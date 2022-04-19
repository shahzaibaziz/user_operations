// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/shahzaibaziz/user_operations/gen/restapi/operations"
	"github.com/shahzaibaziz/user_operations/gen/restapi/operations/service"
)

//go:generate swagger generate server --target ../../gen --name UserOperations --spec ../../swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.UserOperationsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UserOperationsAPI) http.Handler {
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

	if api.ServiceDeleteUserHandler == nil {
		api.ServiceDeleteUserHandler = service.DeleteUserHandlerFunc(func(params service.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation service.DeleteUser has not yet been implemented")
		})
	}
	if api.ServiceGetAllUsersHandler == nil {
		api.ServiceGetAllUsersHandler = service.GetAllUsersHandlerFunc(func(params service.GetAllUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation service.GetAllUsers has not yet been implemented")
		})
	}
	if api.ServiceGetUserHandler == nil {
		api.ServiceGetUserHandler = service.GetUserHandlerFunc(func(params service.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation service.GetUser has not yet been implemented")
		})
	}
	if api.ServiceRegisterUserHandler == nil {
		api.ServiceRegisterUserHandler = service.RegisterUserHandlerFunc(func(params service.RegisterUserParams) middleware.Responder {
			return middleware.NotImplemented("operation service.RegisterUser has not yet been implemented")
		})
	}
	if api.ServiceUpdateUserHandler == nil {
		api.ServiceUpdateUserHandler = service.UpdateUserHandlerFunc(func(params service.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation service.UpdateUser has not yet been implemented")
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
// Scheme value will be set accordingly: "http", "https" or "unix".
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
