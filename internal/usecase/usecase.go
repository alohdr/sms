package usecase

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/payment"
	"hanoman-id/xendit-payment/internal/models"
	"hanoman-id/xendit-payment/internal/repositories"
)

type useCase struct {
	repo repositories.Repositories
}

type UseCase interface {
	GetMakingPayment(ctx context.Context, params payment.GetMakingPaymentsParams) (*models.MakingPayment, error)
}

func NewUseCase(r repositories.Repositories) UseCase {
	return &useCase{repo: r}
}
