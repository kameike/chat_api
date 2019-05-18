package handler

import (
	"github.com/kameike/chat_api/swggen/apimodel"

	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/swggen/restapi/operations/account"
)

func (a *RequestHandler) AccountPostAuthHandler() account.PostAuthHandlerFunc {
	return func(params account.PostAuthParams) middleware.Responder {
		repo := a.p.AuthRepository()
		user, authInfo, err := repo.FindOrCreateUser(
			params.Body.AuthToken,
			params.Body.AccountHash,
		)

		if err != nil {
			return errorResponse(err)
		}

		res := account.NewPostAuthOK().WithPayload(&apimodel.AuthInfo{
			Account: &apimodel.Account{
				Name:     user.Name,
				ImageURL: user.Url,
				ID:       int64(user.ID),
				Hash:     user.UserHash,
			},
			AccessToken: authInfo.AccessToken,
		})

		return res
	}
}
