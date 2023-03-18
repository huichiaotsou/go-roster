package model

import (
	"fmt"
	"strings"

	"github.com/huichiaotsou/go-roster/types"
	"github.com/lib/pq"
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

	stmt := fmt.Sprintf(
		"INSERT INTO teams (team_name) VALUES %s ON CONFLICT DO NOTHING",
		strings.Join(placeholders, ", "),
	)
	_, err := db.Sqlx.Exec(stmt, values...)
	if err != nil {
		return fmt.Errorf("error while inserting teams: %s", err)
	}

	return nil
}

func (db *Database) InsertCampus(c types.Campuses) error {
	// Prepare the query string with placeholders for each team name.
	// For example: "INSERT INTO campus (campus_name) VALUES ($1), ($2), ($3)"
	placeholders := make([]string, 0, len(c.CampusNames))
	values := make([]interface{}, 0, len(c.CampusNames))
	for i, cn := range c.CampusNames {
		placeholders = append(placeholders, fmt.Sprintf("($%d)", i+1))
		values = append(values, cn)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO campus (campus_name) VALUES %s ON CONFLICT DO NOTHING",
		strings.Join(placeholders, ", "),
	)
	_, err := db.Sqlx.Exec(stmt, values...)
	if err != nil {
		return fmt.Errorf("error while inserting campus: %s", err)
	}

	return nil
}

func (db *Database) InsertUserTeams(userTeams types.UserTeams) error {
	// Prepare the SQL statement with placeholders for the user ID and team IDs.
	// We use the unnest() function to convert the team ID array into rows.
	stmt := "INSERT INTO user_teams (user_id, team_id) " +
		"SELECT $1, unnest($2::int[]) " +
		"ON CONFLICT (user_id, team_id) DO NOTHING"

	// Execute the SQL statement with the user ID and team IDs.
	_, err := db.Sqlx.Exec(stmt, userTeams.UserID, pq.Array(userTeams.TeamIDs))
	if err != nil {
		return err
	}

	return nil
}
