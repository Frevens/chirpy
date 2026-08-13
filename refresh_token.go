package main

import (
	"net/http"
	"time"

	"github.com/Frevens/chirpy/internal/auth"
)

type refreshResponse struct {
	Token string `json:"token"`
}

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	// Extraer refresh token del header Authorization: Bearer <token>
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "missing or invalid authorization header")
		return
	}

	// Buscar usuario asociado y verificar que el token no esté expirado ni revocado.
	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid or expired refresh token")
		return
	}

	// Generar nuevo access token (1 hora)
	accessToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to generate access token")
		return
	}

	// Responder con el nuevo token
	respondWithJSON(w, http.StatusOK, refreshResponse{Token: accessToken})
}
