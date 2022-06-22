package usecase

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/payment"
	"hanoman-id/xendit-payment/internal/models"
)

func (uc *useCase) GetMakingPayment(ctx context.Context, params payment.GetMakingPaymentsParams) (*models.MakingPayment, error) {
	result := models.MakingPayment{}
	data, err := uc.repo.GetMakingPayment(ctx, params)
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		result = append(result, &models.MakingPaymentItems0{
			Description: v.Description,
			Header:      v.Type,
		})
	}

	return &result, nil
}
