// generated by "charlatan -output=./service_charlatan.go Service".  DO NOT EDIT.

package service

import (
	"testing"
)

// NameInvocation represents a single call of FakeService.Name
type NameInvocation struct {
	Results struct {
		Ident95 string
	}
}

// EndpointsInvocation represents a single call of FakeService.Endpoints
type EndpointsInvocation struct {
	Results struct {
		Ident96 []Endpoint
	}
}

/*
FakeService is a mock implementation of Service for testing.
Use it in your tests as in this example:

	package example

	func TestWithService(t *testing.T) {
		f := &service.FakeService{
			NameHook: func() (ident95 string) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeName ...
		f.AssertNameCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a painc if any
unexpected calls are made to FakeName.
*/
type FakeService struct {
	NameHook      func() string
	EndpointsHook func() []Endpoint

	NameCalls      []*NameInvocation
	EndpointsCalls []*EndpointsInvocation
}

// NewFakeServiceDefaultPanic returns an instance of FakeService with all hooks configured to panic
func NewFakeServiceDefaultPanic() *FakeService {
	return &FakeService{
		NameHook: func() (ident95 string) {
			panic("Unexpected call to Service.Name")
			return
		},
		EndpointsHook: func() (ident96 []Endpoint) {
			panic("Unexpected call to Service.Endpoints")
			return
		},
	}
}

// NewFakeServiceDefaultFatal returns an instance of FakeService with all hooks configured to call t.Fatal
func NewFakeServiceDefaultFatal(t *testing.T) *FakeService {
	return &FakeService{
		NameHook: func() (ident95 string) {
			t.Fatal("Unexpected call to Service.Name")
			return
		},
		EndpointsHook: func() (ident96 []Endpoint) {
			t.Fatal("Unexpected call to Service.Endpoints")
			return
		},
	}
}

// NewFakeServiceDefaultError returns an instance of FakeService with all hooks configured to call t.Error
func NewFakeServiceDefaultError(t *testing.T) *FakeService {
	return &FakeService{
		NameHook: func() (ident95 string) {
			t.Error("Unexpected call to Service.Name")
			return
		},
		EndpointsHook: func() (ident96 []Endpoint) {
			t.Error("Unexpected call to Service.Endpoints")
			return
		},
	}
}

func (_f1 *FakeService) Name() (ident95 string) {
	invocation := new(NameInvocation)

	ident95 = _f1.NameHook()

	invocation.Results.Ident95 = ident95

	_f1.NameCalls = append(_f1.NameCalls, invocation)

	return
}

// NameCalled returns true if FakeService.Name was called
func (f *FakeService) NameCalled() bool {
	return len(f.NameCalls) != 0
}

// AssertNameCalled calls t.Error if FakeService.Name was not called
func (f *FakeService) AssertNameCalled(t *testing.T) {
	t.Helper()
	if len(f.NameCalls) == 0 {
		t.Error("FakeService.Name not called, expected at least one")
	}
}

// NameNotCalled returns true if FakeService.Name was not called
func (f *FakeService) NameNotCalled() bool {
	return len(f.NameCalls) == 0
}

// AssertNameNotCalled calls t.Error if FakeService.Name was called
func (f *FakeService) AssertNameNotCalled(t *testing.T) {
	t.Helper()
	if len(f.NameCalls) != 0 {
		t.Error("FakeService.Name called, expected none")
	}
}

// NameCalledOnce returns true if FakeService.Name was called exactly once
func (f *FakeService) NameCalledOnce() bool {
	return len(f.NameCalls) == 1
}

// AssertNameCalledOnce calls t.Error if FakeService.Name was not called exactly once
func (f *FakeService) AssertNameCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.NameCalls) != 1 {
		t.Errorf("FakeService.Name called %d times, expected 1", len(f.NameCalls))
	}
}

// NameCalledN returns true if FakeService.Name was called at least n times
func (f *FakeService) NameCalledN(n int) bool {
	return len(f.NameCalls) >= n
}

// AssertNameCalledN calls t.Error if FakeService.Name was called less than n times
func (f *FakeService) AssertNameCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.NameCalls) < n {
		t.Errorf("FakeService.Name called %d times, expected >= %d", len(f.NameCalls), n)
	}
}

func (_f2 *FakeService) Endpoints() (ident96 []Endpoint) {
	invocation := new(EndpointsInvocation)

	ident96 = _f2.EndpointsHook()

	invocation.Results.Ident96 = ident96

	_f2.EndpointsCalls = append(_f2.EndpointsCalls, invocation)

	return
}

// EndpointsCalled returns true if FakeService.Endpoints was called
func (f *FakeService) EndpointsCalled() bool {
	return len(f.EndpointsCalls) != 0
}

// AssertEndpointsCalled calls t.Error if FakeService.Endpoints was not called
func (f *FakeService) AssertEndpointsCalled(t *testing.T) {
	t.Helper()
	if len(f.EndpointsCalls) == 0 {
		t.Error("FakeService.Endpoints not called, expected at least one")
	}
}

// EndpointsNotCalled returns true if FakeService.Endpoints was not called
func (f *FakeService) EndpointsNotCalled() bool {
	return len(f.EndpointsCalls) == 0
}

// AssertEndpointsNotCalled calls t.Error if FakeService.Endpoints was called
func (f *FakeService) AssertEndpointsNotCalled(t *testing.T) {
	t.Helper()
	if len(f.EndpointsCalls) != 0 {
		t.Error("FakeService.Endpoints called, expected none")
	}
}

// EndpointsCalledOnce returns true if FakeService.Endpoints was called exactly once
func (f *FakeService) EndpointsCalledOnce() bool {
	return len(f.EndpointsCalls) == 1
}

// AssertEndpointsCalledOnce calls t.Error if FakeService.Endpoints was not called exactly once
func (f *FakeService) AssertEndpointsCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.EndpointsCalls) != 1 {
		t.Errorf("FakeService.Endpoints called %d times, expected 1", len(f.EndpointsCalls))
	}
}

// EndpointsCalledN returns true if FakeService.Endpoints was called at least n times
func (f *FakeService) EndpointsCalledN(n int) bool {
	return len(f.EndpointsCalls) >= n
}

// AssertEndpointsCalledN calls t.Error if FakeService.Endpoints was called less than n times
func (f *FakeService) AssertEndpointsCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.EndpointsCalls) < n {
		t.Errorf("FakeService.Endpoints called %d times, expected >= %d", len(f.EndpointsCalls), n)
	}
}
