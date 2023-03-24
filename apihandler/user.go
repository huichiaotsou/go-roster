package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/huichiaotsou/go-roster/config"
	"github.com/huichiaotsou/go-roster/types"
	"github.com/huichiaotsou/go-roster/utils"
)

var (
	USER_ROUTE = "/user"
)

// /api/v1/user or  /api/v1/user/{id}
func (a *APIHandler) SetUserRoutes() {
	apiVersion := fmt.Sprintf("/api/%s", config.GetApiVersion())
	userAPI := apiVersion + USER_ROUTE

	// Handle create user
	a.router.HandleFunc(userAPI, a.handleCreateUser).Methods(http.MethodPost)

	// Apply permission middlewares to sub router
	apiWithID := fmt.Sprintf(userAPI + "/{id}")

	// Update user requires user permission
	userPermRouter := a.router.PathPrefix(apiWithID).Subrouter()
	userPermRouter.Use(a.mw.CheckUserPerm)
	userPermRouter.HandleFunc("", a.handleUpdateUser).Methods(http.MethodPut)

	// Delete user requires admin permission
	adminPermRouter := a.router.PathPrefix(apiWithID).Subrouter()
	adminPermRouter.Use(a.mw.CheckAdminOrSuperuserPerm)
	adminPermRouter.HandleFunc("", a.handleDeleteUser).Methods(http.MethodDelete)
}

func (a *APIHandler) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body to User struct
	var newUser types.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		handleError(w, err, "error while decoding newUser", http.StatusBadRequest)
		return
	}

	// Check if the email exists
	exist, err := a.db.VerifyEmailExists(newUser.Email)
	if err != nil {
		handleError(w, err, "error while verifying email exists", http.StatusInternalServerError)
		return
	}

	if exist {
		handleError(w, err, "email exists", http.StatusConflict)
		return
	}

	// Hash user password before storing
	hashedPwd := utils.HashPassword(newUser.PwdOrToken)
	newUser.PwdOrToken = hashedPwd

	// Insert new user into database
	userID, err := a.db.InsertOrUpdateUser(newUser)
	if err != nil {
		handleError(w, err, "error while creating user", http.StatusInternalServerError)
		return
	}

	// // Get teamIDs with userID
	// teamPerms, err := a.db.GetTeamPermsByUserID(userID)
	// if err != nil {
	// 	handleError(w, err, "error while getting user team ID", http.StatusInternalServerError)
	// 	return
	// }

	token, err := utils.GenerateJWTToken(userID, nil, newUser.Email)
	if err != nil {
		handleError(w, err, "error while generating JWT token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Authorization", "Bearer "+token)
	respondWithJSON(w, http.StatusCreated, map[string]string{
		"message": "User created successfully",
	})
}

func (a *APIHandler) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body to User struct
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handleError(w, err, "error while decoding user in handleUpdateUser", http.StatusBadRequest)
		return
	}

	// Verify if the req.body email has been modified
	claims, _ := utils.VerifyJWTToken(r)
	if claims[types.Emailclaim].(string) != user.Email {
		handleError(w, fmt.Errorf("email has been modified"), "Forbidden", http.StatusForbidden)
		return
	}

	// Hash user password before storing
	hashedPwd := utils.HashPassword(user.PwdOrToken)
	user.PwdOrToken = hashedPwd

	// Insert new user into database
	userID, err := a.db.InsertOrUpdateUser(user)
	if err != nil {
		handleError(w, err, "error while inserting/updating user in handleUpdateUser", http.StatusInternalServerError)
		return
	}

	// Get teamPermss with userID
	teamPerms, err := a.db.GetTeamPermsByUserID(userID)
	if err != nil {
		handleError(w, err, "error while getting user team ID", http.StatusInternalServerError)
		return
	}

	token, err := utils.GenerateJWTToken(userID, teamPerms, user.Email)
	if err != nil {
		handleError(w, err, "error while generating JWT token", http.StatusInternalServerError)
		return
	}

	// Set token in Authorization header
	w.Header().Set("Authorization", "Bearer "+token)
	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "User updated successfully",
	})
}

func (a *APIHandler) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	// TO-DO

}
