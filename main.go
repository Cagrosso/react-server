package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./web/js"))))
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit '/' endpoint")
	extension, _ := regexp.MatchString("\\.+[a-zA-Z]+", r.URL.EscapedPath())
	// If the url contains an extension, use file server
	if extension {
		http.FileServer(http.Dir("./web/js")).ServeHTTP(w, r)
	} else {
		http.ServeFile(w, r, "./web/index.html")
	}
}
