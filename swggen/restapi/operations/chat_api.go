// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/kameike/chat_api/swggen/restapi/operations/account"
	"github.com/kameike/chat_api/swggen/restapi/operations/chat_rooms"
	"github.com/kameike/chat_api/swggen/restapi/operations/deploy"
)

// NewChatAPI creates a new Chat instance
func NewChatAPI(spec *loads.Document) *ChatAPI {
	return &ChatAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		ChatRoomsGetChatroomsIDHandler: chat_rooms.GetChatroomsIDHandlerFunc(func(params chat_rooms.GetChatroomsIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ChatRoomsGetChatroomsID has not yet been implemented")
		}),
		ChatRoomsGetChatroomsIDMessagesHandler: chat_rooms.GetChatroomsIDMessagesHandlerFunc(func(params chat_rooms.GetChatroomsIDMessagesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ChatRoomsGetChatroomsIDMessages has not yet been implemented")
		}),
		DeployGetHealthHandler: deploy.GetHealthHandlerFunc(func(params deploy.GetHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation DeployGetHealth has not yet been implemented")
		}),
		AccountPostAuthHandler: account.PostAuthHandlerFunc(func(params account.PostAuthParams) middleware.Responder {
			return middleware.NotImplemented("operation AccountPostAuth has not yet been implemented")
		}),
		ChatRoomsPostChatroomsHandler: chat_rooms.PostChatroomsHandlerFunc(func(params chat_rooms.PostChatroomsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ChatRoomsPostChatrooms has not yet been implemented")
		}),
		ChatRoomsPostChatroomsIDMessagesHandler: chat_rooms.PostChatroomsIDMessagesHandlerFunc(func(params chat_rooms.PostChatroomsIDMessagesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ChatRoomsPostChatroomsIDMessages has not yet been implemented")
		}),
		ChatRoomsPostChatroomsIDReadHandler: chat_rooms.PostChatroomsIDReadHandlerFunc(func(params chat_rooms.PostChatroomsIDReadParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ChatRoomsPostChatroomsIDRead has not yet been implemented")
		}),
		AccountPostProfileHandler: account.PostProfileHandlerFunc(func(params account.PostProfileParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation AccountPostProfile has not yet been implemented")
		}),

		// Applies when the "x_chat_access_token" header is set
		APIKeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) x_chat_access_token from header param [x_chat_access_token] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*ChatAPI From the todo list tutorial on goswagger.io */
type ChatAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key x_chat_access_token provided in the header
	APIKeyAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ChatRoomsGetChatroomsIDHandler sets the operation handler for the get chatrooms ID operation
	ChatRoomsGetChatroomsIDHandler chat_rooms.GetChatroomsIDHandler
	// ChatRoomsGetChatroomsIDMessagesHandler sets the operation handler for the get chatrooms ID messages operation
	ChatRoomsGetChatroomsIDMessagesHandler chat_rooms.GetChatroomsIDMessagesHandler
	// DeployGetHealthHandler sets the operation handler for the get health operation
	DeployGetHealthHandler deploy.GetHealthHandler
	// AccountPostAuthHandler sets the operation handler for the post auth operation
	AccountPostAuthHandler account.PostAuthHandler
	// ChatRoomsPostChatroomsHandler sets the operation handler for the post chatrooms operation
	ChatRoomsPostChatroomsHandler chat_rooms.PostChatroomsHandler
	// ChatRoomsPostChatroomsIDMessagesHandler sets the operation handler for the post chatrooms ID messages operation
	ChatRoomsPostChatroomsIDMessagesHandler chat_rooms.PostChatroomsIDMessagesHandler
	// ChatRoomsPostChatroomsIDReadHandler sets the operation handler for the post chatrooms ID read operation
	ChatRoomsPostChatroomsIDReadHandler chat_rooms.PostChatroomsIDReadHandler
	// AccountPostProfileHandler sets the operation handler for the post profile operation
	AccountPostProfileHandler account.PostProfileHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ChatAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ChatAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ChatAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ChatAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ChatAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ChatAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ChatAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ChatAPI
func (o *ChatAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "XChatAccessTokenAuth")
	}

	if o.ChatRoomsGetChatroomsIDHandler == nil {
		unregistered = append(unregistered, "chat_rooms.GetChatroomsIDHandler")
	}

	if o.ChatRoomsGetChatroomsIDMessagesHandler == nil {
		unregistered = append(unregistered, "chat_rooms.GetChatroomsIDMessagesHandler")
	}

	if o.DeployGetHealthHandler == nil {
		unregistered = append(unregistered, "deploy.GetHealthHandler")
	}

	if o.AccountPostAuthHandler == nil {
		unregistered = append(unregistered, "account.PostAuthHandler")
	}

	if o.ChatRoomsPostChatroomsHandler == nil {
		unregistered = append(unregistered, "chat_rooms.PostChatroomsHandler")
	}

	if o.ChatRoomsPostChatroomsIDMessagesHandler == nil {
		unregistered = append(unregistered, "chat_rooms.PostChatroomsIDMessagesHandler")
	}

	if o.ChatRoomsPostChatroomsIDReadHandler == nil {
		unregistered = append(unregistered, "chat_rooms.PostChatroomsIDReadHandler")
	}

	if o.AccountPostProfileHandler == nil {
		unregistered = append(unregistered, "account.PostProfileHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ChatAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ChatAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "api_key":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.APIKeyAuth)

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *ChatAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *ChatAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ChatAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ChatAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the chat API
func (o *ChatAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ChatAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/chatrooms/{id}"] = chat_rooms.NewGetChatroomsID(o.context, o.ChatRoomsGetChatroomsIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/chatrooms/{id}/messages"] = chat_rooms.NewGetChatroomsIDMessages(o.context, o.ChatRoomsGetChatroomsIDMessagesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health"] = deploy.NewGetHealth(o.context, o.DeployGetHealthHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth"] = account.NewPostAuth(o.context, o.AccountPostAuthHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/chatrooms"] = chat_rooms.NewPostChatrooms(o.context, o.ChatRoomsPostChatroomsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/chatrooms/{id}/messages"] = chat_rooms.NewPostChatroomsIDMessages(o.context, o.ChatRoomsPostChatroomsIDMessagesHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/chatrooms/{id}/read"] = chat_rooms.NewPostChatroomsIDRead(o.context, o.ChatRoomsPostChatroomsIDReadHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/profile"] = account.NewPostProfile(o.context, o.AccountPostProfileHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ChatAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ChatAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ChatAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ChatAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
