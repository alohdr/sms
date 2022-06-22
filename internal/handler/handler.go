package handler

import (
	"context"
	"hanoman-id/xendit-payment/configs"
	"hanoman-id/xendit-payment/internal/apis/operations/payment"
	"hanoman-id/xendit-payment/internal/models"
	"hanoman-id/xendit-payment/internal/usecase"
)

type handler struct {
	useCase usecase.UseCase
}

type Handlers interface {
	GetMakingPayment(ctx context.Context, params payment.GetMakingPaymentsParams) (*models.MakingPayment, error)
}

func NewHandler() Handlers {
	return &handler{useCase: configs.GetUsecase()}
}
