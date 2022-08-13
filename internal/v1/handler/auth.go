package handler

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/auth"
)

func (h *handler) CreateOtp(ctx context.Context, params auth.PostCreateOtpParams) (*auth.PostCreateOtpOKBodyData, error) {
	res, err := h.useCase.CreateOtp(ctx, params)
	if err != nil {
		return nil, err
	}
	response := &auth.PostCreateOtpOKBodyData{
		Otp: *res,
	}
	return response, nil
}

func (h *handler) ValidateOtp(ctx context.Context, params auth.PostValidateOtpParams) (*auth.PostValidateOtpOKBodyData, error) {
	res, err := h.useCase.ValidateOtp(ctx, params)
	if err != nil {
		return nil, err
	}
	return res, nil
}
