package handler

import (
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/swggen/restapi/operations/chatrooms"
)

func (a *RequestHandler) ChatroomsGetChatroomHandler() chatrooms.GetChatroomsIDHandler {
	return chatrooms.GetChatroomsIDHandlerFunc(func(params chatrooms.GetChatroomsIDParams, user interface{}) middleware.Responder {
		// TODO
		return middleware.NotImplemented("TODO")
	})
}
