package gateway

import (
	stdctx "context"
	"expvar"
	"net"
	"net/http"
	"sort"
	"sync"

	"github.com/ansel1/merry"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/percolate/shisa/auxiliary"
	"github.com/percolate/shisa/httpx"
	"github.com/percolate/shisa/service"
)

type byName []httpx.Field

func (p byName) Len() int           { return len(p) }
func (p byName) Less(i, j int) bool { return p[i].Name < p[j].Name }
func (p byName) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (g *Gateway) Serve(services []service.Service, auxiliaries ...auxiliary.Server) error {
	return g.serve(false, services, auxiliaries)
}

func (g *Gateway) ServeTLS(services []service.Service, auxiliaries ...auxiliary.Server) error {
	return g.serve(true, services, auxiliaries)
}

func (g *Gateway) Shutdown() (err error) {
	g.Logger.Info("shutting down gateway")
	ctx, cancel := stdctx.WithTimeout(stdctx.Background(), g.GracePeriod)
	defer cancel()

	err = merry.Wrap(g.base.Shutdown(ctx))

	for _, aux := range g.auxiliaries {
		g.Logger.Info("shutting down auxiliary", zap.String("name", aux.Name()))
		err = multierr.Append(err, merry.Wrap(aux.Shutdown(g.GracePeriod)))
	}

	g.started = false
	return
}

func (g *Gateway) serve(tls bool, services []service.Service, auxiliaries []auxiliary.Server) (err error) {
	if len(services) == 0 {
		return merry.New("services must not be empty")
	}

	g.init()
	defer g.Logger.Sync()

	if err := g.installServices(services); err != nil {
		return err
	}

	g.auxiliaries = auxiliaries

	ach := make(chan error, len(g.auxiliaries))
	var wg sync.WaitGroup
	for _, aux := range g.auxiliaries {
		wg.Add(1)
		go g.safelyRunAuxiliary(aux, ach, &wg)
	}

	wg.Wait()

	listener, err := httpx.HTTPListenerForAddress(g.Address)
	if err != nil {
		return err
	}

	addr := listener.Addr().String()
	if err := g.register(addr); err != nil {
		return err
	}

	gch := make(chan error, 1)
	go func() {
		g.Logger.Info("gateway started", zap.String("addr", addr))
		if tls {
			gch <- g.base.ServeTLS(l, "", "")
		} else {
			gch <- g.base.Serve(l)
		}
	}()

	for i := len(g.auxiliaries) + 1; i != 0; i-- {
		select {
		case aerr := <-ach:
			if !merry.Is(aerr, http.ErrServerClosed) {
				err = multierr.Append(err, merry.Wrap(aerr))
			}
		case gerr := <-gch:
			if gerr != http.ErrServerClosed {
				err = multierr.Append(err, merry.Wrap(gerr))
			}
		}
	}

	return
}

func (g *Gateway) register(addr string) error {
	if g.Registrar == nil {
		return nil
	}

	if err = g.Registrar.Register(g.Name, addr); err != nil {
		return err
	}
	defer func() {
		if err1 := g.Registrar.Deregister(g.Name); err1 != nil {
			if err == nil {
				err = err1.Prepend("deregister")
			}
		}
	}()

	if g.CheckURLHook == nil {
		return nil
	}

	u, err := g.CheckURLHook()
	if err != nil {
		return err
	}

	if err = g.Registrar.AddCheck(g.Name, u); err != nil {
		return err
	}
	defer func() {
		if err1 := g.Registrar.RemoveChecks(g.Name); err1 != nil {
			if err == nil {
				err = err1.Prepend("remove checks")
			}
		}
	}()
}

func (g *Gateway) safelyRunAuxiliary(server auxiliary.Server, ch chan error, wg *sync.WaitGroup) {
	var once sync.Once
	done := func() { wg.Done() }
	defer func() {
		arg := recover()
		if arg == nil {
			return
		}

		once.Do(done)
		if err, ok := arg.(error); ok {
			ch <- merry.WithMessage(err, "panic in auxiliary")
			return
		}

		ch <- merry.New("panic in auxiliary").WithValue("context", arg)
	}()

	err := server.Listen()
	once.Do(done)
	if err != nil {
		ch <- err
		return
	}
	g.Logger.Info("starting auxiliary server", zap.String("name", server.Name()), zap.String("addr", server.Address()))
	ch <- server.Serve()
}

