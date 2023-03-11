package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/types"
)

// /api/v1/user or  /api/v1/user/{id}
func (a *APIHandler) SetUserRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	userApi := apiVersion + "/user"

	// Handle create user
	a.router.HandleFunc(userApi, a.CreateUser).Methods(http.MethodPost)

	// Apply CheckUserPerm middleware to the sub router userPermRouter
	apiWithID := fmt.Sprintf(userApi + "/{id}")
	userPermRouter := a.router.PathPrefix(apiWithID).Subrouter()
	userPermRouter.Use(a.mw.CheckUserPerm)
	userPermRouter.HandleFunc(apiWithID, a.UpdateUser).Methods(http.MethodPut)
	userPermRouter.HandleFunc(apiWithID, a.DeleteUser).Methods(http.MethodDelete)
}

func (a *APIHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body to User struct
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the email exists
	exist, err := a.db.VerifyEmailExists(newUser.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exist {
		http.Error(w, "email exists", http.StatusConflict)
		return
	}

	// Hash user password before storing
	hashedPwd := types.HashPassword(newUser.PwdOrToken)
	newUser.PwdOrToken = hashedPwd

	// Insert new user into database
	userId, err := a.db.InsertOrUpdateUser(newUser)
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Write response
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *APIHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// TO-DO

}

func (a *APIHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// TO-DO

}
