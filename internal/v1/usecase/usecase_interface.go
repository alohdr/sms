package usecase

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/provider"
	"hanoman-id/xendit-payment/internal/apis/operations/template"
	"hanoman-id/xendit-payment/internal/models"
	"hanoman-id/xendit-payment/internal/v1/repositories"
)

type useCase struct {
	repo repositories.Repositories
}

type UseCase interface {
	CreateTemplate(ctx context.Context, params template.PostTemplateParams) (*string, error)
	GetListTemplate(ctx context.Context) (*models.GetTemplate, error)
	EditTemplate(ctx context.Context, params template.PutTemplateTemplateIDParams) (*string, error)
	DeleteTemplate(ctx context.Context, params template.DeleteTemplateTemplateIDParams) (*string, error)

	UpdateProvider(ctx context.Context, params provider.PutProviderProviderIDParams) (*string, error)
	GetProvider(ctx context.Context) (*provider.GetProviderOKBodyResponseData, error)
}

func NewUseCase(r repositories.Repositories) UseCase {
	return &useCase{repo: r}
}
