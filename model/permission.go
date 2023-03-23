package model

import (
	"fmt"
	"strings"

	"github.com/huichiaotsou/go-roster/types"
)

func (db *Database) InsertPerms(p types.Permissions) error {
	// Prepare the query string with placeholders for each team name.
	// For example: "INSERT INTO teams (team_name) VALUES ($1), ($2), ($3)"
	placeholders := make([]string, 0, len(p.PermissionNames))
	values := make([]interface{}, 0, len(p.PermissionNames))
	for i, team := range p.PermissionNames {
		placeholders = append(placeholders, fmt.Sprintf("($%d)", i+1))
		values = append(values, team)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO perms (permission_name) VALUES %s ON CONFLICT DO NOTHING",
		strings.Join(placeholders, ", "),
	)
	_, err := db.Sqlx.Exec(stmt, values...)
	if err != nil {
		return fmt.Errorf("error while inserting perms: %s", err)
	}

	return nil
}

func (db *Database) InsertUserTeamPerms(userPerms []types.UserTeamPerm) error {
	placeholders := make([]string, len(userPerms))
	values := make([]interface{}, 0, len(userPerms))

	for i, u := range userPerms {
		ai := i * 3
		placeholders[i] = fmt.Sprintf("($%d, $%d, $%d)", ai+1, ai+2, ai+3)
		values = append(values, u.UserID, u.TeamID, u.PermissionID)
	}

	stmt := fmt.Sprintf(`
		INSERT INTO user_teams_perms (user_id, team_id, perm_id) VALUES %s 
		ON CONFLICT DO NOTHING`,
		strings.Join(placeholders, ", "),
	)

	_, err := db.Sqlx.Exec(stmt, values...)
	if err != nil {
		return fmt.Errorf("error while inserting user perms: %s", err)
	}

	return nil
}
