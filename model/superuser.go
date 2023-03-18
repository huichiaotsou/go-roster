package model

import "fmt"

func (db *Database) UpdateIsSuperuser(userID string, is_super_user bool) error {
	stmt := `UPDATE users SET is_super_user = $1 WHERE id = $2`
	_, err := db.Sqlx.Exec(stmt, is_super_user, userID)
	if err != nil {
		return fmt.Errorf("error while updating is_super_user = %v : %s", is_super_user, err)
	}

	return nil
}
