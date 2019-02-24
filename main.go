package main

import (
	"log"
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/repository"
	"github.com/kameike/chat_api/swggen/restapi"
	"github.com/kameike/chat_api/swggen/restapi/operations"
	"github.com/kameike/chat_api/swggen/restapi/operations/deploy"
)

var repo repository.ReposotryProvidable

func main() {
	repo = repository.CreateAppRepositoryProvider()
	defer repo.Close()

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewChatAPI(swaggerSpec)

	api.DeployGetHealthHandler = healthHandler

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 1323

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

var healthHandler = deploy.GetHealthHandlerFunc(func(params deploy.GetHealthParams) middleware.Responder {
	message, isHealthy := repo.CheckHealth()
	if isHealthy {
		return deploy.NewGetHealthOK().WithPayload(message)
	} else {
		return notHealthy(message)
	}
})

func notHealthy(message string) middleware.ResponderFunc {
	return func(res http.ResponseWriter, pro runtime.Producer) {
		res.WriteHeader(503)
		res.Write([]byte(message))
	}
}

var hello = runtime.OperationHandlerFunc(func(params interface{}) (interface{}, error) {
	log.Println("received 'findTodos'")
	log.Printf("%#v\n", params)

	return "hello", nil
})
