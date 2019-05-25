package handler

import (
	"github.com/kameike/chat_api/apierror"
	"github.com/kameike/chat_api/model"
	"github.com/kameike/chat_api/swggen/apimodel"

	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/swggen/restapi/operations/messages"
)

func (a *RequestHandler) MessagesPostMessageHandler() messages.PostChatroomsChatroomHashMessagesHandlerFunc {
	return func(params messages.PostChatroomsChatroomHashMessagesParams, principal interface{}) middleware.Responder {
		u := principal.(*model.User)

		if u == nil {
			return errorResponse(apierror.NewError(apierror.INVALID_POST_MESSAGE))
		}

		cr, err := a.p.ChatRepository(*u, params.ChatroomHash)
		if err != nil {
			return errorResponse(err)
		}

		err = cr.CreateMessage(params.Body.Content)
		if err != nil {
			return errorResponse(err)
		}

		// TODO: correct message here
		res := &apimodel.MessagesResponse{
			Messages: []*apimodel.Message{},
			ReadAts:  apimodel.ReadAts{},
		}

		return messages.NewPostChatroomsChatroomHashMessagesOK().WithPayload(res)
	}
}
