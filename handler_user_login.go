package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Frevens/chirpy/internal/auth"
	"github.com/Frevens/chirpy/internal/database"
)

type loginResponse struct {
	User
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func (cfg *apiConfig) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
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

	// Generar access token JWT con expiración fija de 1 hora.
	const accessTokenTTL = time.Hour
	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, accessTokenTTL)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	// Generar refresh token, guardarlo en la base de datos con expiración a 60 días.
	refreshToken := auth.MakeRefreshToken()

	const refreshDays = 60
	expiresAt := time.Now().UTC().Add(time.Hour * 24 * refreshDays)

	_, err = cfg.db.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to store refresh token")
		return
	}

	// Responder con la forma requerida, incluyendo el token
	respondWithJSON(w, http.StatusOK, loginResponse{
		User: User{
			ID:          user.ID,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
			Email:       user.Email,
			IsChirpyRed: user.IsChirpyRed,
		},
		Token:        token,
		RefreshToken: refreshToken,
	})
}
