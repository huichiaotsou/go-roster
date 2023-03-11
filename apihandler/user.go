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
	a.router.HandleFunc(userApi, a.handleCreateUser).Methods(http.MethodPost)

	// Apply permission middlewares to sub router
	apiWithID := fmt.Sprintf(userApi + "/{id}")
	permRouter := a.router.PathPrefix(apiWithID).Subrouter()

	// Update user requires user permission
	permRouter.Use(a.mw.CheckUserPerm)
	permRouter.HandleFunc("", a.handleUpdateUser).Methods(http.MethodPut)

	// Delete user requires admin permission
	permRouter.Use(a.mw.CheckAdminPerm)
	permRouter.HandleFunc("", a.handleDeleteUser).Methods(http.MethodDelete)
}

func (a *APIHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body to User struct
	var newUser types.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		err = fmt.Errorf("error while decoding newUser: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the email exists
	exist, err := a.db.VerifyEmailExists(newUser.Email)
	if err != nil {
		err = fmt.Errorf("error while verifying email exists: %s", err)
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
		err = fmt.Errorf("error while creating user: %s", err)
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
		err = fmt.Errorf("error while writing response in handleCreateUser: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *APIHandler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body to User struct
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		err = fmt.Errorf("error while decoding user in handleUpdateUser: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash user password before storing
	hashedPwd := types.HashPassword(user.PwdOrToken)
	user.PwdOrToken = hashedPwd

	// Insert new user into database
	userId, err := a.db.InsertOrUpdateUser(user)
	if err != nil {
		err = fmt.Errorf("error while inserting/updating user in handleUpdateUser: %s", err)
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
		err = fmt.Errorf("error while writing response in handleUpdateUser: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *APIHandler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	// TO-DO

}
