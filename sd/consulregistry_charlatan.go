// generated by "charlatan -output=./consulregistry_charlatan.go consulRegistry".  DO NOT EDIT.

package sd

import "reflect"

import consul "github.com/hashicorp/consul/api"

// consulRegistryServiceRegisterInvocation represents a single call of FakeconsulRegistry.ServiceRegister
type consulRegistryServiceRegisterInvocation struct {
	Parameters struct {
		Ident1 *consul.AgentServiceRegistration
	}
	Results struct {
		Ident2 error
	}
}

// consulRegistryServiceDeregisterInvocation represents a single call of FakeconsulRegistry.ServiceDeregister
type consulRegistryServiceDeregisterInvocation struct {
	Parameters struct {
		Ident1 string
	}
	Results struct {
		Ident2 error
	}
}

// consulRegistryCheckRegisterInvocation represents a single call of FakeconsulRegistry.CheckRegister
type consulRegistryCheckRegisterInvocation struct {
	Parameters struct {
		Ident1 *consul.AgentCheckRegistration
	}
	Results struct {
		Ident2 error
	}
}

// consulRegistryCheckDeregisterInvocation represents a single call of FakeconsulRegistry.CheckDeregister
type consulRegistryCheckDeregisterInvocation struct {
	Parameters struct {
		Ident1 string
	}
	Results struct {
		Ident2 error
	}
}

// consulRegistryTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type consulRegistryTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeconsulRegistry is a mock implementation of consulRegistry for testing.
Use it in your tests as in this example:

	package example

	func TestWithconsulRegistry(t *testing.T) {
		f := &sd.FakeconsulRegistry{
			ServiceRegisterHook: func(ident1 *consul.AgentServiceRegistration) (ident2 error) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeServiceRegister ...
		f.AssertServiceRegisterCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeServiceRegister.
*/
type FakeconsulRegistry struct {
	ServiceRegisterHook   func(*consul.AgentServiceRegistration) error
	ServiceDeregisterHook func(string) error
	CheckRegisterHook     func(*consul.AgentCheckRegistration) error
	CheckDeregisterHook   func(string) error

	ServiceRegisterCalls   []*consulRegistryServiceRegisterInvocation
	ServiceDeregisterCalls []*consulRegistryServiceDeregisterInvocation
	CheckRegisterCalls     []*consulRegistryCheckRegisterInvocation
	CheckDeregisterCalls   []*consulRegistryCheckDeregisterInvocation
}

// NewFakeconsulRegistryDefaultPanic returns an instance of FakeconsulRegistry with all hooks configured to panic
func NewFakeconsulRegistryDefaultPanic() *FakeconsulRegistry {
	return &FakeconsulRegistry{
		ServiceRegisterHook: func(*consul.AgentServiceRegistration) (ident2 error) {
			panic("Unexpected call to consulRegistry.ServiceRegister")
		},
		ServiceDeregisterHook: func(string) (ident2 error) {
			panic("Unexpected call to consulRegistry.ServiceDeregister")
		},
		CheckRegisterHook: func(*consul.AgentCheckRegistration) (ident2 error) {
			panic("Unexpected call to consulRegistry.CheckRegister")
		},
		CheckDeregisterHook: func(string) (ident2 error) {
			panic("Unexpected call to consulRegistry.CheckDeregister")
		},
	}
}

// NewFakeconsulRegistryDefaultFatal returns an instance of FakeconsulRegistry with all hooks configured to call t.Fatal
func NewFakeconsulRegistryDefaultFatal(t consulRegistryTestingT) *FakeconsulRegistry {
	return &FakeconsulRegistry{
		ServiceRegisterHook: func(*consul.AgentServiceRegistration) (ident2 error) {
			t.Fatal("Unexpected call to consulRegistry.ServiceRegister")
			return
		},
		ServiceDeregisterHook: func(string) (ident2 error) {
			t.Fatal("Unexpected call to consulRegistry.ServiceDeregister")
			return
		},
		CheckRegisterHook: func(*consul.AgentCheckRegistration) (ident2 error) {
			t.Fatal("Unexpected call to consulRegistry.CheckRegister")
			return
		},
		CheckDeregisterHook: func(string) (ident2 error) {
			t.Fatal("Unexpected call to consulRegistry.CheckDeregister")
			return
		},
	}
}

// NewFakeconsulRegistryDefaultError returns an instance of FakeconsulRegistry with all hooks configured to call t.Error
func NewFakeconsulRegistryDefaultError(t consulRegistryTestingT) *FakeconsulRegistry {
	return &FakeconsulRegistry{
		ServiceRegisterHook: func(*consul.AgentServiceRegistration) (ident2 error) {
			t.Error("Unexpected call to consulRegistry.ServiceRegister")
			return
		},
		ServiceDeregisterHook: func(string) (ident2 error) {
			t.Error("Unexpected call to consulRegistry.ServiceDeregister")
			return
		},
		CheckRegisterHook: func(*consul.AgentCheckRegistration) (ident2 error) {
			t.Error("Unexpected call to consulRegistry.CheckRegister")
			return
		},
		CheckDeregisterHook: func(string) (ident2 error) {
			t.Error("Unexpected call to consulRegistry.CheckDeregister")
			return
		},
	}
}

func (f *FakeconsulRegistry) Reset() {
	f.ServiceRegisterCalls = []*consulRegistryServiceRegisterInvocation{}
	f.ServiceDeregisterCalls = []*consulRegistryServiceDeregisterInvocation{}
	f.CheckRegisterCalls = []*consulRegistryCheckRegisterInvocation{}
	f.CheckDeregisterCalls = []*consulRegistryCheckDeregisterInvocation{}
}

func (_f1 *FakeconsulRegistry) ServiceRegister(ident1 *consul.AgentServiceRegistration) (ident2 error) {
	if _f1.ServiceRegisterHook == nil {
		panic("consulRegistry.ServiceRegister() called but FakeconsulRegistry.ServiceRegisterHook is nil")
	}

	invocation := new(consulRegistryServiceRegisterInvocation)
	_f1.ServiceRegisterCalls = append(_f1.ServiceRegisterCalls, invocation)

	invocation.Parameters.Ident1 = ident1

	ident2 = _f1.ServiceRegisterHook(ident1)

	invocation.Results.Ident2 = ident2

	return
}

// ServiceRegisterCalled returns true if FakeconsulRegistry.ServiceRegister was called
func (f *FakeconsulRegistry) ServiceRegisterCalled() bool {
	return len(f.ServiceRegisterCalls) != 0
}

// AssertServiceRegisterCalled calls t.Error if FakeconsulRegistry.ServiceRegister was not called
func (f *FakeconsulRegistry) AssertServiceRegisterCalled(t consulRegistryTestingT) {
	t.Helper()
	if len(f.ServiceRegisterCalls) == 0 {
		t.Error("FakeconsulRegistry.ServiceRegister not called, expected at least one")
	}
}

// ServiceRegisterNotCalled returns true if FakeconsulRegistry.ServiceRegister was not called
func (f *FakeconsulRegistry) ServiceRegisterNotCalled() bool {
	return len(f.ServiceRegisterCalls) == 0
}

// AssertServiceRegisterNotCalled calls t.Error if FakeconsulRegistry.ServiceRegister was called
func (f *FakeconsulRegistry) AssertServiceRegisterNotCalled(t consulRegistryTestingT) {
	t.Helper()
	if len(f.ServiceRegisterCalls) != 0 {
		t.Error("FakeconsulRegistry.ServiceRegister called, expected none")
	}
}

// ServiceRegisterCalledOnce returns true if FakeconsulRegistry.ServiceRegister was called exactly once
func (f *FakeconsulRegistry) ServiceRegisterCalledOnce() bool {
	return len(f.ServiceRegisterCalls) == 1
}

// AssertServiceRegisterCalledOnce calls t.Error if FakeconsulRegistry.ServiceRegister was not called exactly once
func (f *FakeconsulRegistry) AssertServiceRegisterCalledOnce(t consulRegistryTestingT) {
	t.Helper()
	if len(f.ServiceRegisterCalls) != 1 {
		t.Errorf("FakeconsulRegistry.ServiceRegister called %d times, expected 1", len(f.ServiceRegisterCalls))
	}
}

// ServiceRegisterCalledN returns true if FakeconsulRegistry.ServiceRegister was called at least n times
func (f *FakeconsulRegistry) ServiceRegisterCalledN(n int) bool {
	return len(f.ServiceRegisterCalls) >= n
}

// AssertServiceRegisterCalledN calls t.Error if FakeconsulRegistry.ServiceRegister was called less than n times
func (f *FakeconsulRegistry) AssertServiceRegisterCalledN(t consulRegistryTestingT, n int) {
	t.Helper()
	if len(f.ServiceRegisterCalls) < n {
		t.Errorf("FakeconsulRegistry.ServiceRegister called %d times, expected >= %d", len(f.ServiceRegisterCalls), n)
	}
}

// ServiceRegisterCalledWith returns true if FakeconsulRegistry.ServiceRegister was called with the given values
func (_f2 *FakeconsulRegistry) ServiceRegisterCalledWith(ident1 *consul.AgentServiceRegistration) (found bool) {
	for _, call := range _f2.ServiceRegisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertServiceRegisterCalledWith calls t.Error if FakeconsulRegistry.ServiceRegister was not called with the given values
func (_f3 *FakeconsulRegistry) AssertServiceRegisterCalledWith(t consulRegistryTestingT, ident1 *consul.AgentServiceRegistration) {
	t.Helper()
	var found bool
	for _, call := range _f3.ServiceRegisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeconsulRegistry.ServiceRegister not called with expected parameters")
	}
}

// ServiceRegisterCalledOnceWith returns true if FakeconsulRegistry.ServiceRegister was called exactly once with the given values
func (_f4 *FakeconsulRegistry) ServiceRegisterCalledOnceWith(ident1 *consul.AgentServiceRegistration) bool {
	var count int
	for _, call := range _f4.ServiceRegisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertServiceRegisterCalledOnceWith calls t.Error if FakeconsulRegistry.ServiceRegister was not called exactly once with the given values
func (_f5 *FakeconsulRegistry) AssertServiceRegisterCalledOnceWith(t consulRegistryTestingT, ident1 *consul.AgentServiceRegistration) {
	t.Helper()
	var count int
	for _, call := range _f5.ServiceRegisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeconsulRegistry.ServiceRegister called %d times with expected parameters, expected one", count)
	}
}

// ServiceRegisterResultsForCall returns the result values for the first call to FakeconsulRegistry.ServiceRegister with the given values
func (_f6 *FakeconsulRegistry) ServiceRegisterResultsForCall(ident1 *consul.AgentServiceRegistration) (ident2 error, found bool) {
	for _, call := range _f6.ServiceRegisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			found = true
			break
		}
	}

	return
}

