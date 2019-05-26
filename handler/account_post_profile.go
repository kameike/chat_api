package handler

import (
	"github.com/kameike/chat_api/model"
	"github.com/kameike/chat_api/swggen/apimodel"

	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/swggen/restapi/operations/account"
)

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

func (a *RequestHandler) AccountPostProfileHandler() account.PostProfileHandlerFunc {
	return func(params account.PostProfileParams, principal interface{}) middleware.Responder {
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
