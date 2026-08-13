package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoBearerToken = errors.New("no bearer token found in headers")

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoBearerToken
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", ErrNoBearerToken
	}

	token := strings.TrimSpace(parts[1])
	if token == "" {
		return "", ErrNoBearerToken
	}

	return token, nil
}
