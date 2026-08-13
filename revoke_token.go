package main

import (
	"net/http"

	"github.com/Frevens/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	// Extraer refresh token del header Authorization: Bearer <token>
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "missing or invalid authorization header")
		return
	}

	// Marcar como revocado en la base de datos (actualiza revoked_at y updated_at)
	if err := cfg.db.RevokeRefreshToken(r.Context(), refreshToken); err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to revoke refresh token")
		return
	}

	// Responder 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
