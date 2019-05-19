package handler

import (
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/swggen/restapi/operations/chatrooms"
)

func (a *RequestHandler) PostChatroomsChannelHashReadHandler() chatrooms.PostChatroomsChannelHashReadHandler {
	return chatrooms.PostChatroomsChannelHashReadHandlerFunc(func(params chatrooms.PostChatroomsChannelHashReadParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation chatrooms.PostChatroomsChannelHashRead has not yet been implemented")
	})

}
