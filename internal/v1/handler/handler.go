package handler

import (
	"context"
	"hanoman-id/xendit-payment/configs"
	"hanoman-id/xendit-payment/internal/apis/operations/auth"
	"hanoman-id/xendit-payment/internal/apis/operations/health"
	"hanoman-id/xendit-payment/internal/v1/usecase"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
	CreateOtp(ctx context.Context, params auth.PostCreateOtpParams) (*auth.PostCreateOtpOKBodyData, error)
	ValidateOtp(ctx context.Context, params auth.PostValidateOtpParams) (*auth.PostValidateOtpOKBodyData, error)
	GetWelcome(ctx context.Context, params health.GetWelcomeParams, principal interface{}) (health.GetWelcomeOKBody, error)
	ReadirectToken(ctx context.Context, params health.GetRedirectParams) (health.GetRedirectOKBody, error)
}

func NewHandler() Handlers {
	return &handler{useCase: configs.GetUsecase()}
}
