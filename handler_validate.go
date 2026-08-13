package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type cleanResponse struct {
	CleanedBody string `json:"cleaned_body"`
}

type parameters struct {
	Body   string `json:"body"`
	UserID string `json:"user_id"`
}

var listOfBannedWords = []string{
	"kerfuffle",
	"sharbert",
	"fornax",
}

func (cfg *apiConfig) handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var requestData parameters

	if err := decoder.Decode(&requestData); err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to decode JSON")
		return
	}

	if len(requestData.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	cleanedBody := cleanChirp(requestData.Body)

	respondWithJSON(w, http.StatusOK, cleanResponse{
		CleanedBody: cleanedBody,
	})
}

func cleanChirp(body string) string {
	words := strings.Split(body, " ")

	for i, word := range words {
		lowerWord := strings.ToLower(word)

		for _, banned := range listOfBannedWords {
			if lowerWord == banned {
				words[i] = "****"
			}
		}
	}

	return strings.Join(words, " ")
}
