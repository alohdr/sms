package repositories

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/sms"
	"hanoman-id/xendit-payment/internal/repository/query"
)

func (r *repositories) CreateSms(ctx context.Context, params sms.PostSmsParams, id string) error {
	arg := &query.CreatSmsParams{
		ID:          id,
		Sender:      params.Body.AppsName,
		PhoneNumber: params.Body.PhoneNumber,
		Type:        params.Body.Type,
		Status:      "SUCCESS",
	}
	_, err := r.qry.CreatSms(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositories) GetSmsHistory(ctx context.Context) ([]*query.GetSmsHistoryRow, error) {
	data, err := r.qry.GetSmsHistory(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}
