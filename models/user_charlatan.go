// generated by "charlatan -output=./user_charlatan.go User".  DO NOT EDIT.

package models

import (
	"testing"
)

// StringInvocation represents a single call of FakeUser.String
type StringInvocation struct {
	Results struct {
		Ident5 string
	}
}

// IDInvocation represents a single call of FakeUser.ID
type IDInvocation struct {
	Results struct {
		Ident6 string
	}
}

/*
FakeUser is a mock implementation of User for testing.
Use it in your tests as in this example:

	package example

	func TestWithUser(t *testing.T) {
		f := &models.FakeUser{
			StringHook: func() (ident5 string) {
				// ensure parameters meet expections, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeString ...
		f.AssertStringCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeString.
*/
type FakeUser struct {
	StringHook func() string
	IDHook     func() string

	StringCalls []*StringInvocation
	IDCalls     []*IDInvocation
}

// NewFakeUserDefaultPanic returns an instance of FakeUser with all hooks configured to panic
func NewFakeUserDefaultPanic() *FakeUser {
	return &FakeUser{
		StringHook: func() (ident5 string) {
			panic("Unexpected call to User.String")
		},
		IDHook: func() (ident6 string) {
			panic("Unexpected call to User.ID")
		},
	}
}

// NewFakeUserDefaultFatal returns an instance of FakeUser with all hooks configured to call t.Fatal
func NewFakeUserDefaultFatal(t *testing.T) *FakeUser {
	return &FakeUser{
		StringHook: func() (ident5 string) {
			t.Fatal("Unexpected call to User.String")
			return
		},
		IDHook: func() (ident6 string) {
			t.Fatal("Unexpected call to User.ID")
			return
		},
	}
}

// NewFakeUserDefaultError returns an instance of FakeUser with all hooks configured to call t.Error
func NewFakeUserDefaultError(t *testing.T) *FakeUser {
	return &FakeUser{
		StringHook: func() (ident5 string) {
			t.Error("Unexpected call to User.String")
			return
		},
		IDHook: func() (ident6 string) {
			t.Error("Unexpected call to User.ID")
			return
		},
	}
}

func (_f1 *FakeUser) String() (ident5 string) {
	invocation := new(StringInvocation)

	ident5 = _f1.StringHook()

	invocation.Results.Ident5 = ident5

	_f1.StringCalls = append(_f1.StringCalls, invocation)

	return
}

// StringCalled returns true if FakeUser.String was called
func (f *FakeUser) StringCalled() bool {
	return len(f.StringCalls) != 0
}

// AssertStringCalled calls t.Error if FakeUser.String was not called
func (f *FakeUser) AssertStringCalled(t *testing.T) {
	t.Helper()
	if len(f.StringCalls) == 0 {
		t.Error("FakeUser.String not called, expected at least one")
	}
}

// StringNotCalled returns true if FakeUser.String was not called
func (f *FakeUser) StringNotCalled() bool {
	return len(f.StringCalls) == 0
}

// AssertStringNotCalled calls t.Error if FakeUser.String was called
func (f *FakeUser) AssertStringNotCalled(t *testing.T) {
	t.Helper()
	if len(f.StringCalls) != 0 {
		t.Error("FakeUser.String called, expected none")
	}
}

// StringCalledOnce returns true if FakeUser.String was called exactly once
func (f *FakeUser) StringCalledOnce() bool {
	return len(f.StringCalls) == 1
}

// AssertStringCalledOnce calls t.Error if FakeUser.String was not called exactly once
func (f *FakeUser) AssertStringCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.StringCalls) != 1 {
		t.Errorf("FakeUser.String called %d times, expected 1", len(f.StringCalls))
	}
}

// StringCalledN returns true if FakeUser.String was called at least n times
func (f *FakeUser) StringCalledN(n int) bool {
	return len(f.StringCalls) >= n
}

// AssertStringCalledN calls t.Error if FakeUser.String was called less than n times
func (f *FakeUser) AssertStringCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.StringCalls) < n {
		t.Errorf("FakeUser.String called %d times, expected >= %d", len(f.StringCalls), n)
	}
}

func (_f2 *FakeUser) ID() (ident6 string) {
	invocation := new(IDInvocation)

	ident6 = _f2.IDHook()

	invocation.Results.Ident6 = ident6

	_f2.IDCalls = append(_f2.IDCalls, invocation)

	return
}

// IDCalled returns true if FakeUser.ID was called
func (f *FakeUser) IDCalled() bool {
	return len(f.IDCalls) != 0
}

// AssertIDCalled calls t.Error if FakeUser.ID was not called
func (f *FakeUser) AssertIDCalled(t *testing.T) {
	t.Helper()
	if len(f.IDCalls) == 0 {
		t.Error("FakeUser.ID not called, expected at least one")
	}
}

// IDNotCalled returns true if FakeUser.ID was not called
func (f *FakeUser) IDNotCalled() bool {
	return len(f.IDCalls) == 0
}

// AssertIDNotCalled calls t.Error if FakeUser.ID was called
func (f *FakeUser) AssertIDNotCalled(t *testing.T) {
	t.Helper()
	if len(f.IDCalls) != 0 {
		t.Error("FakeUser.ID called, expected none")
	}
}

// IDCalledOnce returns true if FakeUser.ID was called exactly once
func (f *FakeUser) IDCalledOnce() bool {
	return len(f.IDCalls) == 1
}

// AssertIDCalledOnce calls t.Error if FakeUser.ID was not called exactly once
func (f *FakeUser) AssertIDCalledOnce(t *testing.T) {
	t.Helper()
	if len(f.IDCalls) != 1 {
		t.Errorf("FakeUser.ID called %d times, expected 1", len(f.IDCalls))
	}
}

// IDCalledN returns true if FakeUser.ID was called at least n times
func (f *FakeUser) IDCalledN(n int) bool {
	return len(f.IDCalls) >= n
}

// AssertIDCalledN calls t.Error if FakeUser.ID was called less than n times
func (f *FakeUser) AssertIDCalledN(t *testing.T, n int) {
	t.Helper()
	if len(f.IDCalls) < n {
		t.Errorf("FakeUser.ID called %d times, expected >= %d", len(f.IDCalls), n)
	}
}