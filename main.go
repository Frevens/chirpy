package main

import (
	"log"
	"fmt"
	"net/http"
	"sync/atomic"

	
)
type apiConfig struct {
fileserverHits atomic.Int32
}
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Atomically increment the counter by 1
		cfg.fileserverHits.Add(1)
		
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func readiness(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Write the status code (e.g., 200 OK)
	w.WriteHeader(http.StatusOK)

	// Write the response body
	w.Write([]byte("OK"))
}
func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	// Load the current value atomically
	hits := cfg.fileserverHits.Load()
	
	// Set the content type to plain text (optional but recommended)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	// Write the formatted response
	fmt.Fprintf(w, "Hits: %d\n", hits)
}   
func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	// Atomically reset the counter to 0
	cfg.fileserverHits.Store(0)
	
	// Set content type to plain text
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	
	// Confirm the reset in the response
	fmt.Fprintf(w, "Hits: 0\n")
}
func main() {
	mux := http.NewServeMux()
	cfg := &apiConfig{}

	fs := http.FileServer(http.Dir("."))
	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app/", fs)))
	mux.HandleFunc("GET /healthz", readiness)
	mux.HandleFunc("GET /metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /reset", cfg.handlerReset)
	
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}