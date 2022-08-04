package handler

import (
	"context"
	"fmt"
	"hanoman-id/xendit-payment/internal/apis/operations/health"
)

func (h *handler) GetWelcome(ctx context.Context, params health.GetWelcomeParams, principal interface{}) (health.GetWelcomeOKBody, error) {
	id := principal.(map[string]interface{})["id"].(string)
	name := principal.(map[string]interface{})["username"].(string)

	res := fmt.Sprintf("Hai %s, ID anda adalah %s", name, id)
	return health.GetWelcomeOKBody{
		Message: res,
	}, nil
}

func (h *handler) ReadirectToken(ctx context.Context, params health.GetRedirectParams) (health.GetRedirectOKBody, error) {
	return health.GetRedirectOKBody{
		AccessToken: params.AccessToken,
	}, nil
}
