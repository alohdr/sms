package repositories

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/provider"
	"hanoman-id/xendit-payment/internal/apis/operations/template"
	"hanoman-id/xendit-payment/internal/repository/query"
)

type repositories struct {
	qry *query.Queries
}

type Repositories interface {
	CreateTemplate(ctx context.Context, params template.PostTemplateParams, id string) error
	GetListTemplate(ctx context.Context) ([]*query.GetListTemplateRow, error)
	EditTemplate(ctx context.Context, params template.PutTemplateTemplateIDParams) error
	DeleteTemplate(ctx context.Context, params template.DeleteTemplateTemplateIDParams) error

	GetProvider(ctx context.Context) ([]*query.GetProviderRow, error)
	UpdateProvider(ctx context.Context, params provider.PutProviderParams) error
	UpdateAllFalse(ctx context.Context) error

	GetUSer(ctx context.Context, params provider.PutProviderParams) (*query.GetUserRow, error)
}

func NewRepositories(q *query.Queries) Repositories {
	return &repositories{qry: q}
}
