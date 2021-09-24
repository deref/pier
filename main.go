package pier

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/nojima/httpie-go"
)

func Main(handler http.Handler) {
	options := &httpie.Options{
		Transport: &Transport{
			Handler: handler,
		},
	}
	if err := httpie.MainWithOptions(options); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

func Interactive(handler http.Handler) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		args := strings.Split(scanner.Text(), " ")
		os.Args = append([]string{os.Args[0]}, args...)
		options := &httpie.Options{
			Transport: &Transport{
				Handler: handler,
			},
		}
		if err := httpie.MainWithOptions(options); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		}
		fmt.Printf("> ")
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
