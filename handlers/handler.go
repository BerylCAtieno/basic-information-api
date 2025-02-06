package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Email          string `json:"email"`
	CurrentDatetime string `json:"current_datetime"`
	GitHubURL      string `json:"github_url"`
}

func GetInfoHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Email:          "berylatieno30@gmail.com",
		CurrentDatetime: time.Now().UTC().Format(time.RFC3339),
		GitHubURL:      "https://github.com/BerylCAtieno/basic-information-api",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
