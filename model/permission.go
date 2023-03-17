package model

import (
	"fmt"
	"strings"

	"github.com/huichiaotsou/go-roster/types"
)

func (db *Database) InsertUserPerms(userPerms []types.UserPerms) error {
	placeholders := make([]string, len(userPerms))
	values := make([]interface{}, 0, len(userPerms))

	for i, u := range userPerms {
		ai := i * 3
		placeholders[i] = fmt.Sprintf("($%d, $%d, $%d)", ai+1, ai+2, ai+3)
		values = append(values, u.UserID, u.TeamID, u.PermissionName)
	}

	stmt := fmt.Sprintf(`
		INSERT INTO permissions (user_id, team_id, permission_name) VALUES %s 
		ON CONFLICT (user_id, team_id) DO UPDATE 
		SET permission_name = EXCLUDED.permission_name`,
		strings.Join(placeholders, ", "),
	)

	_, err := db.Sqlx.Exec(stmt, values...)
	if err != nil {
		return fmt.Errorf("error while inserting user perms: %s", err)
	}

	return nil
}
