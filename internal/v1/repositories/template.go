package repositories

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/template"
	"hanoman-id/xendit-payment/internal/repository/query"
)

func (r *repositories) CreateTemplate(ctx context.Context, params template.PostTemplateParams, id string) error {
	args := &query.CreateTemplateParams{
		ID:       id,
		Type:     params.Body.Type,
		AppsName: params.Body.AppsName,
		Text:     params.Body.Text,
	}
	_, err := r.qry.CreateTemplate(ctx, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositories) GetListTemplate(ctx context.Context) ([]*query.GetListTemplateRow, error) {
	data, err := r.qry.GetListTemplate(ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *repositories) EditTemplate(ctx context.Context, params template.PutTemplateTemplateIDParams) error {
	args := &query.UpdateTemplateParams{
		Type:     params.Body.Type,
		AppsName: params.Body.AppsName,
		Text:     params.Body.Text,
		ID:       params.TemplateID,
	}
	err := r.qry.UpdateTemplate(ctx, args)
	if err != nil {
		return err
	}

	return nil
}

func (r *repositories) DeleteTemplate(ctx context.Context, params template.DeleteTemplateTemplateIDParams) error {
	err := r.qry.DeleteTemplate(ctx, params.TemplateID)
	if err != nil {
		return err
	}

	return nil
}
