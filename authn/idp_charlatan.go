// generated by "charlatan -output=./idp_charlatan.go IdentityProvider".  DO NOT EDIT.

package authn

import "github.com/ansel1/merry"
import "github.com/percolate/shisa/context"

import "github.com/percolate/shisa/models"
import "reflect"

// IdentityProviderAuthenticateInvocation represents a single call of FakeIdentityProvider.Authenticate
type IdentityProviderAuthenticateInvocation struct {
	Parameters struct {
		Ident1 context.Context
		Ident2 string
	}
	Results struct {
		Ident3 models.User
		Ident4 merry.Error
	}
}

// IdentityProviderTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type IdentityProviderTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeIdentityProvider is a mock implementation of IdentityProvider for testing.
Use it in your tests as in this example:

	package example

	func TestWithIdentityProvider(t *testing.T) {
		f := &authn.FakeIdentityProvider{
			AuthenticateHook: func(ident1 context.Context, ident2 string) (ident3 models.User, ident4 merry.Error) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeAuthenticate ...
		f.AssertAuthenticateCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeAuthenticate.
*/
type FakeIdentityProvider struct {
	AuthenticateHook func(context.Context, string) (models.User, merry.Error)

	AuthenticateCalls []*IdentityProviderAuthenticateInvocation
}

// NewFakeIdentityProviderDefaultPanic returns an instance of FakeIdentityProvider with all hooks configured to panic
func NewFakeIdentityProviderDefaultPanic() *FakeIdentityProvider {
	return &FakeIdentityProvider{
		AuthenticateHook: func(context.Context, string) (ident3 models.User, ident4 merry.Error) {
			panic("Unexpected call to IdentityProvider.Authenticate")
		},
	}
}

// NewFakeIdentityProviderDefaultFatal returns an instance of FakeIdentityProvider with all hooks configured to call t.Fatal
func NewFakeIdentityProviderDefaultFatal(t IdentityProviderTestingT) *FakeIdentityProvider {
	return &FakeIdentityProvider{
		AuthenticateHook: func(context.Context, string) (ident3 models.User, ident4 merry.Error) {
			t.Fatal("Unexpected call to IdentityProvider.Authenticate")
			return
		},
	}
}

// NewFakeIdentityProviderDefaultError returns an instance of FakeIdentityProvider with all hooks configured to call t.Error
func NewFakeIdentityProviderDefaultError(t IdentityProviderTestingT) *FakeIdentityProvider {
	return &FakeIdentityProvider{
		AuthenticateHook: func(context.Context, string) (ident3 models.User, ident4 merry.Error) {
			t.Error("Unexpected call to IdentityProvider.Authenticate")
			return
		},
	}
}

func (f *FakeIdentityProvider) Reset() {
	f.AuthenticateCalls = []*IdentityProviderAuthenticateInvocation{}
}

func (_f1 *FakeIdentityProvider) Authenticate(ident1 context.Context, ident2 string) (ident3 models.User, ident4 merry.Error) {
	invocation := new(IdentityProviderAuthenticateInvocation)

	invocation.Parameters.Ident1 = ident1
	invocation.Parameters.Ident2 = ident2

	ident3, ident4 = _f1.AuthenticateHook(ident1, ident2)

	invocation.Results.Ident3 = ident3
	invocation.Results.Ident4 = ident4

	_f1.AuthenticateCalls = append(_f1.AuthenticateCalls, invocation)

	return
}

// AuthenticateCalled returns true if FakeIdentityProvider.Authenticate was called
func (f *FakeIdentityProvider) AuthenticateCalled() bool {
	return len(f.AuthenticateCalls) != 0
}

// AssertAuthenticateCalled calls t.Error if FakeIdentityProvider.Authenticate was not called
func (f *FakeIdentityProvider) AssertAuthenticateCalled(t IdentityProviderTestingT) {
	t.Helper()
	if len(f.AuthenticateCalls) == 0 {
		t.Error("FakeIdentityProvider.Authenticate not called, expected at least one")
	}
}

// AuthenticateNotCalled returns true if FakeIdentityProvider.Authenticate was not called
func (f *FakeIdentityProvider) AuthenticateNotCalled() bool {
	return len(f.AuthenticateCalls) == 0
}

// AssertAuthenticateNotCalled calls t.Error if FakeIdentityProvider.Authenticate was called
func (f *FakeIdentityProvider) AssertAuthenticateNotCalled(t IdentityProviderTestingT) {
	t.Helper()
	if len(f.AuthenticateCalls) != 0 {
		t.Error("FakeIdentityProvider.Authenticate called, expected none")
	}
}

// AuthenticateCalledOnce returns true if FakeIdentityProvider.Authenticate was called exactly once
func (f *FakeIdentityProvider) AuthenticateCalledOnce() bool {
	return len(f.AuthenticateCalls) == 1
}

// AssertAuthenticateCalledOnce calls t.Error if FakeIdentityProvider.Authenticate was not called exactly once
func (f *FakeIdentityProvider) AssertAuthenticateCalledOnce(t IdentityProviderTestingT) {
	t.Helper()
	if len(f.AuthenticateCalls) != 1 {
		t.Errorf("FakeIdentityProvider.Authenticate called %d times, expected 1", len(f.AuthenticateCalls))
	}
}

// AuthenticateCalledN returns true if FakeIdentityProvider.Authenticate was called at least n times
func (f *FakeIdentityProvider) AuthenticateCalledN(n int) bool {
	return len(f.AuthenticateCalls) >= n
}

// AssertAuthenticateCalledN calls t.Error if FakeIdentityProvider.Authenticate was called less than n times
func (f *FakeIdentityProvider) AssertAuthenticateCalledN(t IdentityProviderTestingT, n int) {
	t.Helper()
	if len(f.AuthenticateCalls) < n {
		t.Errorf("FakeIdentityProvider.Authenticate called %d times, expected >= %d", len(f.AuthenticateCalls), n)
	}
}

// AuthenticateCalledWith returns true if FakeIdentityProvider.Authenticate was called with the given values
func (_f2 *FakeIdentityProvider) AuthenticateCalledWith(ident1 context.Context, ident2 string) (found bool) {
	for _, call := range _f2.AuthenticateCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			found = true
			break
		}
	}

	return
}

