package handler

import (
	"github.com/kameike/chat_api/model"
	"github.com/kameike/chat_api/repository"
	"github.com/kameike/chat_api/swggen/apimodel"

	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/swggen/restapi/operations/chatrooms"
)

func (a *RequestHandler) ChatRoomsPostChatroomsHandler() chatrooms.PostChatroomsHandlerFunc {
	return chatrooms.PostChatroomsHandlerFunc(func(params chatrooms.PostChatroomsParams, principal interface{}) middleware.Responder {
		u := principal.(*model.User)
		repo, err := a.p.UserRepository(*u)
		if err != nil {
		}

		rooms, err := repo.GetChatRooms(repository.ChatRoomsInfoDescriable{params.Body.Chatrooms})

		if err != nil {
			return errorResponse(err)
		}

		chatroomsResult := make([]*apimodel.Chatroom, len(rooms), len(rooms))

		for i, r := range rooms {
			data := apimodel.Chatroom{
				ID:           "",
				Hash:         r.RoomHash,
				Accounts:     mapUsers(r.Users),
				Messages:     mapMessages(r.Messages),
				Name:         r.Name,
				UnreadsCount: []*apimodel.UnreadCount{},
				ReadAts:      apimodel.ReadAts{},
			}
			chatroomsResult[i] = &data
		}

		result := &chatrooms.PostChatroomsOKBody{
			Chatrooms: chatroomsResult,
		}

		return chatrooms.NewPostChatroomsOK().WithPayload(result)
	})
}
