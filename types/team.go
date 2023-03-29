package types

type Teams struct {
	TeamNames []string `json:"teamNames"`
}

type Campuses struct {
	CampusNames []string `json:"campusNames"`
}

type UserTeams struct {
	UserID  int64   `json:"userID"`
	TeamIDs []int64 `json:"teamIDs"`
}

type Permissions struct {
	PermissionNames []string `json:"permissionNames"`
}

type UserTeamPerm struct {
	UserID       int64 `json:"userID"`
	TeamID       int64 `json:"teamID"`
	PermissionID int64 `json:"permissionID"`
}

type TeamPermission struct {
	TeamID       int64 `json:"teamID"`
	PermissionID int64 `json:"permissionID"`
}

func NewTeamPermission(teamID int64, permissionID int64) *TeamPermission {
	return &TeamPermission{
		TeamID:       teamID,
		PermissionID: permissionID,
	}
}
