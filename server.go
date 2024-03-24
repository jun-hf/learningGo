package main

import (
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found", http.StatusNotFound)
	})
	server.Handle("/", http.FileServer(http.Dir("./index.html")))
	http.ListenAndServe(":8080", server)
}