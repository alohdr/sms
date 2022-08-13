package handler

import (
	"context"
	"hanoman-id/xendit-payment/configs"
	"hanoman-id/xendit-payment/internal/apis/operations/template"
	"hanoman-id/xendit-payment/internal/v1/usecase"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
	CreateTemplate(ctx context.Context, params template.PostTemplateParams) (*template.PostTemplateOKBody, error)
	GetListTemplate(ctx context.Context) (*template.GetTemplateOKBody, error)
	EditTemplate(ctx context.Context, params template.PutTemplateTemplateIDParams) (*template.PutTemplateTemplateIDOKBody, error)
}

func NewHandler() Handlers {
	return &handler{useCase: configs.GetUsecase()}
}
