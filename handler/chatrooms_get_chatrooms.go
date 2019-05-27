package handler

import (
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/apierror"
	"github.com/kameike/chat_api/model"
	"github.com/kameike/chat_api/repository"
	"github.com/kameike/chat_api/swggen/apimodel"
	"github.com/kameike/chat_api/swggen/restapi/operations/chatrooms"
)

func (a *RequestHandler) ChatroomsGetChatroomHandler() chatrooms.GetChatroomsIDHandler {
	return chatrooms.GetChatroomsIDHandlerFunc(func(params chatrooms.GetChatroomsIDParams, user interface{}) middleware.Responder {
		u := user.(*model.User)
		print("start")
		print(u)

		if u == nil {
			return errorResponse(apierror.NewError(apierror.INVALID_POST_MESSAGE))
		}

		cr, err := a.p.UserRepository(*u)
		if err != nil {
			return errorResponse(err)
		}

		r, err := cr.GetChatRoom(repository.GetChatRoomRequest{
			Hash: params.ID,
		})

		cr.GetMessages(repository.GetMessageRequest{
			RoomHash: r.RoomHash,
			User:     *u,
		})

		data := apimodel.Chatroom{
			ID:           int64(r.ID),
			Hash:         r.RoomHash,
			Accounts:     mapUsers(r.Users),
			Messages:     mapMessages(r.Messages),
			Name:         r.Name,
			UnreadsCount: []*apimodel.UnreadCount{},
			ReadAts:      apimodel.ReadAts{},
		}

		print("ok")

		return chatrooms.NewGetChatroomsIDOK().WithPayload(&data)
	})
}
