package core

import "github.com/google/uuid"

type User struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"-"`
	Email          string `json:"email"`
}

func NewUser(username, email, hashedPassword string) *User {
	return &User{
		ID:             uuid.NewString(),
		Username:       username,
		HashedPassword: hashedPassword,
		Email:          email,
	}
}
