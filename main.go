package main

import (
	"log"
	"net/http"

	"github.com/mj-hagonoy/get-top-ten/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		w.Write([]byte(`{"message": "OK"}`))
	})

	mux.HandleFunc("/top10", handler.GetTopTen)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("http.ListenAndServe: %s", err.Error())
	}
}
