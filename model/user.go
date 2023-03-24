package model

import (
	"database/sql"
	"fmt"
	"strconv"

	dbtypes "github.com/huichiaotsou/go-roster/model/types"
	"github.com/huichiaotsou/go-roster/types"
)

func (db *Database) VerifyEmailExists(email string) (bool, error) {
	var count int
	err := db.Sqlx.Get(&count, "SELECT COUNT(*) FROM users WHERE email = $1", email)
	if err != nil {
		return false, fmt.Errorf("error while verifying email exists: %s", err)
	}
	return count > 0, nil
}

// InsertOrUpdateUser inserts or updates the provided user in the database
// If the user already exists (based on email), it updates the existing user
// Returns the ID of the inserted or updated user, or an error if something goes wrong
func (db *Database) InsertOrUpdateUser(user types.User) (int64, error) {
	// Define the SQL statement to insert a user and handle conflicts on email
	query := `
        INSERT INTO users (
            first_name_en, last_name_en, first_name_zh, last_name_zh, email, pwd_hash_or_token, date_of_birth, created_date
        )
        VALUES (
            $1, $2, $3, $4, $5, $6, $7, NOW()
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

	// Execute the statement and get the generated user ID
	var userID int64
	err := db.Sqlx.Get(&userID, query, user.FirstNameEn, user.LastNameEn, user.FirstNameZh, user.LastNameZh, user.Email, user.PwdOrToken, user.DateOfBirth)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, fmt.Errorf("error while inserting/updating user: %s", err)
	}

	return userID, nil
}

func (db *Database) GetTeamPermsByUserID(userID int64) ([]*types.TeamPermission, error) {
	// stmt := "SELECT * FROM permissions WHERE user_id=$1"
	var teamPerms []dbtypes.DbPermission
	err := db.Sqlx.Select(&teamPerms, "SELECT * FROM user_teams_perms WHERE user_id = $1", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var teamPermsType []*types.TeamPermission
	for _, t := range teamPerms {
		teamID, _ := strconv.ParseInt(t.TeamID, 10, 64)
		permID, _ := strconv.ParseInt(t.PermissionID, 10, 64)
		teamPermsType = append(teamPermsType, types.NewTeamPermission(teamID, permID))
	}
	return teamPermsType, nil
}

func (db *Database) IsSuperUser(userID int64) (bool, error) {
	var user dbtypes.DbUser
	query := "SELECT id, is_super_user FROM users WHERE id= $1"
	err := db.Sqlx.Get(&user, query, userID)
	if err != nil {
		return false, fmt.Errorf("error while verifying super user by user ID: %s", err)
	}

	return user.IsSuperUser || userID == 1, nil
}
