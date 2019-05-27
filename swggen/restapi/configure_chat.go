// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"net/http"

	"github.com/kameike/chat_api/swggen/restapi/operations"
	"github.com/rs/cors"
)

//go:generate swagger generate server --target ../../swggen --name Chat --spec ../../swagger.yaml --model-package apimodel --exclude-main

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

	// Applies when the "X-CHAT-ACCESS-TOKEN" header is set
	//	api.APIKeyAuth = func(token string) (interface{}, error) {
	//		return nil, errors.NotImplemented("api key auth (apiKey) X-CHAT-ACCESS-TOKEN from header param [X-CHAT-ACCESS-TOKEN] has not yet been implemented")
	//	}
	//
	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	// if api.ChatroomsGetAdminSearchChatroomsHandler == nil {
	// 	api.ChatroomsGetAdminSearchChatroomsHandler = chatrooms.GetAdminSearchChatroomsHandlerFunc(func(params chatrooms.GetAdminSearchChatroomsParams, principal interface{}) middleware.Responder {
	// 		return middleware.NotImplemented("operation chatrooms.GetAdminSearchChatrooms has not yet been implemented")
	// 	})
	// }
	// if api.MessagesGetChatroomsChatroomHashMessagesHandler == nil {
	// 	api.MessagesGetChatroomsChatroomHashMessagesHandler = messages.GetChatroomsChatroomHashMessagesHandlerFunc(func(params messages.GetChatroomsChatroomHashMessagesParams, principal interface{}) middleware.Responder {
	// 		return middleware.NotImplemented("operation messages.GetChatroomsChatroomHashMessages has not yet been implemented")
	// 	})
	// }
	// if api.ChatroomsGetChatroomsIDHandler == nil {
	// 	api.ChatroomsGetChatroomsIDHandler = chatrooms.GetChatroomsIDHandlerFunc(func(params chatrooms.GetChatroomsIDParams, principal interface{}) middleware.Responder {
	// 		return middleware.NotImplemented("operation chatrooms.GetChatroomsID has not yet been implemented")
	// 	})
	// }
	// if api.DeployGetHealthHandler == nil {
	// 	api.DeployGetHealthHandler = deploy.GetHealthHandlerFunc(func(params deploy.GetHealthParams) middleware.Responder {
	// 		return middleware.NotImplemented("operation deploy.GetHealth has not yet been implemented")
	// 	})
	// }
	// if api.AccountPostAuthHandler == nil {
	// 	api.AccountPostAuthHandler = account.PostAuthHandlerFunc(func(params account.PostAuthParams) middleware.Responder {
	// 		return middleware.NotImplemented("operation account.PostAuth has not yet been implemented")
	// 	})
	// }
	// if api.ChatroomsPostChatroomsHandler == nil {
	// 	api.ChatroomsPostChatroomsHandler = chatrooms.PostChatroomsHandlerFunc(func(params chatrooms.PostChatroomsParams, principal interface{}) middleware.Responder {
	// 		return middleware.NotImplemented("operation chatrooms.PostChatrooms has not yet been implemented")
	// 	})
	// }
	// if api.ChatroomsPostChatroomsChannelHashReadHandler == nil {
	// 	api.ChatroomsPostChatroomsChannelHashReadHandler = chatrooms.PostChatroomsChannelHashReadHandlerFunc(func(params chatrooms.PostChatroomsChannelHashReadParams, principal interface{}) middleware.Responder {
	// 		return middleware.NotImplemented("operation chatrooms.PostChatroomsChannelHashRead has not yet been implemented")
	// 	})
	// }
	// if api.MessagesPostChatroomsChatroomHashMessagesHandler == nil {
	// 	api.MessagesPostChatroomsChatroomHashMessagesHandler = messages.PostChatroomsChatroomHashMessagesHandlerFunc(func(params messages.PostChatroomsChatroomHashMessagesParams, principal interface{}) middleware.Responder {
	// 		return middleware.NotImplemented("operation messages.PostChatroomsChatroomHashMessages has not yet been implemented")
	// 	})
	// }
	// if api.AccountPostProfileHandler == nil {
	// 	api.AccountPostProfileHandler = account.PostProfileHandlerFunc(func(params account.PostProfileParams, principal interface{}) middleware.Responder {
	// 		return middleware.NotImplemented("operation account.PostProfile has not yet been implemented")
	// 	})
	// }

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
	corsHandler := cors.New(cors.Options{
		Debug:          false,
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{},
		MaxAge:         1000,
	})
	return corsHandler.Handler(handler)
}