func (_f7 *FakeconsulRegistry) ServiceDeregister(ident1 string) (ident2 error) {
	if _f7.ServiceDeregisterHook == nil {
		panic("consulRegistry.ServiceDeregister() called but FakeconsulRegistry.ServiceDeregisterHook is nil")
	}

	invocation := new(consulRegistryServiceDeregisterInvocation)
	_f7.ServiceDeregisterCalls = append(_f7.ServiceDeregisterCalls, invocation)

	invocation.Parameters.Ident1 = ident1

	ident2 = _f7.ServiceDeregisterHook(ident1)

	invocation.Results.Ident2 = ident2

	return
}

// ServiceDeregisterCalled returns true if FakeconsulRegistry.ServiceDeregister was called
func (f *FakeconsulRegistry) ServiceDeregisterCalled() bool {
	return len(f.ServiceDeregisterCalls) != 0
}

// AssertServiceDeregisterCalled calls t.Error if FakeconsulRegistry.ServiceDeregister was not called
func (f *FakeconsulRegistry) AssertServiceDeregisterCalled(t consulRegistryTestingT) {
	t.Helper()
	if len(f.ServiceDeregisterCalls) == 0 {
		t.Error("FakeconsulRegistry.ServiceDeregister not called, expected at least one")
	}
}

// ServiceDeregisterNotCalled returns true if FakeconsulRegistry.ServiceDeregister was not called
func (f *FakeconsulRegistry) ServiceDeregisterNotCalled() bool {
	return len(f.ServiceDeregisterCalls) == 0
}

// AssertServiceDeregisterNotCalled calls t.Error if FakeconsulRegistry.ServiceDeregister was called
func (f *FakeconsulRegistry) AssertServiceDeregisterNotCalled(t consulRegistryTestingT) {
	t.Helper()
	if len(f.ServiceDeregisterCalls) != 0 {
		t.Error("FakeconsulRegistry.ServiceDeregister called, expected none")
	}
}

// ServiceDeregisterCalledOnce returns true if FakeconsulRegistry.ServiceDeregister was called exactly once
func (f *FakeconsulRegistry) ServiceDeregisterCalledOnce() bool {
	return len(f.ServiceDeregisterCalls) == 1
}

