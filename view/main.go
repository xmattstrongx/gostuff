package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"goji.io"
	"goji.io/pat"
)

//Paydata dsd fucking
type PayData struct {
	Name    string `json:"name"`
	Date    int64  `json:"date"`
	Payrate int    `json:"payrate"`
	Paytime int    `json:"paytime"`
}

func main() {
	mux := goji.NewMux()
	// mux.HandleFuncC(pat.Get("/hello/:name"), hello)
	mux.HandleFuncC(pat.Post("/view"), view)
	// mux.HandleFuncC(pat.Post("/details"), details)

	mux.Use(JSONContentTypeMiddleware)

	http.ListenAndServe("0.0.0.0:8004", mux)

}

func view(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := &PayData{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("FUCK! Could not parse request: %s", err)
		return
	}
	log.Printf("rbody:: %+v", req)

	json.NewEncoder(w).Encode(req)
	return
}

// JSONContentTypeMiddleware forces every response to be JSON-formatted
func JSONContentTypeMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
