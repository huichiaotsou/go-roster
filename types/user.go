package types

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

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

func HashPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("error while hashing password: %s", err)
	}
	return string(hash)
}

func ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
