package repositories

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/auth"
	"hanoman-id/xendit-payment/internal/repository/query"
)

type repositories struct {
	qry *query.Queries
}

type Repositories interface {
	CreateOtp(ctx context.Context, params auth.PostCreateOtpParams, otp string, date string) error
	GetPhoneNumber(ctx context.Context, params string) (*query.Otp, error)
	UpdateOtp(ctx context.Context, params auth.PostCreateOtpParams, otp string, date string) error
}

func NewRepositories(q *query.Queries) Repositories {
	return &repositories{qry: q}
}
