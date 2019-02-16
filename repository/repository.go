package repository

import (
	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/error"
	"github.com/kameike/chat_api/model"
)

type UserRepositable interface {
	createUser(user model.User)
	updateToken(user model.User)
}

type ChatRepostitable interface {
	createChatRoom()
	getChatRoom()
	postMessage()
	getMessages()
	getUnreads()
}

type ReposotryProvidable interface {
	CheckHealth() (string, bool)
	UserRepository() (UserRepositable, error.ChatAPIError)
	ChatRepository() (ChatRepostitable, error.ChatAPIError)
	Close()
}

type applicationRepositoryProvidable struct {
	datasource datasource.DataSourceDescriptor
}

func CreateAppRepositoryProvider() ReposotryProvidable {
	ds := datasource.PrepareDatasource()
	return &applicationRepositoryProvidable{
		datasource: ds,
	}
}

func (r *applicationRepositoryProvidable) CheckHealth() (string, bool) {
	return r.datasource.CheckHealth()
}

func (r *applicationRepositoryProvidable) Close() {
	r.datasource.Close()
}

func (r *applicationRepositoryProvidable) UserRepository() (UserRepositable, error.ChatAPIError) {
	return nil, nil
}

func (r *applicationRepositoryProvidable) ChatRepository() (ChatRepostitable, error.ChatAPIError) {
	return nil, nil
}

type AuthInfoProvidable interface {
	RefreshToken() string
	Name() string
}

type AuthAccessRequestable interface {
	AccessToken() string
}
