package handler

import (
	"log"
	"net/http"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/repository"
	"github.com/kameike/chat_api/swggen/restapi/operations/account"
	"github.com/kameike/chat_api/swggen/restapi/operations/chat_rooms"
	"github.com/kameike/chat_api/swggen/restapi/operations/deploy"
)

var repo repository.ReposotryProvidable

type RequestHandlable interface {
	DeployGetHealthHandler() deploy.GetHealthHandlerFunc

	AccountPostAuthHandler() account.PostAuthHandlerFunc
	AccountPostProfileHandler() account.PostProfileHandlerFunc

	ChatRoomsGetChatroomsIDMessagesHandler() chat_rooms.GetChatroomsIDMessagesHandlerFunc
	ChatRoomsGetChatroomsIDHandler() chat_rooms.GetChatroomsIDHandlerFunc
	ChatRoomsPostChatroomsHandler() chat_rooms.PostChatroomsHandlerFunc
	ChatRoomsPostChatroomsIDMessagesHandler() chat_rooms.PostChatroomsIDMessagesHandlerFunc
	ChatRoomsPostChatroomsIDReadHandler() chat_rooms.PostChatroomsIDReadHandlerFunc
}

func SetUpHandler() RequestHandlable {
	return &appRequestHandler{}
}

type appRequestHandler struct {
}

func (a *appRequestHandler) AccountPostAuthHandler() account.PostAuthHandlerFunc {
	return func(params account.PostAuthParams) middleware.Responder {
		res := account.NewPostAuth()
		return middleware.NotImplemented("not yet")
	}
}

func (a *appRequestHandler) AccountPostProfileHandler() account.PostProfileHandlerFunc {
	panic("not implemented")
}

func (a *appRequestHandler) ChatRoomsGetChatroomsIDMessagesHandler() chat_rooms.GetChatroomsIDMessagesHandlerFunc {
	panic("not implemented")
}

func (a *appRequestHandler) ChatRoomsGetChatroomsIDHandler() chat_rooms.GetChatroomsIDHandlerFunc {
	panic("not implemented")
}

func (a *appRequestHandler) ChatRoomsPostChatroomsHandler() chat_rooms.PostChatroomsHandlerFunc {
	panic("not implemented")
}

func (a *appRequestHandler) ChatRoomsPostChatroomsIDMessagesHandler() chat_rooms.PostChatroomsIDMessagesHandlerFunc {
	panic("not implemented")
}

func (a *appRequestHandler) ChatRoomsPostChatroomsIDReadHandler() chat_rooms.PostChatroomsIDReadHandlerFunc {
	panic("not implemented")
}

func (h *appRequestHandler) DeployGetHealthHandler() deploy.GetHealthHandlerFunc {
	return func(params deploy.GetHealthParams) middleware.Responder {
		message, isHealthy := repo.CheckHealth()
		if isHealthy {
			return deploy.NewGetHealthOK().WithPayload(message)
		} else {
			return notHealthy(message)
		}
	}
}

// TODO: いい感じにエラーを整形する
func errorResponse(code int, message string) middleware.ResponderFunc {
	return func(res http.ResponseWriter, pro runtime.Producer) {
		res.WriteHeader(code)
		res.Write([]byte(message))
	}
}

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
