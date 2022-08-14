package handler

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/sms"
	"hanoman-id/xendit-payment/pkg/utils"
)

func (h *handler) CreateSms(ctx context.Context, params sms.PostSmsParams) (*sms.PostSmsOKBody, error) {
	_, err := h.useCase.CreateSms(ctx, params)
	if err != nil {
		return nil, err
	}
	res := &sms.PostSmsOKBody{
		ResponseCode:    utils.StatusCodeSuccess,
		ResponseData:    nil,
		ResponseMessege: utils.SuccessCreateTemplate,
	}
	return res, nil
}

func (h *handler) GetSmsHistory(ctx context.Context) (*sms.GetSmsHistoryOKBody, error) {
	data, err := h.useCase.GetSmsHistory(ctx)
	if err != nil {
		return nil, err
	}
	res := &sms.GetSmsHistoryOKBody{
		ResponseCode:    utils.StatusCodeSuccess,
		ResponseData:    *data,
		ResponseMessege: utils.SuccessResponseMessege,
	}
	return res, nil
}
