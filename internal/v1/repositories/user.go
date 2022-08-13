package repositories

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/provider"
	"hanoman-id/xendit-payment/internal/repository/query"
)

func (r *repositories) GetUSer(ctx context.Context, params provider.PutProviderParams) (*query.GetUserRow, error) {
	data, err := r.qry.GetUser(ctx, params.Body.Username)
	if err != nil {
		return nil, err
	}

	return data, nil
}
