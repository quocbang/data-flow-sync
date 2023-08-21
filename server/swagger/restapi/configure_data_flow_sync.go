// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/limitary_hour"
)

//go:generate swagger generate server --target ..\..\swagger --name DataFlowSync --spec ..\..\..\swagger.yml --principal models.Principal

func configureFlags(api *operations.DataFlowSyncAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DataFlowSyncAPI) http.Handler {
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

	// Applies when the "x-data-flow-sync-auth-key" header is set
	if api.APIKeyAuth == nil {
		api.APIKeyAuth = func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (api_key) x-data-flow-sync-auth-key from header param [x-data-flow-sync-auth-key] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.LimitaryHourCreateMergeRequestHandler == nil {
		api.LimitaryHourCreateMergeRequestHandler = limitary_hour.CreateMergeRequestHandlerFunc(func(params limitary_hour.CreateMergeRequestParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation limitary_hour.CreateMergeRequest has not yet been implemented")
		})
	}
	if api.LimitaryHourGetLimitaryHourDiffHandler == nil {
		api.LimitaryHourGetLimitaryHourDiffHandler = limitary_hour.GetLimitaryHourDiffHandlerFunc(func(params limitary_hour.GetLimitaryHourDiffParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation limitary_hour.GetLimitaryHourDiff has not yet been implemented")
		})
	}
	if api.LimitaryHourListMergeRequestsHandler == nil {
		api.LimitaryHourListMergeRequestsHandler = limitary_hour.ListMergeRequestsHandlerFunc(func(params limitary_hour.ListMergeRequestsParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation limitary_hour.ListMergeRequests has not yet been implemented")
		})
	}
	if api.AccountLoginHandler == nil {
		api.AccountLoginHandler = account.LoginHandlerFunc(func(params account.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation account.Login has not yet been implemented")
		})
	}
	if api.AccountLogoutHandler == nil {
		api.AccountLogoutHandler = account.LogoutHandlerFunc(func(params account.LogoutParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.Logout has not yet been implemented")
		})
	}
	if api.AccountSendMailHandler == nil {
		api.AccountSendMailHandler = account.SendMailHandlerFunc(func(params account.SendMailParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.SendMail has not yet been implemented")
		})
	}
	if api.AccountSignupHandler == nil {
		api.AccountSignupHandler = account.SignupHandlerFunc(func(params account.SignupParams) middleware.Responder {
			return middleware.NotImplemented("operation account.Signup has not yet been implemented")
		})
	}
	if api.LimitaryHourUploadLimitaryHourHandler == nil {
		api.LimitaryHourUploadLimitaryHourHandler = limitary_hour.UploadLimitaryHourHandlerFunc(func(params limitary_hour.UploadLimitaryHourParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation limitary_hour.UploadLimitaryHour has not yet been implemented")
		})
	}
	if api.AccountVerifyAccountHandler == nil {
		api.AccountVerifyAccountHandler = account.VerifyAccountHandlerFunc(func(params account.VerifyAccountParams) middleware.Responder {
			return middleware.NotImplemented("operation account.VerifyAccount has not yet been implemented")
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
