package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type createUserParameters struct {
    Email string `json:"email"`
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

    user, err := cfg.db.CreateUser(r.Context(), requestData.Email)
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
