package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mj-hagonoy/get-top-ten/handler"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		if _, err := w.Write([]byte(`{"message": "OK"}`)); err != nil {
			log.Printf("w.Write: %s\n", err)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/top10", handler.GetTopTen).Methods(http.MethodPost)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("http.ListenAndServe: %s\n", err.Error())
	}
}
