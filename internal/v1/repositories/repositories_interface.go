package repositories

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/template"
	"hanoman-id/xendit-payment/internal/repository/query"
)

type repositories struct {
	qry *query.Queries
}

type Repositories interface {
	CreateTemplate(ctx context.Context, params template.PostTemplateParams, id string) error
	GetListTemplate(ctx context.Context) ([]*query.GetListTemplateRow, error)
}

func NewRepositories(q *query.Queries) Repositories {
	return &repositories{qry: q}
}