func (g *Gateway) installServices(services []service.Service) merry.Error {
	servicesExpvar := new(expvar.Map)
	gatewayExpvar.Set("services", servicesExpvar)
	for _, svc := range services {
		if svc.Name() == "" {
			return merry.New("service name cannot be empty")
		}
		if len(svc.Endpoints()) == 0 {
			return merry.New("service endpoints cannot be empty").WithValue("service", svc.Name())
		}

		serviceVar := new(expvar.Map)
		servicesExpvar.Set(svc.Name(), serviceVar)

		g.Logger.Info("installing service", zap.String("name", svc.Name()))
		for i, endp := range svc.Endpoints() {
			if endp.Route == "" {
				return merry.New("endpoint route cannot be emtpy").WithValue("service", svc.Name()).WithValue("index", i)
			}
			if endp.Route[0] != '/' {
				return merry.New("endpoint route must begin with '/'").WithValue("service", svc.Name()).WithValue("route", endp.Route)
			}

			e := endpoint{
				Endpoint: service.Endpoint{
					Route: endp.Route,
				},
				serviceName:       svc.Name(),
				badQueryHandler:   svc.MalformedRequestHandler(),
				notAllowedHandler: svc.MethodNotAllowedHandler(),
				redirectHandler:   svc.RedirectHandler(),
				iseHandler:        svc.InternalServerErrorHandler(),
			}

			foundMethod := false
			if endp.Head != nil {
				foundMethod = true
				pipeline, err := installPipeline(svc.Handlers(), endp.Head)
				if err != nil {
					return err.WithValue("service", svc.Name()).WithValue("route", endp.Route).WithValue("method", http.MethodHead)
				}
				e.Head = pipeline
			}
			if endp.Get != nil {
				foundMethod = true
				pipeline, err := installPipeline(svc.Handlers(), endp.Get)
				if err != nil {
					return err.WithValue("service", svc.Name()).WithValue("route", endp.Route).WithValue("method", http.MethodGet)
				}
				e.Get = pipeline
			}
			if endp.Put != nil {
				foundMethod = true
				pipeline, err := installPipeline(svc.Handlers(), endp.Put)
				if err != nil {
					return err.WithValue("service", svc.Name()).WithValue("route", endp.Route).WithValue("method", http.MethodPut)
				}
				e.Put = pipeline
			}
			if endp.Post != nil {
				foundMethod = true
				pipeline, err := installPipeline(svc.Handlers(), endp.Post)
				if err != nil {
					return err.WithValue("service", svc.Name()).WithValue("route", endp.Route).WithValue("method", http.MethodPost)
				}
				e.Post = pipeline
			}
			if endp.Patch != nil {
				foundMethod = true
				pipeline, err := installPipeline(svc.Handlers(), endp.Patch)
				if err != nil {
					return err.WithValue("service", svc.Name()).WithValue("route", endp.Route).WithValue("method", http.MethodPatch)
				}
				e.Patch = pipeline
			}
			if endp.Delete != nil {
				foundMethod = true
				pipeline, err := installPipeline(svc.Handlers(), endp.Delete)
				if err != nil {
					return err.WithValue("service", svc.Name()).WithValue("route", endp.Route).WithValue("method", http.MethodDelete)
				}
				e.Delete = pipeline
			}
			if endp.Connect != nil {
				foundMethod = true
				pipeline, err := installPipeline(svc.Handlers(), endp.Connect)
				if err != nil {
					return err.WithValue("service", svc.Name()).WithValue("route", endp.Route).WithValue("method", http.MethodConnect)
				}
				e.Connect = pipeline
			}
			if endp.Options != nil {
				foundMethod = true
				pipeline, err := installPipeline(svc.Handlers(), endp.Options)
				if err != nil {
					return err.WithValue("service", svc.Name()).WithValue("route", endp.Route).WithValue("method", http.MethodOptions)
				}
				e.Options = pipeline
			}
			if endp.Trace != nil {
				foundMethod = true
				pipeline, err := installPipeline(svc.Handlers(), endp.Trace)
				if err != nil {
					return err.WithValue("service", svc.Name()).WithValue("route", endp.Route).WithValue("method", http.MethodTrace)
				}
				e.Trace = pipeline
			}

			if !foundMethod {
				return merry.New("endpoint requires least one method").WithValue("service", svc.Name()).WithValue("index", i)
			}

			g.Logger.Debug("adding endpoint", zap.String("route", endp.Route))
			if err := g.tree.addRoute(endp.Route, &e); err != nil {
				return err
			}

			serviceVar.Set(e.Route, e)
		}
	}

	return nil
}

func installPipeline(handlers []httpx.Handler, pipeline *service.Pipeline) (*service.Pipeline, merry.Error) {
	for _, field := range pipeline.QueryFields {
		if field.Default != "" && field.Name == "" {
			return nil, merry.New("Field default requires name")
		}
	}

	result := &service.Pipeline{
		Policy:      pipeline.Policy,
		Handlers:    append(handlers, pipeline.Handlers...),
		QueryFields: append([]httpx.Field(nil), pipeline.QueryFields...),
	}
	sort.Sort(byName(result.QueryFields))

	return result, nil
}
