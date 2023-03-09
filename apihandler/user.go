package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/types"
)

func (h *APIHandler) SetUserRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	userApi := apiVersion + "/user"

	// create user
	h.Router.HandleFunc(userApi, h.CreateUser).Methods("POST")

	// apply CheckUserPerm middleware to the sub router userPermRouter
	apiWithID := fmt.Sprintf(userApi + "/{id}")
	userPermRouter := h.Router.PathPrefix(apiWithID).Subrouter()
	userPermRouter.Use(h.Middleware.CheckUserPerm)

	// modify & delete user with user ID
	userPermRouter.HandleFunc(apiWithID, h.UpdateUser).Methods("PUT")
	userPermRouter.HandleFunc(apiWithID, h.DeleteUser).Methods("DELETE")
}

func (h *APIHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body to User struct
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the email exists
	exist, err := h.DB.VerifyEmailExists(newUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if exist {
		http.Error(w, "email exists", http.StatusConflict)
		return
	}

	// Insert new user into database
	userId, err := h.DB.InsertOrUpdateUser(newUser)
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

func (h *APIHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TO-DO

}

func (h *APIHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TO-DO

}
