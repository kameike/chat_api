package service

import (
	"github.com/kameike/chat_api/error"
	"github.com/kameike/chat_api/model"
	"github.com/kameike/chat_api/repository"
)

type AuthInfo interface {
	AuthToken() string
	UserHash() string
}

type UserService interface {
	singInOrSingUp(AuthInfo) (*model.User, error.ChatAPIError)
}

type appUserService struct {
	repositoryProvider repository.ReposotryProvidable
}
