package repositories

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/payment"
	"hanoman-id/xendit-payment/internal/repository/query"
)

type repositories struct {
	qry *query.Queries
}

type Repositories interface {
	GetMakingPayment(ctx context.Context, params payment.GetMakingPaymentsParams) (*query.GetMakingPaymentRow, error)
}

func NewRepositories(q *query.Queries) Repositories {
	return &repositories{qry: q}
}
