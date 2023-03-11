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

	// Update user requires user permission
	userPermRouter := a.router.PathPrefix(apiWithID).Subrouter()
	userPermRouter.Use(a.mw.CheckUserPerm)
	userPermRouter.HandleFunc("", a.handleUpdateUser).Methods(http.MethodPut)

	// Delete user requires admin permission
	adminPermRouter := a.router.PathPrefix(apiWithID).Subrouter()
	adminPermRouter.Use(a.mw.CheckAdminPerm)
	adminPermRouter.HandleFunc("", a.handleDeleteUser).Methods(http.MethodDelete)
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

	// Get teamID with userID
	teamID, err := a.db.GetTeamIDByUserID(userId)
	if err != nil {
		err = fmt.Errorf("error while getting user team ID: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := generateJWTToken(userId, teamID, newUser.Email)
	if err != nil {
		err = fmt.Errorf("error while generating JWT token: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set token in Authorization header
	w.Header().Set("Authorization", "Bearer "+token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Return success response
	response := map[string]string{
		"message": "User created successfully",
	}
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

	// Get teamID with userID
	teamID, err := a.db.GetTeamIDByUserID(userId)
	if err != nil {
		err = fmt.Errorf("error while getting user team ID: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := generateJWTToken(userId, teamID, user.Email)
	if err != nil {
		err = fmt.Errorf("error while generating JWT token: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set token in Authorization header
	w.Header().Set("Authorization", "Bearer "+token)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Return success response
	response := map[string]string{
		"message": "User updated successfully",
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		err = fmt.Errorf("error while writing response in handleCreateUser: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *APIHandler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	// TO-DO

}
