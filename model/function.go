package model

import (
	"fmt"
	"strings"

	"github.com/huichiaotsou/go-roster/types"
	"github.com/lib/pq"
)

func (db *Database) InsertFunctions(f types.Functions) error {
	// Prepare the query string with placeholders for each team name.
	// For example: "INSERT INTO teams (team_name) VALUES ($1), ($2), ($3)"
	placeholders := make([]string, 0, len(f.FuncNames))
	values := make([]interface{}, 0, len(f.FuncNames))
	for i, f := range f.FuncNames {
		placeholders = append(placeholders, fmt.Sprintf("($%d)", i+1))
		values = append(values, f)
	}

	stmt := fmt.Sprintf(
		"INSERT INTO functions (func_name) VALUES %s ON CONFLICT DO NOTHING",
		strings.Join(placeholders, ", "),
	)
	_, err := db.Sqlx.Exec(stmt, values...)
	if err != nil {
		return fmt.Errorf("error while inserting func_names: %s", err)
	}

	return nil

}

func (db *Database) ClearUserFuncs(userID int64) error {
	stmt := "DELETE FROM user_funcs WHERE user_id = $1"
	_, err := db.Sqlx.Exec(stmt, userID)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) InsertUserFuncs(userID int64, funcIDs []int64) error {
	stmt := "INSERT INTO user_funcs (user_id, func_id) " +
		"SELECT $1, unnest($2::int[]) " +
		"ON CONFLICT (user_id, func_id) DO NOTHING"

	_, err := db.Sqlx.Exec(stmt, userID, pq.Array(funcIDs))
	if err != nil {
		return err
	}

	return nil
}
