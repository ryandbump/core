package authn

import (
	"github.com/ansel1/merry"

	"github.com/percolate/shisa/context"
	"github.com/percolate/shisa/httpx"
	"github.com/percolate/shisa/models"
)

const (
	AuthnHeaderKey = "Authorization"
)

//go:generate charlatan -output=./authenticator_charlatan.go Authenticator

// Authenticator defines a provder for authenticating the
// principal of a request.
type Authenticator interface {
	// Authenticate extracts a token from the request and
	// resolves it into a user principal.
	Authenticate(context.Context, *httpx.Request) (models.User, merry.Error)
	// Challenge returns the value for the "WWW-Authenticate"
	// header if authentication fails.
	Challenge() string
}
