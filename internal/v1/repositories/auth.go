package repositories

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/auth"
	"hanoman-id/xendit-payment/internal/repository/query"
)

func (r *repositories) CreateOtp(ctx context.Context, params auth.PostCreateOtpParams, otp string, date string) error {
	data := query.CreateOtpParams{
		PhoneNumber: params.Body.PhoneNumber,
		Otp:         otp,
		ExpiredDate: date,
	}

	_, err := r.qry.CreateOtp(ctx, &data)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositories) UpdateOtp(ctx context.Context, params auth.PostCreateOtpParams, otp string, date string) error {
	data := query.UpdateOtpParams{
		PhoneNumber: params.Body.PhoneNumber,
		Otp:         otp,
		ExpiredDate: date,
	}

	_, err := r.qry.UpdateOtp(ctx, &data)
	if err != nil {
		return err
	}
	return nil
}

func (r *repositories) GetPhoneNumber(ctx context.Context, params string) (*query.Otp, error) {
	res, _ := r.qry.GetPhoneNumber(ctx, params)
	return res, nil
}
