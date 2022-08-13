package usecase

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/auth"
	"hanoman-id/xendit-payment/internal/v1/repositories"
)

type useCase struct {
	repo repositories.Repositories
}

type UseCase interface {
	CreateOtp(ctx context.Context, params auth.PostCreateOtpParams) (*string, error)
	ValidateOtp(ctx context.Context, params auth.PostValidateOtpParams) (*auth.PostValidateOtpOKBodyData, error)
}

func NewUseCase(r repositories.Repositories) UseCase {
	return &useCase{repo: r}
}
