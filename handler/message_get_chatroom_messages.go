package handler

import (
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/swggen/apimodel"
	"github.com/kameike/chat_api/swggen/restapi/operations/messages"
)

func (a *RequestHandler) MessagesGetMessagesHandler() messages.GetChatroomsChatroomHashMessagesHandler {
	return messages.GetChatroomsChatroomHashMessagesHandlerFunc(func(request messages.GetChatroomsChatroomHashMessagesParams, user interface{}) middleware.Responder {
		// TODO

		response := apimodel.MessagesResponse{
			Messages: []*apimodel.Message{},
			ReadAts:  apimodel.ReadAts{},
		}

		return messages.NewPostChatroomsChatroomHashMessagesOK().WithPayload(&response)
	})
}
