// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/kameike/chat_api/swggen/restapi/operations"
	"github.com/kameike/chat_api/swggen/restapi/operations/auth"
	"github.com/kameike/chat_api/swggen/restapi/operations/chat_rooms"
	"github.com/kameike/chat_api/swggen/restapi/operations/deploy"
)

//go:generate swagger generate server --target ../../swggen --name Chat --spec ../../swagger.yml --model-package apimodel --exclude-main

func configureFlags(api *operations.ChatAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ChatAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x_chat_access_token" header is set
	api.APIKeyAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (api_key) x_chat_access_token from header param [x_chat_access_token] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.ChatRoomsGetChatroomsIDHandler = chat_rooms.GetChatroomsIDHandlerFunc(func(params chat_rooms.GetChatroomsIDParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation chat_rooms.GetChatroomsID has not yet been implemented")
	})
	api.ChatRoomsGetChatroomsIDMessagesHandler = chat_rooms.GetChatroomsIDMessagesHandlerFunc(func(params chat_rooms.GetChatroomsIDMessagesParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation chat_rooms.GetChatroomsIDMessages has not yet been implemented")
	})
	api.DeployGetHealthHandler = deploy.GetHealthHandlerFunc(func(params deploy.GetHealthParams) middleware.Responder {
		return middleware.NotImplemented("operation deploy.GetHealth has not yet been implemented")
	})
	api.AuthPostAuthHandler = auth.PostAuthHandlerFunc(func(params auth.PostAuthParams) middleware.Responder {
		return middleware.NotImplemented("operation auth.PostAuth has not yet been implemented")
	})
	api.ChatRoomsPostChatroomsHandler = chat_rooms.PostChatroomsHandlerFunc(func(params chat_rooms.PostChatroomsParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation chat_rooms.PostChatrooms has not yet been implemented")
	})
	api.ChatRoomsPostChatroomsIDMessagesHandler = chat_rooms.PostChatroomsIDMessagesHandlerFunc(func(params chat_rooms.PostChatroomsIDMessagesParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation chat_rooms.PostChatroomsIDMessages has not yet been implemented")
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
