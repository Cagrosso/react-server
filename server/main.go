package main

import (
	"http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", servePage).Methods("GET")
	http.Handle("/", r)
}

func servePage() {

}
