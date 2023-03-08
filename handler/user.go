package handler

import (
	"encoding/json"
	"net/http"

	"github.com/huichiaotsou/go-roster/types"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body to User struct
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the email exists
	exist, err := h.Model.VerifyEmailExists(newUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if exist {
		http.Error(w, "email exists", http.StatusConflict)
		return
	}

	// Insert new user into database
	userId, err := h.Model.InsertOrUpdateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response with user ID
	response := struct {
		UserID int64 `json:"userId"`
	}{
		UserID: userId,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TO-DO

}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TO-DO

}
