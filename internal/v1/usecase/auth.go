package usecase

import (
	"context"
	"errors"
	"fmt"
	"hanoman-id/xendit-payment/internal/apis/operations/auth"
	"hanoman-id/xendit-payment/pkg/utils"
	"time"

	"github.com/google/uuid"
)

func (uc *useCase) CreateOtp(ctx context.Context, params auth.PostCreateOtpParams) (*string, error) {
	otp := utils.CreateOtp()
	expDate := time.Now().Add(2 * time.Minute).UTC().Format(time.RFC3339)

	user, err := uc.repo.GetPhoneNumber(ctx, params.Body.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if user.ID != 0 {
		err = uc.repo.UpdateOtp(ctx, params, otp, expDate)
		if err != nil {
			return nil, err
		}
	} else {
		err = uc.repo.CreateOtp(ctx, params, otp, expDate)
		if err != nil {
			return nil, err
		}
	}
	return &otp, nil
}

func (uc *useCase) ValidateOtp(ctx context.Context, params auth.PostValidateOtpParams) (*auth.PostValidateOtpOKBodyData, error) {
	dateNow := time.Now().UTC()

	user, err := uc.repo.GetPhoneNumber(ctx, params.Body.PhoneNumber)
	if err != nil {
		return nil, err
	}

	if params.Body.Otp != user.Otp {
		return nil, errors.New("OTP not valid")
	}

	expired, _ := utils.Expired(dateNow, user.ExpiredDate)
	if expired {
		return nil, errors.New("CODE IS EXPIRED")
	}

	isiData := map[string]string{
		"id":       uuid.New().String(),
		"username": user.PhoneNumber,
		"role":     "ADMIN",
	}

	token, err := utils.GenerateToken(isiData)
	if err != nil {
		fmt.Println("MASUK ERROR")
		return nil, err
	}

	result := &auth.PostValidateOtpOKBodyData{
		AccessToken: *token,
		User:        user.PhoneNumber,
	}

	return result, nil

}
