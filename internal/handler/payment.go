package handler

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/payment"
	"hanoman-id/xendit-payment/internal/models"
)

func (h *handler) GetMakingPayment(ctx context.Context, params payment.GetMakingPaymentsParams) (*models.MakingPayment, error) {
	res, err := h.useCase.GetMakingPayment(ctx, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}
