package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	// "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/handler"
	"github.com/kameike/chat_api/repository"
	"github.com/kameike/chat_api/swggen/restapi"
	"github.com/kameike/chat_api/swggen/restapi/operations"
	// "github.com/kameike/chat_api/swggen/restapi/operations/deploy"
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

	api.MessagesPostChatroomsChatroomHashMessagesHandler = handlers.ChatRoomsPostChatroomsIDMessagesHandler()
	api.ChatroomsPostChatroomsHandler = handlers.ChatRoomsPostChatroomsHandler()
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
