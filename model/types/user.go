package types

import "time"

type DbUser struct {
	ID             string    `db:"id"`
	FirstNameEn    string    `db:"first_name_en"`
	LastNameEn     string    `db:"last_name_en"`
	FirstNameZh    string    `db:"first_name_zh"`
	LastNameZh     string    `db:"last_name_zh"`
	Email          string    `db:"email"`
	EmailVerified  bool      `db:"email_verified"`
	PwdHashOrToken string    `db:"pwd_hash_or_token"`
	DateOfBirth    string    `db:"date_of_birth"`
	CreatedDate    time.Time `db:"created_date"`
}

type DbPermission struct {
	UserID       string `db:"user_id"`
	TeamID       string `db:"team_id"`
	PermissionID string `db:"permission_id"`
}
