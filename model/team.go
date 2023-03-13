package model

import (
	"fmt"
	"strings"

	"github.com/huichiaotsou/go-roster/types"
)

func (db *Database) InsertTeams(t types.Teams) error {
	// Prepare the query string with placeholders for each team name.
	// For example: "INSERT INTO teams (team_name) VALUES ($1), ($2), ($3)"
	placeholders := make([]string, 0, len(t.TeamNames))
	values := make([]interface{}, 0, len(t.TeamNames))
	for i, team := range t.TeamNames {
		placeholders = append(placeholders, fmt.Sprintf("($%d)", i+1))
		values = append(values, team)
	}

	stmt := fmt.Sprintf("INSERT INTO teams (team_name) VALUES %s", strings.Join(placeholders, ", "))
	_, err := db.Sqlx.Exec(stmt, values...)
	if err != nil {
		return fmt.Errorf("error while inserting teams: %s", err)
	}

	return nil
}
