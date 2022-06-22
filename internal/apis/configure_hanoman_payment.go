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
	"hanoman-id/xendit-payment/internal/apis/operations/health"
	"hanoman-id/xendit-payment/internal/apis/operations/payment"
	"hanoman-id/xendit-payment/internal/handler"
	"hanoman-id/xendit-payment/internal/models"
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

	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(func(ghp health.GetHealthParams) middleware.Responder {
		return health.NewGetHealthOK().WithPayload(&health.GetHealthOKBody{
			Data: &models.Health{
				Status: "200",
			},
			Message: "Good",
		})
	})

	api.PaymentGetMakingPaymentsHandler = payment.GetMakingPaymentsHandlerFunc(func(params payment.GetMakingPaymentsParams) middleware.Responder {
		result, err := handler.NewHandler().GetMakingPayment(context.Background(), params)
		if err != nil {
			return payment.NewGetMakingPaymentsBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}
		return payment.NewGetMakingPaymentsOK().WithPayload(&payment.GetMakingPaymentsOKBody{
			Data:    result,
			Message: "success",
		})
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
