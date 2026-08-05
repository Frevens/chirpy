package main

import (
	"encoding/json"
	//"log"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func respondWithJSON(w http.ResponseWriter, status int, payload any) {
	response, err := json.Marshal(payload)

	if err != nil {
		jsonError := errorResponse{
			Error: "Something went wrong",
		}

		response, _ = json.Marshal(jsonError)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, errorResponse{
		Error: message,
	})
}
