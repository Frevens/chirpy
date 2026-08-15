package main

import (
	"log"
	"net/http"
	"sort"

	"github.com/Frevens/chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	authorIDStr := r.URL.Query().Get("author_id")
	sortOrder := r.URL.Query().Get("sort")
	if sortOrder == "" {
		sortOrder = "asc"
	}

	var (
		chirps []database.Chirp
		err    error
	)

	if authorIDStr == "" {
		chirps, err = cfg.db.GetAllChirps(r.Context())
	} else {
		authorID, parseErr := uuid.Parse(authorIDStr)
		if parseErr != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID")
			return
		}
		log.Printf("AUTHOR ID FILTER: %s", authorID)
		chirps, err = cfg.db.GetChirpsByAuthor(r.Context(), authorID)
	}

	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			"Failed to get chirps",
		)
		return
	}

	sort.Slice(chirps, func(i, j int) bool {
		if sortOrder == "desc" {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		}
		return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
	})

	resp := make([]chirpResponse, 0, len(chirps))
	for _, chirp := range chirps {
		resp = append(resp, chirpResponse{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}
	respondWithJSON(w, http.StatusOK, resp)
}
