package pier

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/brandonbloom/httpie-go"
)

func Main(handler http.Handler) {
	peer := &Peer{
		Handler: handler,
	}
	peer.Main()
}

// Peer configures Pier's in-process HTTP exchange.
type Peer struct {
	// Handler receives the constructed HTTP request directly, without opening
	// a network socket. Handler must be non-nil.
	Handler http.Handler
	// BaseContext, if non-nil, supplies the base context for each request.
	BaseContext func(*http.Request) context.Context
	// Stdin is used for request body input. If nil, os.Stdin is used.
	Stdin io.Reader
	// Stdout receives response output. If nil, os.Stdout is used.
	Stdout io.Writer
	// Stderr receives usage and informational output. If nil, os.Stderr is used.
	Stderr io.Writer
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
	if err := httpie.Main(p.options()); err != nil {
		stderr := p.Stderr
		if stderr == nil {
			stderr = os.Stderr
		}
		fmt.Fprintf(stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

// Run executes Pier with explicit command-line arguments, including argv[0].
func (p *Peer) Run(argv []string) error {
	return httpie.Run(argv, p.options())
}

func (p *Peer) options() *httpie.Options {
	return &httpie.Options{
		Transport: p,
		Stdin:     p.Stdin,
		Stdout:    p.Stdout,
		Stderr:    p.Stderr,
	}
}
