package pier

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/nojima/httpie-go"
)

func Main(handler http.Handler) {
	options := &httpie.Options{
		Transport: &Transport{
			Handler: handler,
		},
	}
	if err := httpie.Main(options); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

type Transport struct {
	Handler http.Handler
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	recorder := httptest.NewRecorder()
	t.Handler.ServeHTTP(recorder, req)
	res := recorder.Result()
	return res, nil
}
