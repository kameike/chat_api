package handler

import (
	"github.com/kameike/chat_api/apierror"
	"github.com/kameike/chat_api/model"
	"github.com/kameike/chat_api/swggen/apimodel"

	_ "github.com/joho/godotenv/autoload"
)

func (a *RequestHandler) APIKeyAuthHandler() func(string) (interface{}, error) {
	return func(token string) (interface{}, error) {
		return a.authUser(token)
	}
}

func (a *RequestHandler) authUser(token string) (*model.User, apierror.ChatAPIError) {
	repo := a.p.AuthRepository()
	return repo.FindUser(token)
}

func mapUser(user model.User) *apimodel.Account {
	return &apimodel.Account{
		Name:     user.Name,
		ImageURL: user.Url,
		ID:       int64(user.ID),
		Hash:     user.UserHash,
	}
}
