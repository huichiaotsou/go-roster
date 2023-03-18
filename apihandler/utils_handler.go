package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleError(w http.ResponseWriter, err error, message string, statusCode int) {
	err = fmt.Errorf("%s: %s", message, err)
	http.Error(w, err.Error(), statusCode)
}
