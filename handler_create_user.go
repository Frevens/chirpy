package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Frevens/chirpy/internal/auth"
	"github.com/Frevens/chirpy/internal/database"
	"github.com/google/uuid"
)

type createUserParameters struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var requestData createUserParameters
	if err := decoder.Decode(&requestData); err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to decode JSON")
		return
	}

	hashedPassword, err := auth.HashPassword(requestData.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	user, err := cfg.db.CreateUser(
		r.Context(),
		database.CreateUserParams{
			Email:          requestData.Email,
			HashedPassword: hashedPassword,
		},
	)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusCreated, User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
