package usecase

import (
	"context"

	"matcha/internal/core"
)

type (
	CreateIssue struct {
		r core.Repository
	}
	ListIssues struct {
		r core.Repository
	}
)

func NewCreateIssue(r core.Repository) *CreateIssue {
	return &CreateIssue{
		r: r,
	}
}

var _ core.CreateIssueUsecase = (*CreateIssue)(nil)

func (uc *CreateIssue) Execute(ctx context.Context, args *core.CreateIssueParams) (*core.Issue, error) {
	i := core.NewIssue(args.Title, args.Description)
	return uc.r.Issue().Create(ctx, i)
}

var _ core.ListIssuesUsecase = (*ListIssues)(nil)

func NewListIssues(r core.Repository) *ListIssues {
	return &ListIssues{
		r: r,
	}
}

func (uc *ListIssues) Execute(ctx context.Context) ([]*core.Issue, error) {
	return uc.r.Issue().List(ctx)
}
