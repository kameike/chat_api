package main

import (
	"log"

	"github.com/go-openapi/loads"
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

	api.DeployGetHealthHandler = handlers.DeployGetHealthHandler()

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 1323

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