// AssertServiceDeregisterCalledOnce calls t.Error if FakeconsulRegistry.ServiceDeregister was not called exactly once
func (f *FakeconsulRegistry) AssertServiceDeregisterCalledOnce(t consulRegistryTestingT) {
	t.Helper()
	if len(f.ServiceDeregisterCalls) != 1 {
		t.Errorf("FakeconsulRegistry.ServiceDeregister called %d times, expected 1", len(f.ServiceDeregisterCalls))
	}
}

// ServiceDeregisterCalledN returns true if FakeconsulRegistry.ServiceDeregister was called at least n times
func (f *FakeconsulRegistry) ServiceDeregisterCalledN(n int) bool {
	return len(f.ServiceDeregisterCalls) >= n
}

// AssertServiceDeregisterCalledN calls t.Error if FakeconsulRegistry.ServiceDeregister was called less than n times
func (f *FakeconsulRegistry) AssertServiceDeregisterCalledN(t consulRegistryTestingT, n int) {
	t.Helper()
	if len(f.ServiceDeregisterCalls) < n {
		t.Errorf("FakeconsulRegistry.ServiceDeregister called %d times, expected >= %d", len(f.ServiceDeregisterCalls), n)
	}
}

// ServiceDeregisterCalledWith returns true if FakeconsulRegistry.ServiceDeregister was called with the given values
func (_f8 *FakeconsulRegistry) ServiceDeregisterCalledWith(ident1 string) (found bool) {
	for _, call := range _f8.ServiceDeregisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertServiceDeregisterCalledWith calls t.Error if FakeconsulRegistry.ServiceDeregister was not called with the given values
func (_f9 *FakeconsulRegistry) AssertServiceDeregisterCalledWith(t consulRegistryTestingT, ident1 string) {
	t.Helper()
	var found bool
	for _, call := range _f9.ServiceDeregisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeconsulRegistry.ServiceDeregister not called with expected parameters")
	}
}

// ServiceDeregisterCalledOnceWith returns true if FakeconsulRegistry.ServiceDeregister was called exactly once with the given values
func (_f10 *FakeconsulRegistry) ServiceDeregisterCalledOnceWith(ident1 string) bool {
	var count int
	for _, call := range _f10.ServiceDeregisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertServiceDeregisterCalledOnceWith calls t.Error if FakeconsulRegistry.ServiceDeregister was not called exactly once with the given values
func (_f11 *FakeconsulRegistry) AssertServiceDeregisterCalledOnceWith(t consulRegistryTestingT, ident1 string) {
	t.Helper()
	var count int
	for _, call := range _f11.ServiceDeregisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeconsulRegistry.ServiceDeregister called %d times with expected parameters, expected one", count)
	}
}

// ServiceDeregisterResultsForCall returns the result values for the first call to FakeconsulRegistry.ServiceDeregister with the given values
func (_f12 *FakeconsulRegistry) ServiceDeregisterResultsForCall(ident1 string) (ident2 error, found bool) {
	for _, call := range _f12.ServiceDeregisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			found = true
			break
		}
	}

	return
}

func (_f13 *FakeconsulRegistry) CheckRegister(ident1 *consul.AgentCheckRegistration) (ident2 error) {
	if _f13.CheckRegisterHook == nil {
		panic("consulRegistry.CheckRegister() called but FakeconsulRegistry.CheckRegisterHook is nil")
	}

	invocation := new(consulRegistryCheckRegisterInvocation)
	_f13.CheckRegisterCalls = append(_f13.CheckRegisterCalls, invocation)

	invocation.Parameters.Ident1 = ident1

	ident2 = _f13.CheckRegisterHook(ident1)

	invocation.Results.Ident2 = ident2

	return
}

// CheckRegisterCalled returns true if FakeconsulRegistry.CheckRegister was called
func (f *FakeconsulRegistry) CheckRegisterCalled() bool {
	return len(f.CheckRegisterCalls) != 0
}

// AssertCheckRegisterCalled calls t.Error if FakeconsulRegistry.CheckRegister was not called
func (f *FakeconsulRegistry) AssertCheckRegisterCalled(t consulRegistryTestingT) {
	t.Helper()
	if len(f.CheckRegisterCalls) == 0 {
		t.Error("FakeconsulRegistry.CheckRegister not called, expected at least one")
	}
}

// CheckRegisterNotCalled returns true if FakeconsulRegistry.CheckRegister was not called
func (f *FakeconsulRegistry) CheckRegisterNotCalled() bool {
	return len(f.CheckRegisterCalls) == 0
}

