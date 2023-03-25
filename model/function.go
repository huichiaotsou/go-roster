package model

// func (db *Database) InsertFunctions(f types.Functions) error {
// 	stmt := `INSERT INTO functions (team_id, func_name)
//     SELECT $1, unnest($2::text[])
//     ON CONFLICT DO NOTHING`

// 	_, err := db.Sqlx.Exec(stmt, f.TeamID, pq.Array(f.FuncNames))
// 	if err != nil {
// 		return fmt.Errorf("error while inserting functions: %s", err)
// 	}

// 	return nil
// }

// TO-DO remove team id from functions, to re-writes
