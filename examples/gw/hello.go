package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"time"

	"github.com/ansel1/merry"

	"github.com/percolate/shisa/context"
	"github.com/percolate/shisa/examples/rpc/service"
	"github.com/percolate/shisa/httpx"
	"github.com/percolate/shisa/sd"
	"github.com/percolate/shisa/service"
)

var (
	language = httpx.Field{
		Name:         "language",
		Default:      hello.AmericanEnglish,
		Validator:    httpx.StringSliceValidator{Target: hello.SupportedLanguages}.Validate,
		Multiplicity: 1,
	}
)

type Greeting struct {
	Message string
}

func (g Greeting) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"greeting\": %q}", g.Message)), nil
}

type HelloService struct {
	service.Service
	resolver sd.Resolver
}

func NewHelloService(res sd.Resolver) *HelloService {
	policy := service.Policy{
		TimeBudget:                  time.Millisecond * 5,
		AllowTrailingSlashRedirects: true,
	}

	svc := &HelloService{
		Service: service.Service{
			Name: "hello",
		},
		resolver: res,
	}

	greeting := service.GetEndpointWithPolicy("/api/greeting", policy, svc.Greeting)
	greeting.Get.QueryFields = []httpx.Field{
		language,
		{Name: "name", Multiplicity: 1},
	}

	svc.Endpoints = []service.Endpoint{greeting}

	return svc
}

func (s *HelloService) Name() string {
	return s.Service.Name
}

func (s *HelloService) Greeting(ctx context.Context, r *httpx.Request) httpx.Response {
	client, err := s.connect()
	if err != nil {
		return httpx.NewEmptyError(http.StatusInternalServerError, err)
	}

	message := hello.Message{
		RequestID: ctx.RequestID(),
		UserID:    ctx.Actor().ID(),
	}

	for _, param := range r.QueryParams {
		switch param.Name {
		case "language":
			message.Language = param.Values[0]
		case "name":
			message.Name = param.Values[0]
		}
	}

	var reply string
	rpcErr := client.Call("Hello.Greeting", &message, &reply)
	if rpcErr != nil {
		return httpx.NewEmptyError(http.StatusInternalServerError, rpcErr)
	}

	response := httpx.NewOK(Greeting{reply})
	addCommonHeaders(response)

	return response
}

func (s *HelloService) Healthcheck(ctx context.Context) merry.Error {
	addrs, err := s.resolver.Resolve(s.Service.Name)
	if err != nil {
		return err.Prepend("healthcheck")
	}

	if len(addrs) < 1 {
		return merry.New("no healthy hosts")
	}

	return nil
}

func (s *HelloService) connect() (*rpc.Client, merry.Error) {
	addrs, err := s.resolver.Resolve(s.Service.Name)
	if err != nil {
		return nil, err.Prepend("connect")
	}

	if len(addrs) < 1 {
		return nil, merry.New("no healthy hosts")
	}

	client, rpcErr := rpc.DialHTTP("tcp", addrs[0])
	if rpcErr != nil {
		return nil, merry.Prepend(rpcErr, "connect")
	}

	return client, nil
}
