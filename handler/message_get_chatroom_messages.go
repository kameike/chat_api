package handler

import (
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/apierror"
	"github.com/kameike/chat_api/model"
	"github.com/kameike/chat_api/swggen/apimodel"
	"github.com/kameike/chat_api/swggen/restapi/operations/messages"
)

func (a *RequestHandler) MessagesGetMessagesHandler() messages.GetChatroomsChatroomHashMessagesHandler {
	return messages.GetChatroomsChatroomHashMessagesHandlerFunc(func(request messages.GetChatroomsChatroomHashMessagesParams, user interface{}) middleware.Responder {
		u := user.(*model.User)

		if u == nil {
			return errorResponse(apierror.NewError(apierror.INVALID_USER))
		}

		repo, err := a.p.ChatRepository(*u, request.ChatroomHash)

		if err != nil {
			return errorResponse(apierror.Error(apierror.CHATROOM_NOT_FOUND, err))
		}

		result, err := repo.GetMessageAndReadStatus()

		if err != nil {
			return errorResponse(apierror.Error(apierror.CHATROOM_NOT_FOUND, err))
		}

		msgs := make([]*apimodel.Message, len(result.Messages), len(result.Messages))

		for i, m := range result.Messages {
			apimsg := apimodel.Message{
				Content:   m.Text,
				CreatedAt: m.CreatedAt.Unix(),
				ID:        int64(m.ID),
				Account:   mapUser(*u),
			}

			msgs[i] = &apimsg
		}

		// TODO
		response := apimodel.MessagesResponse{
			Messages: []*apimodel.Message{},
			ReadAts:  apimodel.ReadAts{},
		}

		return messages.NewPostChatroomsChatroomHashMessagesOK().WithPayload(&response)
	})
}
