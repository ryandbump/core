package httpx

import (
	"github.com/ansel1/merry"

	"github.com/percolate/shisa/context"
)

// StringExtractor is a function type that extracts a string from
// the given `context.Context` and `*httpx.Request`.
// An error is returned if the string could not be extracted.
type StringExtractor func(context.Context, *Request) (string, merry.Error)

func (h StringExtractor) InvokeSafely(ctx context.Context, request *Request) (str string, err merry.Error) {
	defer func() {
		arg := recover()
		if arg == nil {
			return
		}

		if e1, ok := arg.(error); ok {
			err = merry.WithMessage(e1, "panic in extractor")
			return
		}

		err = merry.New("panic in extractor").WithValue("context", arg)
	}()

	return h(ctx, request)
}