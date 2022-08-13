package usecase

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/provider"
	"hanoman-id/xendit-payment/pkg/utils"
)

func (uc *useCase) UpdateProvider(ctx context.Context, params provider.PutProviderProviderIDParams) (*string, error) {
	err := uc.repo.UpdateProvider(ctx, params)
	if err != nil {
		return nil, err
	}

	resp := utils.SuccessUpdateProvider

	return &resp, nil

}

func (uc *useCase) GetProvider(ctx context.Context) (*provider.GetProviderOKBodyResponseData, error) {
	data, err := uc.repo.GetProvider(ctx)
	if err != nil {
		return nil, err
	}

	response := &provider.GetProviderOKBodyResponseData{
		ID:   data.ID,
		Name: data.Name,
	}

	return response, nil

}
