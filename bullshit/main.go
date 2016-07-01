package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(ctx, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

// JSONContentTypeMiddleware forces every response to be JSON-formatted
func JSONContentTypeMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/hello/:name"), hello)

	// mux.Use(JSONContentTypeMiddleware)

	http.ListenAndServe(":8000", mux)
}
