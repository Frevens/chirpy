package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Frevens/chirpy/internal/auth"
)

type loginResponse struct {
	User
	Token string `json:"token"`
}

func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Email            string `json:"email"`
		Password         string `json:"password"`
		ExpiresInSeconds int    `json:"expires_in_seconds,omitempty"` // opcional
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to decode JSON")
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), requestData.Email)
	if err != nil {
		respondWithError(
			w,
			http.StatusUnauthorized,
			"Incorrect email or password",
		)
		return
	}

	match, err := auth.CheckPasswordHash(
		requestData.Password,
		user.HashedPassword,
	)
	if err != nil || !match {
		respondWithError(
			w,
			http.StatusUnauthorized,
			"Incorrect email or password",
		)
		return
	}

	// Calcular TTL: por defecto 1 hora, si se especifica usarlo, pero nunca > 1 hora.
	const maxSeconds = 3600
	expires := requestData.ExpiresInSeconds
	if expires <= 0 {
		expires = maxSeconds
	}
	if expires > maxSeconds {
		expires = maxSeconds
	}

	// Generar token (la función en el paquete auth debe aceptar id y duración).
	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Duration(expires)*time.Second)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Responder con la forma requerida, incluyendo el token
	respondWithJSON(w, http.StatusOK, loginResponse{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
		Token: token,
	})
}
