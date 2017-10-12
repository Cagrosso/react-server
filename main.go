package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

var (
	// Path to the Project
	ProjectPath string
)

func init() {
	ProjectPath = os.Getenv("PROJECT_PATH")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println("Hit '/' endpoint")
	http.FileServer(http.Dir(filepath.Join(ProjectPath, "web", "index.html")))
}
