package types

type Teams struct {
	TeamNames []string `json:"teams"`
}

type Campuses struct {
	CampusNames []string `json:"campus"`
}

type UserTeams struct {
	UserID  int64   `json:"userID"`
	TeamIDs []int64 `json:"teamIDs"`
}

type UserPerms struct {
	UserID         int64  `json:"userID"`
	TeamID         int64  `json:"teamID"`
	PermissionName string `json:"permissionName"`
}

type TeamPermission struct {
	TeamID         int64  `json:"teamID"`
	PermissionName string `json:"permissionName"`
}

func NewTeamPermission(teamID int64, permissionName string) *TeamPermission {
	return &TeamPermission{
		TeamID:         teamID,
		PermissionName: permissionName,
	}
}
