package repository

import (
	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/error"
	"github.com/kameike/chat_api/model"
)

type AuthRepositable interface {
	FindOrCreateUser(token string, hash string) (*model.User, *model.AccessToken, error.ChatAPIError)
	FindUser(token string) (*model.User, error.ChatAPIError)
}

type UserUpdateInfoDescriable interface {
	Name() *string
	ImageURL() *string
}

type ChatRoomsInfoDescriable interface {
	RoomHashes() []string
}

type UserRepositable interface {
	UpdateUser(UserUpdateInfoDescriable) (*model.User, error.ChatAPIError)
	GetChatRooms(ChatRoomsInfoDescriable) (*model.ChatRoom, error.ChatAPIError)
}

type ChatRepostitable interface {
	getChatRoom()
	postMessage()
	getMessages()
	getUnreads()
}

type ReposotryProvidable interface {
	CheckHealth() (string, bool)
	AuthRepository() AuthRepositable
	UserRepository(model.User) (UserRepositable, error.ChatAPIError)
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

func (r *applicationRepositoryProvidable) AuthRepository() AuthRepositable {
	repo := authRepository{
		d: r.datasource,
	}

	return &repo
}

func (r *applicationRepositoryProvidable) CheckHealth() (string, bool) {
	return r.datasource.CheckHealth()
}

func (r *applicationRepositoryProvidable) Close() {
	r.datasource.Close()
}

func (r *applicationRepositoryProvidable) UserRepository(user model.User) (UserRepositable, error.ChatAPIError) {
	u := userRepository{
		user: user,
		ds:   r.datasource,
	}
	return &u, nil
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
