// This file is safe to edit. Once it exists it will not be overwritten

package apis

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"hanoman-id/xendit-payment/internal/apis/operations"
	"hanoman-id/xendit-payment/internal/apis/operations/auth"
	"hanoman-id/xendit-payment/internal/apis/operations/health"
	"hanoman-id/xendit-payment/internal/handler"
	"hanoman-id/xendit-payment/internal/models"
	"hanoman-id/xendit-payment/internal/utils"
)

//go:generate swagger generate server --target ../../../xendit-payment --name HanomanPayment --spec ../../api/swagger.yml --model-package internal/models --server-package internal/apis --principal interface{}

func configureFlags(api *operations.HanomanPaymentAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HanomanPaymentAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	if api.BearerAuth != nil {
		api.BearerAuth = func(token string) (interface{}, error) {
			// return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
			ctx, err := utils.ValidateHeader(token)
			if err != nil {
				return nil, err
			}
			return ctx, nil
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.HealthGetWelcomeHandler != nil {
		api.HealthGetWelcomeHandler = health.GetWelcomeHandlerFunc(func(params health.GetWelcomeParams, principal interface{}) middleware.Responder {
			// return middleware.NotImplemented("operation health.GetWelcome has not yet been implemented")

			result, err := handler.NewHandler().GetWelcome(context.Background(), params, principal)
			if err != nil {
				return nil
			}

			return health.NewGetWelcomeOK().WithPayload(&health.GetWelcomeOKBody{
				Message: result.Message,
			})
		})
	}
	if api.AuthPostCreateOtpHandler != nil {
		api.AuthPostCreateOtpHandler = auth.PostCreateOtpHandlerFunc(func(params auth.PostCreateOtpParams) middleware.Responder {
			// return middleware.NotImplemented("operation auth.PostCreateOtp has not yet been implemented")
			result, err := handler.NewHandler().CreateOtp(context.Background(), params)
			if err != nil {
				return auth.NewPostCreateOtpBadRequest().WithPayload(&models.Error{
					Code:    "400",
					Message: err.Error(),
				})
			}
			return auth.NewPostCreateOtpOK().WithPayload(&auth.PostCreateOtpOKBody{
				Data:    result,
				Message: "Success",
			})
		})
	}
	if api.AuthPostValidateOtpHandler != nil {
		api.AuthPostValidateOtpHandler = auth.PostValidateOtpHandlerFunc(func(params auth.PostValidateOtpParams) middleware.Responder {
			// return middleware.NotImplemented("operation auth.PostValidateOtp has not yet been implemented")
			result, err := handler.NewHandler().ValidateOtp(context.Background(), params)
			if err != nil {
				return auth.NewPostValidateOtpBadRequest().WithPayload(&models.Error{})
			}

			return auth.NewPostValidateOtpOK().WithPayload(&auth.PostValidateOtpOKBody{
				Data:    result,
				Message: "SUCCESS",
			})
		})
	}

	api.HealthGetRedirectHandler = health.GetRedirectHandlerFunc(func(params health.GetRedirectParams) middleware.Responder {
		resp, err := handler.NewHandler().ReadirectToken(context.Background(), params)
		if err != nil {
			return health.NewGetWelcomeBadRequest().WithPayload(&models.Error{
				Code:    "400",
				Message: err.Error(),
			})
		}
		return health.NewGetRedirectOK().WithPayload(&resp)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
