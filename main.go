package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"
		
	"github.com/Frevens/chirpy/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db *database.Queries
	platform       string
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}	
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}	

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	mux := http.NewServeMux()
	cfg := &apiConfig{
		db: dbQueries,
		platform: platform,
	}

	fs := http.FileServer(http.Dir("."))

	mux.Handle("/app/", cfg.middlewareMetricsInc(http.StripPrefix("/app/", fs)))

	mux.HandleFunc("GET /api/healthz", readiness)
	mux.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", cfg.handlerReset)
	mux.HandleFunc("POST /api/validate_chirp", cfg.handlerValidateChirp)
	mux.HandleFunc("POST /api/users", cfg.handlerCreateUser)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}