// AssertCheckRegisterNotCalled calls t.Error if FakeconsulRegistry.CheckRegister was called
func (f *FakeconsulRegistry) AssertCheckRegisterNotCalled(t consulRegistryTestingT) {
	t.Helper()
	if len(f.CheckRegisterCalls) != 0 {
		t.Error("FakeconsulRegistry.CheckRegister called, expected none")
	}
}

// CheckRegisterCalledOnce returns true if FakeconsulRegistry.CheckRegister was called exactly once
func (f *FakeconsulRegistry) CheckRegisterCalledOnce() bool {
	return len(f.CheckRegisterCalls) == 1
}

// AssertCheckRegisterCalledOnce calls t.Error if FakeconsulRegistry.CheckRegister was not called exactly once
func (f *FakeconsulRegistry) AssertCheckRegisterCalledOnce(t consulRegistryTestingT) {
	t.Helper()
	if len(f.CheckRegisterCalls) != 1 {
		t.Errorf("FakeconsulRegistry.CheckRegister called %d times, expected 1", len(f.CheckRegisterCalls))
	}
}

// CheckRegisterCalledN returns true if FakeconsulRegistry.CheckRegister was called at least n times
func (f *FakeconsulRegistry) CheckRegisterCalledN(n int) bool {
	return len(f.CheckRegisterCalls) >= n
}

// AssertCheckRegisterCalledN calls t.Error if FakeconsulRegistry.CheckRegister was called less than n times
func (f *FakeconsulRegistry) AssertCheckRegisterCalledN(t consulRegistryTestingT, n int) {
	t.Helper()
	if len(f.CheckRegisterCalls) < n {
		t.Errorf("FakeconsulRegistry.CheckRegister called %d times, expected >= %d", len(f.CheckRegisterCalls), n)
	}
}

// CheckRegisterCalledWith returns true if FakeconsulRegistry.CheckRegister was called with the given values
func (_f14 *FakeconsulRegistry) CheckRegisterCalledWith(ident1 *consul.AgentCheckRegistration) (found bool) {
	for _, call := range _f14.CheckRegisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertCheckRegisterCalledWith calls t.Error if FakeconsulRegistry.CheckRegister was not called with the given values
func (_f15 *FakeconsulRegistry) AssertCheckRegisterCalledWith(t consulRegistryTestingT, ident1 *consul.AgentCheckRegistration) {
	t.Helper()
	var found bool
	for _, call := range _f15.CheckRegisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeconsulRegistry.CheckRegister not called with expected parameters")
	}
}

// CheckRegisterCalledOnceWith returns true if FakeconsulRegistry.CheckRegister was called exactly once with the given values
func (_f16 *FakeconsulRegistry) CheckRegisterCalledOnceWith(ident1 *consul.AgentCheckRegistration) bool {
	var count int
	for _, call := range _f16.CheckRegisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertCheckRegisterCalledOnceWith calls t.Error if FakeconsulRegistry.CheckRegister was not called exactly once with the given values
func (_f17 *FakeconsulRegistry) AssertCheckRegisterCalledOnceWith(t consulRegistryTestingT, ident1 *consul.AgentCheckRegistration) {
	t.Helper()
	var count int
	for _, call := range _f17.CheckRegisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeconsulRegistry.CheckRegister called %d times with expected parameters, expected one", count)
	}
}

// CheckRegisterResultsForCall returns the result values for the first call to FakeconsulRegistry.CheckRegister with the given values
func (_f18 *FakeconsulRegistry) CheckRegisterResultsForCall(ident1 *consul.AgentCheckRegistration) (ident2 error, found bool) {
	for _, call := range _f18.CheckRegisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			found = true
			break
		}
	}

	return
}

func (_f19 *FakeconsulRegistry) CheckDeregister(ident1 string) (ident2 error) {
	if _f19.CheckDeregisterHook == nil {
		panic("consulRegistry.CheckDeregister() called but FakeconsulRegistry.CheckDeregisterHook is nil")
	}

	invocation := new(consulRegistryCheckDeregisterInvocation)
	_f19.CheckDeregisterCalls = append(_f19.CheckDeregisterCalls, invocation)

	invocation.Parameters.Ident1 = ident1

	ident2 = _f19.CheckDeregisterHook(ident1)

	invocation.Results.Ident2 = ident2

	return
}

// CheckDeregisterCalled returns true if FakeconsulRegistry.CheckDeregister was called
func (f *FakeconsulRegistry) CheckDeregisterCalled() bool {
	return len(f.CheckDeregisterCalls) != 0
}

