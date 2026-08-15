package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAPIKey = errors.New("no api key found in headers")

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAPIKey
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "apikey" {
		return "", ErrNoAPIKey
	}

	key := strings.TrimSpace(parts[1])
	if key == "" {
		return "", ErrNoAPIKey
	}

	return key, nil
}
