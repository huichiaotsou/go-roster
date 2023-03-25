package types

var Permission_Names = map[int64]string{
	1: "admin",
	2: "volunteer",
}

var Permission_Values = map[string]int64{
	"admin":     1,
	"volunteer": 2,
}

type PermissionLevel int64

const (
	Admin     PermissionLevel = 1
	Volunteer PermissionLevel = 2
)
