package handler

import (
	"context"
	"hanoman-id/xendit-payment/internal/apis/operations/template"
	"hanoman-id/xendit-payment/pkg/utils"
)

func (h *handler) CreateTemplate(ctx context.Context, params template.PostTemplateParams) (*template.PostTemplateOKBody, error) {
	_, err := h.useCase.CreateTemplate(ctx, params)
	if err != nil {
		return nil, err
	}
	res := &template.PostTemplateOKBody{
		ResponseCode:    utils.StatusCodeSuccess,
		ResponseData:    "",
		ResponseMessege: utils.SuccessCreateTemplate,
	}
	return res, nil
}

func (h *handler) GetListTemplate(ctx context.Context) (*template.GetTemplateOKBody, error) {
	data, err := h.useCase.GetListTemplate(ctx)
	if err != nil {
		return nil, err
	}
	res := &template.GetTemplateOKBody{
		ResponseCode:    utils.StatusCodeSuccess,
		ResponseData:    *data,
		ResponseMessege: utils.SuccessResponseMessege,
	}
	return res, nil
}

func (h *handler) EditTemplate(ctx context.Context, params template.PutTemplateTemplateIDParams) (*template.PutTemplateTemplateIDOKBody, error) {
	_, err := h.useCase.EditTemplate(ctx, params)
	if err != nil {
		return nil, err
	}
	res := &template.PutTemplateTemplateIDOKBody{
		ResponseCode:    utils.StatusCodeSuccess,
		ResponseData:    "",
		ResponseMessege: utils.SuccessEditTemplate,
	}
	return res, nil
}

func (h *handler) DeleteTemplate(ctx context.Context, params template.DeleteTemplateTemplateIDParams) (*template.DeleteTemplateTemplateIDOKBody, error) {
	_, err := h.useCase.DeleteTemplate(ctx, params)
	if err != nil {
		return nil, err
	}
	res := &template.DeleteTemplateTemplateIDOKBody{
		ResponseCode:    utils.StatusCodeSuccess,
		ResponseData:    "",
		ResponseMessege: utils.SuccessDeleteTemplate,
	}
	return res, nil
}
