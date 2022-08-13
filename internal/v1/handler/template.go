package handler

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/template"
	"hanoman-id/xendit-payment/pkg/utils"
)

func (h *handler) CreateTemplate(ctx context.Context, params template.PostTemplateParams) (*template.PostTemplateOKBody, error) {
	data, err := h.useCase.CreateTemplate(ctx, params)
	if err != nil {
		return nil, err
	}
	res := &template.PostTemplateOKBody{
		ResponseCode:    utils.StatusCodeSuccess,
		ResponseData:    *data,
		ResponseMessege: utils.SuccessResponseMessege,
	}
	return res, nil
}

func (h *handler) GetListTemplate(ctx context.Context) (*template.GetTemplateOKBody, error) {
	data, err := h.useCase.GetListTemplate(ctx)
	if err != nil {
		return nil, err
	}
	res := &template.GetTemplateOKBody{
		ResponseCode:    "",
		ResponseData:    *data,
		ResponseMessege: "",
	}
	return res, nil
}

func (h *handler) EditTemplate(ctx context.Context, params template.PutTemplateTemplateIDParams) (*template.PutTemplateTemplateIDOKBody, error) {
	data, err := h.useCase.EditTemplate(ctx, params)
	if err != nil {
		return nil, err
	}
	res := &template.PutTemplateTemplateIDOKBody{
		ResponseCode:    "",
		ResponseData:    *data,
		ResponseMessege: "",
	}
	return res, nil
}
