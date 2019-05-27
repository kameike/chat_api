package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/handler"
	"github.com/kameike/chat_api/repository"
	"github.com/kameike/chat_api/swggen/restapi"
	"github.com/kameike/chat_api/swggen/restapi/operations"
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

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	api.RegisterConsumer("application/json", runtime.JSONConsumer())
	api.RegisterProducer("application/json", runtime.JSONProducer())

	api.APIKeyAuth = handlers.APIKeyAuthHandler()

	api.MessagesPostChatroomsChatroomHashMessagesHandler = handlers.MessagesPostMessageHandler()
	api.ChatroomsPostChatroomsHandler = handlers.ChatRoomsPostChatroomsHandler()
	api.ChatroomsGetChatroomsIDHandler = handlers.ChatroomsGetChatroomHandler()
	api.ChatroomsPostChatroomsChannelHashReadHandler = handlers.PostChatroomsChannelHashReadHandler()
	api.AccountPostProfileHandler = handlers.AccountPostProfileHandler()
	api.AccountPostAuthHandler = handlers.AccountPostAuthHandler()
	api.DeployGetHealthHandler = handlers.DeployGetHealthHandler()
	api.MessagesGetChatroomsChatroomHashMessagesHandler = handlers.MessagesGetMessagesHandler()

	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.ConfigureAPI()
	server.ConfigureFlags()

	server.Port = 80

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
