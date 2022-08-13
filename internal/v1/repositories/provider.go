package repositories

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/provider"
	"hanoman-id/xendit-payment/internal/repository/query"
)

func (r *repositories) GetProvider(ctx context.Context) ([]*query.GetProviderRow, error) {
	data, err := r.qry.GetProvider(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repositories) UpdateProvider(ctx context.Context, params provider.PutProviderParams) error {
	err := r.qry.UpdateProvider(ctx, params.Body.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositories) UpdateAllFalse(ctx context.Context) error {
	err := r.qry.UpdateFalseProvider(ctx)
	if err != nil {
		return err
	}

	return nil
}
