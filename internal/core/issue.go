package core

import "github.com/google/uuid"

type Issue struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewIssue(title, description string) *Issue {
	return &Issue{
		ID:          uuid.NewString(),
		Title:       title,
		Description: description,
	}
}
