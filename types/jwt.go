package types

var (
	UserIDclaim    = "userID"
	TeamPermsClaim = "teamPerms"
	Emailclaim     = "email"
)

type TeamPermission struct {
	TeamID         int64  `json:"teamID" db:"team_id"`
	PermissionName string `json:"permissionName" db:"permission_name"`
}

func NewTeamPermission(teamID int64, permissionName string) *TeamPermission {
	return &TeamPermission{
		TeamID:         teamID,
		PermissionName: permissionName,
	}
}
