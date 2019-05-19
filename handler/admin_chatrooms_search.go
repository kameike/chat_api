package handler

import (
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/swggen/restapi/operations/chatrooms"
)

func (a *RequestHandler) ChatroomsGetAdminSearchChatroomsHandler() chatrooms.GetAdminSearchChatroomsHandler {
	return chatrooms.GetAdminSearchChatroomsHandlerFunc(func(params chatrooms.GetAdminSearchChatroomsParams, principal interface{}) middleware.Responder {
		// TODO
		return middleware.NotImplemented("operation chatrooms.GetAdminSearchChatrooms has not yet been implemented")
	})

}
