package core

import "github.com/gofrs/uuid/v5"

type Issue struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func NewIssue(title, description string) *Issue {
	return &Issue{
		Title:       title,
		Description: description,
	}
}
