package usecase

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/template"
	"hanoman-id/xendit-payment/internal/models"
	"hanoman-id/xendit-payment/pkg/utils"

	"github.com/google/uuid"
)

func (uc *useCase) CreateTemplate(ctx context.Context, params template.PostTemplateParams) (*string, error) {
	id := uuid.New().String()
	err := uc.repo.CreateTemplate(ctx, params, id)
	if err != nil {
		return nil, err
	}

	resp := utils.SuccessCreateTemplate

	return &resp, nil
}

func (uc *useCase) GetListTemplate(ctx context.Context) (*models.GetTemplate, error) {
	var resp models.GetTemplate

	data, err := uc.repo.GetListTemplate(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		resp = append(resp, &models.GetTemplateItems0{
			AppsName: v.AppsName,
			ID:       v.ID,
			Text:     v.Text,
			Type:     v.Type,
		})
	}

	return &resp, nil
}

func (uc *useCase) EditTemplate(ctx context.Context, params template.PutTemplateTemplateIDParams) (*string, error) {
	err := uc.repo.EditTemplate(ctx, params)
	if err != nil {
		return nil, err
	}

	resp := utils.SuccessEditTemplate

	return &resp, nil
}

func (uc *useCase) DeleteTemplate(ctx context.Context, params template.DeleteTemplateTemplateIDParams) (*string, error) {
	err := uc.repo.DeleteTemplate(ctx, params)
	if err != nil {
		return nil, err
	}

	resp := utils.SuccessDeleteTemplate

	return &resp, nil
}
