package pier

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/nojima/httpie-go"
)

func Main(handler http.Handler) {
	peer := &Peer{
		Handler: handler,
	}
	peer.Main()
}

type Peer struct {
	Handler     http.Handler
	BaseContext func(*http.Request) context.Context
}

func (p *Peer) RoundTrip(req *http.Request) (*http.Response, error) {
	recorder := httptest.NewRecorder()
	if p.BaseContext != nil {
		req = req.WithContext(p.BaseContext(req))
	}
	p.Handler.ServeHTTP(recorder, req)
	res := recorder.Result()
	return res, nil
}

func (p *Peer) Main() {
	options := &httpie.Options{
		Transport: p,
	}
	if err := httpie.Main(options); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}
