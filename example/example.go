package main

import (
	"io/ioutil"
	"net/http"

	"github.com/deref/pier"
)

func main() {
	pier.Main(http.HandlerFunc(handle))
}

func handle(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/echo" {
		w.WriteHeader(404)
		return
	}
	if req.Body != nil {
		bs, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bs)
	}
}
