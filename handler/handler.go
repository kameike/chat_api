package handler

import (
	"fmt"
	"net/http"

	"github.com/kameike/chat_api/error"
	"github.com/kameike/chat_api/swggen/apimodel"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/repository"
	"github.com/kameike/chat_api/swggen/restapi/operations/account"
	"github.com/kameike/chat_api/swggen/restapi/operations/chat_rooms"
	"github.com/kameike/chat_api/swggen/restapi/operations/deploy"
)

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
	return &appRequestHandler{
		p: repository.CreateAppRepositoryProvider(),
	}
}

type appRequestHandler struct {
	p repository.ReposotryProvidable
}

func (a *appRequestHandler) AccountPostAuthHandler() account.PostAuthHandlerFunc {
	return func(params account.PostAuthParams) middleware.Responder {
		repo := a.p.AuthRepository()
		user, authInfo, err := repo.FindOrCreateUser(params.AuthToken, params.UserHash)

		if err != nil {
			return errorResponse(err)
		}

		res := account.NewPostAuthOK().WithPayload(&apimodel.AuthInfo{
			User: &apimodel.User{
				Name:     user.Name,
				ImageURL: user.Url,
				ID:       fmt.Sprint(user.ID),
				Hash:     user.UserHash,
			},
			AccessToken: authInfo.AccessToken,
		})

		return res
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
		repo := h.p
		message, isHealthy := repo.CheckHealth()
		if isHealthy {
			return deploy.NewGetHealthOK().WithPayload(message)
		} else {
			return notHealthy(message)
		}
	}
}

// TODO: いい感じにエラーを整形する
func errorResponseWithCode(code int, message string) middleware.ResponderFunc {
	return func(res http.ResponseWriter, pro runtime.Producer) {
		res.WriteHeader(code)
		res.Write([]byte(message))
	}
}

func errorResponse(err error.ChatAPIError) middleware.ResponderFunc {
	return errorResponseWithCode(500, err.Localize())
}

func notHealthy(message string) middleware.ResponderFunc {
	return func(res http.ResponseWriter, pro runtime.Producer) {
		res.WriteHeader(503)
		res.Write([]byte(message))
	}
}
