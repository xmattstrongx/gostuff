package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)
 
func main() {
	mux := goji.NewMux()
	mux.HandleFuncC(pat.Get("/hello/:name"), hello)

	http.ListenAndServe(":8000", mux)
}

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := pat.Param(ctx, "name")
	fmt.Printf("hello %s!", name)
}
