package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	mux.HandleFuncC(pat.Post("/aggregate"), aggregate)
	// mux.HandleFuncC(pat.Post("/details"), details)

	mux.Use(JSONContentTypeMiddleware)

	http.ListenAndServe("0.0.0.0:8003", mux)

}

func aggregate(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := &PayData{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("FUCK! Could not parse request: %s", err)
		return
	}
	log.Printf("rbody:: %+v", req)

	json.NewEncoder(w).Encode(req)
	View(*req)
	return
}

func View(data PayData) {
	fmt.Printf("This is the data in the aggregate: %+v\n", data)

	url := "http://dockerhost:8004/view"
	fmt.Println("URL:>", url)

	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

// JSONContentTypeMiddleware forces every response to be JSON-formatted
func JSONContentTypeMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
