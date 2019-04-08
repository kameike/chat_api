package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/handler"
	"github.com/kameike/chat_api/repository"
	"github.com/kameike/chat_api/swggen/restapi"
	"github.com/kameike/chat_api/swggen/restapi/operations"
	"github.com/kameike/chat_api/swggen/restapi/operations/account"
	"github.com/kameike/chat_api/swggen/restapi/operations/chat_rooms"
	"github.com/kameike/chat_api/swggen/restapi/operations/deploy"
)

var repo repository.ReposotryProvidable

func main() {
	handlers := handler.SetUpHandler()
	repo = repository.CreateAppRepositoryProvider()
	defer repo.Close()

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewChatAPI(swaggerSpec)

	api.APIKeyAuth = func(token string) (interface{}, error) {
		result, err := handlers.AuthUser(token)
		return result, err
	}

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.RegisterConsumer("application/json", runtime.JSONConsumer())
	api.RegisterProducer("application/json", runtime.JSONProducer())

	api.ChatRoomsGetChatroomsIDHandler = chat_rooms.GetChatroomsIDHandlerFunc(func(params chat_rooms.GetChatroomsIDParams, principal interface{}) middleware.Responder {

		return middleware.NotImplemented("operation chat_rooms.GetChatroomsID has not yet been implemented")
	})
	api.ChatRoomsGetChatroomsIDMessagesHandler = chat_rooms.GetChatroomsIDMessagesHandlerFunc(func(params chat_rooms.GetChatroomsIDMessagesParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation chat_rooms.GetChatroomsIDMessages has not yet been implemented")
	})
	api.DeployGetHealthHandler = deploy.GetHealthHandlerFunc(func(params deploy.GetHealthParams) middleware.Responder {
		return middleware.NotImplemented("operation deploy.GetHealth has not yet been implemented")
	})
	api.AccountPostAuthHandler = account.PostAuthHandlerFunc(func(params account.PostAuthParams) middleware.Responder {
		return middleware.NotImplemented("operation account.PostAuth has not yet been implemented")
	})
	api.ChatRoomsPostChatroomsIDMessagesHandler = chat_rooms.PostChatroomsIDMessagesHandlerFunc(func(params chat_rooms.PostChatroomsIDMessagesParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation chat_rooms.PostChatroomsIDMessages has not yet been implemented")
	})
	api.ChatRoomsPostChatroomsIDReadHandler = chat_rooms.PostChatroomsIDReadHandlerFunc(func(params chat_rooms.PostChatroomsIDReadParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation chat_rooms.PostChatroomsIDRead has not yet been implemented")
	})

	api.ChatRoomsPostChatroomsHandler = handlers.ChatRoomsPostChatroomsHandler()
	api.AccountPostProfileHandler = handlers.AccountPostProfileHandler()
	api.DeployGetHealthHandler = handlers.DeployGetHealthHandler()
	api.AccountPostAuthHandler = handlers.AccountPostAuthHandler()

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 1323

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
