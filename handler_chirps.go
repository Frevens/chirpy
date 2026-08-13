package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Frevens/chirpy/internal/auth"
	"github.com/Frevens/chirpy/internal/database"
	"github.com/google/uuid"
)

type createChirpRequest struct {
	Body string `json:"body"`
}

type chirpResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerCreateChirp(w http.ResponseWriter, r *http.Request) {

	var params createChirpRequest

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(
			w,
			http.StatusBadRequest,
			"Invalid JSON",
		)
		return
	}

	if len(params.Body) > 140 {
		respondWithError(
			w,
			http.StatusBadRequest,
			"Chirp is too long",
		)
		return
	}

	cleanedBody := cleanChirp(params.Body)

	// Autenticación: extraer bearer token y validar JWT

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "missing or invalid authorization header")
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid token")
		return
	}

	chirp, err := cfg.db.CreateChirp(
		r.Context(),
		database.CreateChirpParams{
			Body:   cleanedBody,
			UserID: userID,
		},
	)

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Could not create chirp",
		)
		return
	}

	response := chirpResponse{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}

	respondWithJSON(
		w,
		http.StatusCreated,
		response,
	)
}
