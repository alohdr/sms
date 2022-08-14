package handler

import (
	"context"
	"hanoman-id/xendit-payment/configs"
	"hanoman-id/xendit-payment/internal/apis/operations/provider"
	"hanoman-id/xendit-payment/internal/apis/operations/sms"
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
	DeleteTemplate(ctx context.Context, params template.DeleteTemplateTemplateIDParams) (*template.DeleteTemplateTemplateIDOKBody, error)

	UpdateProvider(ctx context.Context, params provider.PutProviderParams) (*provider.PutProviderProviderIDOKBody, error)
	GetProvider(ctx context.Context) (*provider.GetProviderOKBody, error)

	CreateSms(ctx context.Context, params sms.PostSmsParams) (*sms.PostSmsOKBody, error)
	GetSmsHistory(ctx context.Context) (*sms.GetSmsHistoryOKBody, error)
}

func NewHandler() Handlers {
	return &handler{useCase: configs.GetUsecase()}
}