// AssertAuthenticateCalledWith calls t.Error if FakeIdentityProvider.Authenticate was not called with the given values
func (_f3 *FakeIdentityProvider) AssertAuthenticateCalledWith(t IdentityProviderTestingT, ident1 context.Context, ident2 string) {
	t.Helper()
	var found bool
	for _, call := range _f3.AuthenticateCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeIdentityProvider.Authenticate not called with expected parameters")
	}
}

// AuthenticateCalledOnceWith returns true if FakeIdentityProvider.Authenticate was called exactly once with the given values
func (_f4 *FakeIdentityProvider) AuthenticateCalledOnceWith(ident1 context.Context, ident2 string) bool {
	var count int
	for _, call := range _f4.AuthenticateCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			count++
		}
	}

	return count == 1
}

// AssertAuthenticateCalledOnceWith calls t.Error if FakeIdentityProvider.Authenticate was not called exactly once with the given values
func (_f5 *FakeIdentityProvider) AssertAuthenticateCalledOnceWith(t IdentityProviderTestingT, ident1 context.Context, ident2 string) {
	t.Helper()
	var count int
	for _, call := range _f5.AuthenticateCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeIdentityProvider.Authenticate called %d times with expected parameters, expected one", count)
	}
}

// AuthenticateResultsForCall returns the result values for the first call to FakeIdentityProvider.Authenticate with the given values
func (_f6 *FakeIdentityProvider) AuthenticateResultsForCall(ident1 context.Context, ident2 string) (ident3 models.User, ident4 merry.Error, found bool) {
	for _, call := range _f6.AuthenticateCalls {
		if reflect.DeepEqual(call.Parameters.Ident1, ident1) && reflect.DeepEqual(call.Parameters.Ident2, ident2) {
			ident3 = call.Results.Ident3
			ident4 = call.Results.Ident4
			found = true
			break
		}
	}

	return
}