// AssertCheckDeregisterCalled calls t.Error if FakeconsulRegistry.CheckDeregister was not called
func (f *FakeconsulRegistry) AssertCheckDeregisterCalled(t consulRegistryTestingT) {
	t.Helper()
	if len(f.CheckDeregisterCalls) == 0 {
		t.Error("FakeconsulRegistry.CheckDeregister not called, expected at least one")
	}
}

// CheckDeregisterNotCalled returns true if FakeconsulRegistry.CheckDeregister was not called
func (f *FakeconsulRegistry) CheckDeregisterNotCalled() bool {
	return len(f.CheckDeregisterCalls) == 0
}

// AssertCheckDeregisterNotCalled calls t.Error if FakeconsulRegistry.CheckDeregister was called
func (f *FakeconsulRegistry) AssertCheckDeregisterNotCalled(t consulRegistryTestingT) {
	t.Helper()
	if len(f.CheckDeregisterCalls) != 0 {
		t.Error("FakeconsulRegistry.CheckDeregister called, expected none")
	}
}

// CheckDeregisterCalledOnce returns true if FakeconsulRegistry.CheckDeregister was called exactly once
func (f *FakeconsulRegistry) CheckDeregisterCalledOnce() bool {
	return len(f.CheckDeregisterCalls) == 1
}

// AssertCheckDeregisterCalledOnce calls t.Error if FakeconsulRegistry.CheckDeregister was not called exactly once
func (f *FakeconsulRegistry) AssertCheckDeregisterCalledOnce(t consulRegistryTestingT) {
	t.Helper()
	if len(f.CheckDeregisterCalls) != 1 {
		t.Errorf("FakeconsulRegistry.CheckDeregister called %d times, expected 1", len(f.CheckDeregisterCalls))
	}
}

// CheckDeregisterCalledN returns true if FakeconsulRegistry.CheckDeregister was called at least n times
func (f *FakeconsulRegistry) CheckDeregisterCalledN(n int) bool {
	return len(f.CheckDeregisterCalls) >= n
}

// AssertCheckDeregisterCalledN calls t.Error if FakeconsulRegistry.CheckDeregister was called less than n times
func (f *FakeconsulRegistry) AssertCheckDeregisterCalledN(t consulRegistryTestingT, n int) {
	t.Helper()
	if len(f.CheckDeregisterCalls) < n {
		t.Errorf("FakeconsulRegistry.CheckDeregister called %d times, expected >= %d", len(f.CheckDeregisterCalls), n)
	}
}

// CheckDeregisterCalledWith returns true if FakeconsulRegistry.CheckDeregister was called with the given values
func (_f20 *FakeconsulRegistry) CheckDeregisterCalledWith(ident1 string) (found bool) {
	for _, call := range _f20.CheckDeregisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	return
}

// AssertCheckDeregisterCalledWith calls t.Error if FakeconsulRegistry.CheckDeregister was not called with the given values
func (_f21 *FakeconsulRegistry) AssertCheckDeregisterCalledWith(t consulRegistryTestingT, ident1 string) {
	t.Helper()
	var found bool
	for _, call := range _f21.CheckDeregisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeconsulRegistry.CheckDeregister not called with expected parameters")
	}
}

// CheckDeregisterCalledOnceWith returns true if FakeconsulRegistry.CheckDeregister was called exactly once with the given values
func (_f22 *FakeconsulRegistry) CheckDeregisterCalledOnceWith(ident1 string) bool {
	var count int
	for _, call := range _f22.CheckDeregisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	return count == 1
}

// AssertCheckDeregisterCalledOnceWith calls t.Error if FakeconsulRegistry.CheckDeregister was not called exactly once with the given values
func (_f23 *FakeconsulRegistry) AssertCheckDeregisterCalledOnceWith(t consulRegistryTestingT, ident1 string) {
	t.Helper()
	var count int
	for _, call := range _f23.CheckDeregisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeconsulRegistry.CheckDeregister called %d times with expected parameters, expected one", count)
	}
}

// CheckDeregisterResultsForCall returns the result values for the first call to FakeconsulRegistry.CheckDeregister with the given values
func (_f24 *FakeconsulRegistry) CheckDeregisterResultsForCall(ident1 string) (ident2 error, found bool) {
	for _, call := range _f24.CheckDeregisterCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) {
			ident2 = call.Results.Ident2
			found = true
			break
		}
	}

	return
}
