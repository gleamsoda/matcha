package core

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

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

type (
	CreateIssueUsecase interface {
		Execute(ctx context.Context, args *CreateIssueParams) (*Issue, error)
	}
	CreateIssueParams struct {
		Title       string
		Description string
	}
	ListIssuesUsecase interface {
		Execute(ctx context.Context) ([]*Issue, error)
	}
	IssueRepository interface {
		Create(ctx context.Context, u *Issue) (*Issue, error)
		List(ctx context.Context) ([]*Issue, error)
	}
)
