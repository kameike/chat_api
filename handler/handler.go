package handler

import (
	"net/http"

	"github.com/kameike/chat_api/apierror"
	"github.com/kameike/chat_api/model"
	"github.com/kameike/chat_api/swggen/apimodel"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/repository"
	"github.com/kameike/chat_api/swggen/restapi/operations/account"
	"github.com/kameike/chat_api/swggen/restapi/operations/chatrooms"
	"github.com/kameike/chat_api/swggen/restapi/operations/deploy"
	"github.com/kameike/chat_api/swggen/restapi/operations/messages"
)

type RequestHandlable interface {
	DeployGetHealthHandler() deploy.GetHealthHandlerFunc

	AccountPostAuthHandler() account.PostAuthHandlerFunc
	AccountPostProfileHandler() account.PostProfileHandlerFunc
	APIKeyAuthHandler() func(string) (interface{}, error)

	// ChatRoomsGetChatroomsIDMessagesHandler() chatrooms.GetAdminSearchChatroomsHandlerFunc
	// ChatRoomsGetChatroomsIDHandler() chatrooms.GetChatroomsIDHandlerFunc
	// ChatRoomsPostChatroomsIDReadHandler() chatrooms.PostChatroomsChannelHashReadHandlerFunc

	ChatRoomsPostChatroomsHandler() chatrooms.PostChatroomsHandlerFunc
	ChatRoomsPostChatroomsIDMessagesHandler() messages.PostChatroomsChatroomHashMessagesHandlerFunc
}

func SetUpHandler() RequestHandlable {
	return &appRequestHandler{
		p: repository.CreateAppRepositoryProvider(),
	}
}

type appRequestHandler struct {
	p repository.ReposotryProvidable
}

func (a *appRequestHandler) APIKeyAuthHandler() func(string) (interface{}, error) {
	return func(token string) (interface{}, error) {
		return a.AuthUser(token)
	}
}

func (a *appRequestHandler) AuthUser(token string) (*model.User, apierror.ChatAPIError) {
	repo := a.p.AuthRepository()
	return repo.FindUser(token)
}

func (a *appRequestHandler) AccountPostAuthHandler() account.PostAuthHandlerFunc {
	return func(params account.PostAuthParams) middleware.Responder {
		repo := a.p.AuthRepository()
		user, authInfo, err := repo.FindOrCreateUser(
			params.Body.AuthToken,
			params.Body.AccountHash,
		)

		if err != nil {
			return errorResponse(err)
		}

		res := account.NewPostAuthOK().WithPayload(&apimodel.AuthInfo{
			Account: &apimodel.Account{
				Name:     user.Name,
				ImageURL: user.Url,
				ID:       int64(user.ID),
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
		println("account-update")
		user := principal.(*model.User)
		repo, err := a.p.UserRepository(*user)

		if err != nil {
			return errorResponse(err)
		}

		user, err = repo.UpdateUser(userUpdateData{
			name: &params.Body.Name,
			url:  &params.Body.ImageURL,
		})

		if err != nil {
			return errorResponse(err)
		}

		return account.NewPostProfileOK().WithPayload(&apimodel.Account{
			Name:     user.Name,
			ImageURL: user.Url,
			ID:       int64(user.ID),
			Hash:     user.UserHash,
		})

	}
}

// func (a *appRequestHandler) ChatRoomsGetChatroomsIDMessagesHandler() chatrooms.GetChatroomsIDMessagesHandlerFunc {
// 	panic("not implemented")
// }
//
// func (a *appRequestHandler) ChatRoomsGetChatroomsIDHandler() chatrooms.GetChatroomsIDHandlerFunc {
// 	panic("not implemented")
// }

func (a *appRequestHandler) ChatRoomsPostChatroomsHandler() chatrooms.PostChatroomsHandlerFunc {
	return chatrooms.PostChatroomsHandlerFunc(func(params chatrooms.PostChatroomsParams, principal interface{}) middleware.Responder {
		u := principal.(*model.User)
		repo, err := a.p.UserRepository(*u)
		if err != nil {
		}

		rooms, err := repo.GetChatRooms(chatRoomMapper{params})

		if err != nil {
			return errorResponse(err)
		}

		chatroomsResult := make([]*apimodel.Chatroom, len(rooms), len(rooms))

		for i, r := range rooms {
			data := apimodel.Chatroom{
				ID:           r.RoomHash,
				Accounts:     mapUsers(r.Users),
				Messages:     []*apimodel.Message{},
				Name:         r.Name,
				UnreadsCount: []*apimodel.UnreadCount{},
			}
			chatroomsResult[i] = &data
		}

		result := &chatrooms.PostChatroomsOKBody{
			Chatrooms: chatroomsResult,
		}

		return chatrooms.NewPostChatroomsOK().WithPayload(result)
	})
}

func mapUser(user model.User) *apimodel.Account {
	return &apimodel.Account{
		Name:     user.Name,
		ImageURL: user.Url,
		ID:       int64(user.ID),
		Hash:     user.UserHash,
	}
}

func mapUsers(users []model.User) []*apimodel.Account {
	result := make([]*apimodel.Account, len(users), len(users))

	for i, u := range users {
		result[i] = mapUser(u)
	}
	return result
}

type chatRoomMapper struct {
	data chatrooms.PostChatroomsParams
}

func (d chatRoomMapper) RoomHashes() []string {
	return d.data.Body.Chatrooms
}

func (a *appRequestHandler) ChatRoomsPostChatroomsIDMessagesHandler() messages.PostChatroomsChatroomHashMessagesHandlerFunc {
	return func(params messages.PostChatroomsChatroomHashMessagesParams, principal interface{}) middleware.Responder {
		u := principal.(*model.User)

		if u == nil {
			return errorResponse(apierror.ErrorNoPermission())
		}

		cr, err := a.p.ChatRepository(*u, params.ChatroomHash)
		if err != nil {
			return errorResponse(err)
		}

		cr.CreateMessage(params.Body.Content)

		// TODO: correct message here
		res := &apimodel.MessagesResponse{
			Messages: []*apimodel.Message{},
		}

		return messages.NewPostChatroomsChatroomHashMessagesOK().WithPayload(res)
	}
}

// func (a *appRequestHandler) ChatRoomsPostChatroomsIDReadHandler() chatrooms.PostChatroomsIDReadHandlerFunc {
// 	panic("not implemented")
// }

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

func errorResponse(err apierror.ChatAPIError) middleware.ResponderFunc {
	return errorResponseWithCode(500, err.Localize())
}

func notHealthy(message string) middleware.ResponderFunc {
	return func(res http.ResponseWriter, pro runtime.Producer) {
		res.WriteHeader(503)
		res.Write([]byte(message))
	}
}
