package types

type Teams struct {
	TeamNames []string `json:"teams"`
}

type UserTeams struct {
	UserID  int64   `json:"userID"`
	TeamIDs []int64 `json:"teamIDs"`
}

// type UserTeams struct {
// 	UserID  int64   `json:"userID"`
// 	TeamIDs []int64 `json:"teamIDs"`
// }
