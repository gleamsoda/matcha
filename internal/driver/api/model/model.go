// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateIssue struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
