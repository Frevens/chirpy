package main

import (
	"log"
	"net/http"
)
func readiness(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Write the status code (e.g., 200 OK)
	w.WriteHeader(http.StatusOK)

	// Write the response body
	w.Write([]byte("OK"))
}
func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("."))
	mux.Handle("/app/", http.StripPrefix("/app/", fs))
	mux.HandleFunc("/healthz/", readiness)


	
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}