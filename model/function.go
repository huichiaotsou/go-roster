package model

import (
	"fmt"
	"strings"

	"github.com/huichiaotsou/go-roster/types"
)

func (db *Database) InsertFunctions(f types.Functions) error {
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
		return fmt.Errorf("error while inserting functions: %s", err)
	}

	return nil
}
