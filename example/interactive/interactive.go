package main

import (
	"encoding/json"
	"net/http"

	"github.com/deref/pier"
)

var calls = 0

func main() {
	pier.Interactive(http.HandlerFunc(handle))
}

func handle(w http.ResponseWriter, req *http.Request) {
	calls++
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"calls": calls,
	})
}
