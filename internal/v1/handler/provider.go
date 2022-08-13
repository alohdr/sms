package handler

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/provider"
	"hanoman-id/xendit-payment/pkg/utils"
)

func (h *handler) UpdateProvider(ctx context.Context, params provider.PutProviderProviderIDParams) (*provider.PutProviderProviderIDOKBody, error) {
	_, err := h.useCase.UpdateProvider(ctx, params)
	if err != nil {
		return nil, err
	}

	resp := &provider.PutProviderProviderIDOKBody{
		ResponseCode:    utils.StatusCodeSuccess,
		ResponseData:    "",
		ResponseMessege: utils.SuccessUpdateProvider,
	}

	return resp, nil

}

func (h *handler) GetProvider(ctx context.Context) (*provider.GetProviderOKBody, error) {
	data, err := h.useCase.GetProvider(ctx)
	if err != nil {
		return nil, err
	}

	resp := &provider.GetProviderOKBody{
		ResponseCode:    utils.StatusCodeSuccess,
		ResponseData:    data,
		ResponseMessege: utils.SuccessResponseMessege,
	}

	return resp, nil

}
