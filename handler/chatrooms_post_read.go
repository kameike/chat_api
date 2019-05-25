package handler

import (
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/apierror"
	"github.com/kameike/chat_api/model"
	"github.com/kameike/chat_api/swggen/apimodel"
	"github.com/kameike/chat_api/swggen/restapi/operations/chatrooms"
)

func (a *RequestHandler) PostChatroomsChannelHashReadHandler() chatrooms.PostChatroomsChannelHashReadHandler {
	return chatrooms.PostChatroomsChannelHashReadHandlerFunc(func(params chatrooms.PostChatroomsChannelHashReadParams, principal interface{}) middleware.Responder {
		u := principal.(*model.User)

		if u == nil {
			return errorResponse(apierror.NewError(apierror.INVALID_POST_MESSAGE))
		}

		_, err := a.p.UserRepository(*u)
		if err != nil {
			return errorResponse(err)
		}

		response := apimodel.MessagesResponse{
			Messages: []*apimodel.Message{},
			ReadAts:  apimodel.ReadAts{},
		}

		return chatrooms.NewPostChatroomsChannelHashReadOK().WithPayload(&response)
	})
}
