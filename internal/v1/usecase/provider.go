package usecase

import (
	"context"
	"database/sql"
	"errors"
	"hanoman-id/xendit-payment/internal/apis/operations/provider"
	"hanoman-id/xendit-payment/pkg/utils"
)

func (uc *useCase) UpdateProvider(ctx context.Context, params provider.PutProviderParams) (*string, error) {
	data, err := uc.repo.GetUSer(ctx, params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New(utils.ErrUserNotFound)
		}
		return nil, err
	}

	err = utils.ComparePassword(data.Password, params.Body.Password)
	if err != nil {
		return nil, errors.New(utils.ErrInvalidPassword)
	}

	err = uc.repo.UpdateAllFalse(ctx)
	if err != nil {
		return nil, err
	}

	err = uc.repo.UpdateProvider(ctx, params)
	if err != nil {
		return nil, err
	}

	resp := utils.SuccessUpdateProvider

	return &resp, nil

}

func (uc *useCase) GetProvider(ctx context.Context) ([]*provider.GetProviderOKBodyResponseDataItems0, error) {
	var resp []*provider.GetProviderOKBodyResponseDataItems0
	data, err := uc.repo.GetProvider(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		resp = append(resp, &provider.GetProviderOKBodyResponseDataItems0{
			ID:        v.ID,
			IsActived: v.IsSelected,
			Name:      v.Name,
		})
	}

	return resp, nil

}
