package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	cards = make([]TimeCard, 0)
)

func main() {
	http.HandleFunc("/api/timeCards", timecards)

	log.Println("Before listening")
	http.ListenAndServe(":8030", nil)
	log.Println("After listening")
}

type TimeCardEntry struct {
	punchIn bool
	time    int64
}

type TimeCard struct {
	entries []TimeCardEntry
}

func timecards(w http.ResponseWriter, r *http.Request) {
	var tc TimeCard
	err := json.NewDecoder(r.Body).Decode(&tc)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	cards = append(cards, tc)
	marshalled, err := json.Marshal(tc)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(marshalled)
}
