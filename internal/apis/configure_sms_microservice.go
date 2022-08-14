// This file is safe to edit. Once it exists it will not be overwritten

package apis

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/net/context"

	"hanoman-id/xendit-payment/internal/apis/operations"
	"hanoman-id/xendit-payment/internal/apis/operations/provider"
	"hanoman-id/xendit-payment/internal/apis/operations/sms"
	"hanoman-id/xendit-payment/internal/apis/operations/template"
	"hanoman-id/xendit-payment/internal/models"
	"hanoman-id/xendit-payment/internal/v1/handler"
	"hanoman-id/xendit-payment/pkg/utils"
)

//go:generate swagger generate server --target ../../../sms-microservice --name SmsMicroservice --spec ../../api/v1/swagger.yml --model-package internal/models --server-package internal/apis --principal interface{}

func configureFlags(api *operations.SmsMicroserviceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SmsMicroserviceAPI) http.Handler {
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

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.TemplateDeleteTemplateTemplateIDHandler != nil {
		api.TemplateDeleteTemplateTemplateIDHandler = template.DeleteTemplateTemplateIDHandlerFunc(func(params template.DeleteTemplateTemplateIDParams) middleware.Responder {
			// return middleware.NotImplemented("operation template.DeleteTemplateTemplateID has not yet been implemented")
			result, err := handler.NewHandler().DeleteTemplate(context.Background(), params)
			if err != nil {
				template.NewDeleteTemplateTemplateIDBadRequest().WithPayload(&models.Error{
					Code:    utils.SuccessDeleteTemplate,
					Message: err.Error(),
				})
			}
			return template.NewDeleteTemplateTemplateIDOK().WithPayload(result)
		})
	}
	if api.ProviderGetProviderHandler != nil {
		api.ProviderGetProviderHandler = provider.GetProviderHandlerFunc(func(params provider.GetProviderParams) middleware.Responder {
			// return middleware.NotImplemented("operation provider.GetProvider has not yet been implemented")
			result, err := handler.NewHandler().GetProvider(context.Background())
			if err != nil {
				return provider.NewGetProviderBadRequest().WithPayload(&models.Error{
					Code:    utils.StatusCodeBadRequest,
					Message: err.Error(),
				})
			}
			return provider.NewGetProviderOK().WithPayload(result)
		})
	}
	if api.SmsGetSmsHistoryHandler != nil {
		api.SmsGetSmsHistoryHandler = sms.GetSmsHistoryHandlerFunc(func(params sms.GetSmsHistoryParams) middleware.Responder {
			// return middleware.NotImplemented("operation sms.GetSmsHistory has not yet been implemented")
			result, err := handler.NewHandler().GetSmsHistory(context.Background())
			if err != nil {
				return sms.NewGetSmsHistoryBadRequest().WithPayload(&models.Error{
					Code:    utils.StatusCodeBadRequest,
					Message: err.Error(),
				})
			}
			return sms.NewGetSmsHistoryOK().WithPayload(result)
		})
	}
	if api.TemplateGetTemplateHandler != nil {
		api.TemplateGetTemplateHandler = template.GetTemplateHandlerFunc(func(params template.GetTemplateParams) middleware.Responder {
			// return middleware.NotImplemented("operation template.GetTemplate has not yet been implemented")
			result, err := handler.NewHandler().GetListTemplate(context.Background())
			if err != nil {
				return template.NewGetTemplateBadRequest().WithPayload(&models.Error{
					Code:    utils.StatusCodeBadRequest,
					Message: err.Error(),
				})
			}
			return template.NewGetTemplateOK().WithPayload(result)
		})
	}
	if api.SmsPostSmsHandler != nil {
		api.SmsPostSmsHandler = sms.PostSmsHandlerFunc(func(params sms.PostSmsParams) middleware.Responder {
			// return middleware.NotImplemented("operation sms.PostSms has not yet been implemented")
			result, err := handler.NewHandler().CreateSms(context.Background(), params)
			if err != nil {
				return sms.NewGetSmsHistoryBadRequest().WithPayload(&models.Error{
					Code:    utils.StatusCodeBadRequest,
					Message: err.Error(),
				})
			}
			return sms.NewPostSmsOK().WithPayload(result)
		})
	}
	if api.TemplatePostTemplateHandler != nil {
		api.TemplatePostTemplateHandler = template.PostTemplateHandlerFunc(func(params template.PostTemplateParams) middleware.Responder {
			// return middleware.NotImplemented("operation template.PostTemplate has not yet been implemented")
			result, err := handler.NewHandler().CreateTemplate(context.Background(), params)
			if err != nil {
				return template.NewPostTemplateBadRequest().WithPayload(&models.Error{
					Code:    utils.StatusCodeBadRequest,
					Message: err.Error(),
				})
			}
			return template.NewPostTemplateOK().WithPayload(result)
		})
	}
	if api.ProviderPutProviderHandler != nil {
		api.ProviderPutProviderHandler = provider.PutProviderHandlerFunc(func(params provider.PutProviderParams) middleware.Responder {
			// return middleware.NotImplemented("operation provider.PutProviderProviderID has not yet been implemented")
			result, err := handler.NewHandler().UpdateProvider(context.Background(), params)
			if err != nil {
				return provider.NewPutProviderProviderIDBadRequest().WithPayload(&models.Error{
					Code:    utils.StatusCodeBadRequest,
					Message: err.Error(),
				})
			}
			return provider.NewPutProviderProviderIDOK().WithPayload(result)
		})
	}
	if api.TemplatePutTemplateTemplateIDHandler != nil {
		api.TemplatePutTemplateTemplateIDHandler = template.PutTemplateTemplateIDHandlerFunc(func(params template.PutTemplateTemplateIDParams) middleware.Responder {
			// return middleware.NotImplemented("operation template.PutTemplateTemplateID has not yet been implemented")
			result, err := handler.NewHandler().EditTemplate(context.Background(), params)
			if err != nil {
				return template.NewPutTemplateTemplateIDBadRequest().WithPayload(&models.Error{
					Code:    utils.StatusCodeBadRequest,
					Message: err.Error(),
				})
			}
			return template.NewPutTemplateTemplateIDOK().WithPayload(result)
		})
	}

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
