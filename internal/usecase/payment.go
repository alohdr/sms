package usecase

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/payment"
	"hanoman-id/xendit-payment/internal/models"
)

func (uc *useCase) GetMakingPayment(ctx context.Context, params payment.GetMakingPaymentsParams) (*models.MakingPayment, error) {
	result, err := uc.repo.GetMakingPayment(ctx, params)
	if err != nil {
		return nil, err
	}

	resp := &models.MakingPayment{
		Atm:             result.Atm,
		BankCode:        result.BankCode,
		InternetBanking: result.InternetBanking,
		MobileBanking:   result.MobileBanking,
	}
	return resp, nil
}
