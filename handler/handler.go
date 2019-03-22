package handler

import (
	"fmt"
	"net/http"

	"github.com/kameike/chat_api/error"
	"github.com/kameike/chat_api/model"
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
	AuthUser(token string) (*model.User, error.ChatAPIError)

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

func (a *appRequestHandler) AuthUser(token string) (*model.User, error.ChatAPIError) {
	repo := a.p.AuthRepository()
	return repo.FindUser(token)
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

type userUpdateData struct {
	name *string
	url  *string
}

func (d userUpdateData) Name() *string {
	return d.name
}

func (d userUpdateData) ImageURL() *string {
	return d.url
}

func (a *appRequestHandler) AccountPostProfileHandler() account.PostProfileHandlerFunc {
	return func(params account.PostProfileParams, principal interface{}) middleware.Responder {
		user := principal.(*model.User)
		repo, err := a.p.UserRepository(*user)

		if err != nil {
			return errorResponse(err)
		}

		user, err = repo.UpdateUser(userUpdateData{
			name: params.Name,
			url:  params.ImageURL,
		})

		if err != nil {
			return errorResponse(err)
		}

		return account.NewPostProfileOK().WithPayload(&apimodel.User{
			Name:     user.Name,
			ImageURL: user.Url,
			ID:       fmt.Sprint(user.ID),
			Hash:     user.UserHash,
		})

	}
}

func (a *appRequestHandler) ChatRoomsGetChatroomsIDMessagesHandler() chat_rooms.GetChatroomsIDMessagesHandlerFunc {
	panic("not implemented")
}

func (a *appRequestHandler) ChatRoomsGetChatroomsIDHandler() chat_rooms.GetChatroomsIDHandlerFunc {
	panic("not implemented")
}

func (a *appRequestHandler) ChatRoomsPostChatroomsHandler() chat_rooms.PostChatroomsHandlerFunc {
	return chat_rooms.PostChatroomsHandlerFunc(func(params chat_rooms.PostChatroomsParams, principal interface{}) middleware.Responder {
		u := principal.(*model.User)
		repo, err := a.p.UserRepository(*u)
		if err != nil {
			return errorResponse(err)
		}

		rooms, err := repo.GetChatRooms(chatRoomMapper{params})
		result := make([]*apimodel.Chatroom, len(rooms), len(rooms))

		for i, r := range rooms {
			data := apimodel.Chatroom{
				ID:           r.RoomHash,
				Participants: mapUsers(r.Users),
				PeekedChat:   []*apimodel.Message{},
				Unreads:      []*apimodel.ChatroomUnreadsItems0{},
			}
			result[i] = &data
		}

		return chat_rooms.NewPostChatroomsOK().WithPayload(result)
	})
}

func mapUser(user model.User) *apimodel.User {
	return &apimodel.User{
		Name:     user.Name,
		ImageURL: user.Url,
		ID:       fmt.Sprint(user.ID),
		Hash:     user.UserHash,
	}
}

func mapUsers(users []model.User) []*apimodel.User {
	result := make([]*apimodel.User, len(users), len(users))

	for i, u := range users {
		result[i] = mapUser(u)
	}
	return result
}

type chatRoomMapper struct {
	data chat_rooms.PostChatroomsParams
}

func (d chatRoomMapper) RoomHashes() []string {
	return d.data.Body.Request
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
