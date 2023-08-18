// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"github.com/quocbang/data-flow-sync/server"
	apiService "github.com/quocbang/data-flow-sync/server/api"
	"github.com/quocbang/data-flow-sync/server/config"
	mw "github.com/quocbang/data-flow-sync/server/middleware"
	"github.com/quocbang/data-flow-sync/server/swagger/models"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/account"
	"github.com/quocbang/data-flow-sync/server/swagger/restapi/operations/limitary_hour"
)

//go:generate swagger generate server --target ..\..\swagger --name DataFlowSync --spec ..\..\..\swagger.yml --principal models.Principal

var (
	options        = new(config.Options)
	configurations = new(config.Configs)
)

func configureFlags(api *operations.DataFlowSyncAPI) {
	api.CommandLineOptionsGroups = append(api.CommandLineOptionsGroups, swag.CommandLineOptionsGroup{
		ShortDescription: "Configuration Options",
		LongDescription:  "Configuration Options",
		Options:          options,
	})
}

func parseConfig(filePath string) (*config.Configs, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfs config.Configs
	if err := yaml.Unmarshal(data, &cfs); err != nil {
		return nil, err
	}

	return &cfs, nil
}

func configureAPI(api *operations.DataFlowSyncAPI) http.Handler {
	configs, err := parseConfig(options.ConfigPath)
	if err != nil {
		log.Fatalf("failed to parse config file, error: %v", err)
	}

	// configure the api here
	api.ServeError = errors.ServeError

	// initialize logger
	setLogger(false)

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

	repo, err := server.RegisterRepository(configs.Database)
	if err != nil {
		log.Fatalf("failed to register repository, error: %v", err)
	}

	serviceConfig := apiService.ServiceConfig{
		Repo:         repo,
		MRExpiryTime: configs.MRExpiryTime,
	}

	apiService.RegisterAPI(api, serviceConfig)

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
		api.AccountLogoutHandler = account.LogoutHandlerFunc(func(params account.LogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation account.Logout has not yet been implemented")
		})
	}
	if api.LimitaryHourUploadLimitaryHourHandler == nil {
		api.LimitaryHourUploadLimitaryHourHandler = limitary_hour.UploadLimitaryHourHandlerFunc(func(params limitary_hour.UploadLimitaryHourParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation limitary_hour.UploadLimitaryHour has not yet been implemented")
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
	handler = mw.LoggingMiddleware(handler)
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)
	handler = allowCORS(handler)
	return handler
}

// setLogger replaces global logger and redirects STD logger.
func setLogger(devMode bool) {
	var config zap.Config

	if devMode {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	logger, err := config.Build()
	if err != nil {
		log.Fatalln("failed to initialize logger", err)
	}
	zap.ReplaceGlobals(logger)
	zap.RedirectStdLog(logger)
	defer logger.Sync()
	logger.Info("logger initialized")
}

func allowCORS(handler http.Handler) http.Handler {
	return cors.New(
		cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodHead,
				http.MethodOptions,
			},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: false,
		},
	).Handler(handler)
}
