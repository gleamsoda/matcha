package repository

import (
	"context"

	"github.com/aarondl/opt/omit"
	"github.com/stephenafamo/bob"

	"matcha/internal/core"
	bobgen "matcha/internal/core/repository/bob/gen"
	sqlcgen "matcha/internal/core/repository/sqlc/gen"
)

type Issue struct {
	sqlc *sqlcgen.Queries
	bob  bob.Executor
}

func NewIssue(db Executor) *Issue {
	return &Issue{
		sqlc: sqlcgen.New(db),
		bob:  bob.New(db),
	}
}

var _ core.IssueRepository = (*Issue)(nil)

func (r *Issue) Create(ctx context.Context, i *core.Issue) (*core.Issue, error) {
	row, err := bobgen.Issues.Insert(ctx, r.bob, &bobgen.IssueSetter{
		Title:       omit.From(i.Title),
		Description: omit.From(i.Description),
	})
	if err != nil {
		return nil, err
	}

	return &core.Issue{
		ID:          row.ID,
		Title:       row.Title,
		Description: row.Description,
	}, err
}

func (r *Issue) List(ctx context.Context) ([]*core.Issue, error) {
	rows, err := bobgen.Issues.Query(ctx, r.bob).All()
	if err != nil {
		return nil, err
	}

	issues := make([]*core.Issue, len(rows))
	for i, row := range rows {
		issues[i] = &core.Issue{
			ID:          row.ID,
			Title:       row.Title,
			Description: row.Description,
		}
	}

	return issues, nil
}
