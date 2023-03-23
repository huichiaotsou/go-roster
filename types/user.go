package types

var (
	USER_ROLE_ADMIN = "admin"
)

type User struct {
	FirstNameEn string `json:"firstNameEn"`
	LastNameEn  string `json:"lastNameEn"`
	FirstNameZh string `json:"firstNameZh"`
	LastNameZh  string `json:"lastNameZh"`
	Email       string `json:"email"`
	PwdOrToken  string `json:"passwordOrToken"`
	DateOfBirth string `json:"dateOfBirth"`
}

func NewUser(
	firstNameEn, lastNameEn, firstNameZh, lastNameZh,
	email, pwdOrToken, dateOfBirth string) *User {
	return &User{
		FirstNameEn: firstNameEn,
		LastNameEn:  lastNameEn,
		FirstNameZh: firstNameZh,
		LastNameZh:  lastNameZh,
		Email:       email,
		PwdOrToken:  pwdOrToken,
		DateOfBirth: dateOfBirth,
	}
}
