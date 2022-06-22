package repositories

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/payment"
	"hanoman-id/xendit-payment/internal/repository/query"
)

func (r *repositories) GetMakingPayment(ctx context.Context, params payment.GetMakingPaymentsParams) (*query.GetMakingPaymentRow, error) {
	result, err := r.qry.GetMakingPayment(ctx, params.BankCode)
	if err != nil {
		return nil, err
	}
	return result, nil
}
