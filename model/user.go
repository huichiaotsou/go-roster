package model

import (
	"database/sql"

	"github.com/huichiaotsou/go-roster/types"
)

func (m *Model) VerifyEmailExists(email string) (bool, error) {
	var count int
	err := m.Db.Get(&count, "SELECT COUNT(*) FROM users WHERE email = $1", email)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// InsertOrUpdateUser inserts or updates the provided user in the database
// If the user already exists (based on email), it updates the existing user
// Returns the ID of the inserted or updated user, or an error if something goes wrong
func (m *Model) InsertOrUpdateUser(user types.User) (int64, error) {
	// Define the SQL statement to insert a user and handle conflicts on email
	query := `
        INSERT INTO users (
			first_name_en, last_name_en, first_name_zh, last_name_zh, email, pwd_hash_or_token, date_of_birth, create_date
		)
        VALUES (
			:first_name_en, :last_name_en, :first_name_zh, :last_name_zh, :email, :pwd_hash_or_token, :date_of_birth, NOW()
		)
        ON CONFLICT (email) DO UPDATE SET
            first_name_en = EXCLUDED.first_name_en,
            last_name_en = EXCLUDED.last_name_en,
            first_name_zh = EXCLUDED.first_name_zh,
            last_name_zh = EXCLUDED.last_name_zh,
            pwd_hash_or_token = EXCLUDED.pwd_hash_or_token,
            date_of_birth = EXCLUDED.date_of_birth
        RETURNING id;
    `

	// Prepare the statement
	stmt, err := m.Db.PrepareNamed(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// Execute the statement and get the generated user ID
	var userID int64
	err = stmt.Get(&userID, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}

	return userID, nil
}

type Permission struct {
	UserID         string `db:"user_id"`
	TeamID         string `db:"team_id"`
	PermissionName string `db:"permission_name"`
}

func (m *Model) GetPermissionsByUserID(userID string) ([]Permission, error) {
	var permissions []Permission
	query := `SELECT team_id, permission_name FROM permissions WHERE user_id=$1`
	err := m.Db.Select(&permissions, query, userID)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}
