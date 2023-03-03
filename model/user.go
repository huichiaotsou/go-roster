package model

type User struct {
}

func (m *Model) GetUser(id int) *User {
	return &User{}
}
