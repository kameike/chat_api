package repository

import (
	"github.com/kameike/chat_api/apierror"
	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/model"
)

// AuthRepositable Support Auth Repository
type AuthRepositable interface {
	FindOrCreateUser(token string, hash string) (*model.User, *model.AccessToken, apierror.ChatAPIError)
	FindUser(token string) (*model.User, apierror.ChatAPIError)
}

// UserUpdateInfoDescriable Require users info which can be updated
type UserUpdateInfoDescriable interface {
	Name() *string
	ImageURL() *string
}

//ChatRoomsInfoDescriable support signed json strings
type ChatRoomsInfoDescriable struct {
	RoomHashes []string
}

// GetChatRoomRequest exprain chat room with long uniqe hashed string
type GetChatRoomRequest struct {
	Hash string
}

// UserRepositable with can be dispended from  repository providable
type UserRepositable interface {
	UpdateUser(UserUpdateInfoDescriable) (*model.User, apierror.ChatAPIError)
	GetChatRooms(ChatRoomsInfoDescriable) ([]*model.ChatRoom, apierror.ChatAPIError)
	CreateMessage(CreateMessageRequest) apierror.ChatAPIError
	GetMessages(GetMessageRequest) ([]*model.Message, apierror.ChatAPIError)
	GetChatRoom(GetChatRoomRequest) (*model.ChatRoom, apierror.ChatAPIError)
}

type Unread struct {
	UserUnreads map[string]int
}

type ChatRepositable interface {
	CreateMessage(string) apierror.ChatAPIError
}

type ReposotryProvidable interface {
	CheckHealth() (string, bool)
	AuthRepository() AuthRepositable
	UserRepository(model.User) (UserRepositable, apierror.ChatAPIError)
	ChatRepository(model.User, string) (ChatRepositable, apierror.ChatAPIError)
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

func (r *applicationRepositoryProvidable) UserRepository(user model.User) (UserRepositable, apierror.ChatAPIError) {
	u := userRepository{
		user: user,
		ds:   r.datasource,
	}
	return &u, nil
}

func (r *applicationRepositoryProvidable) ChatRepository(u model.User, chatId string) (ChatRepositable, apierror.ChatAPIError) {
	ur, err := r.UserRepository(u)
	if err != nil {
		return nil, err
	}

	res, err := ur.GetChatRoom(GetChatRoomRequest{
		Hash: chatId,
	})

	if err != nil {
		return nil, err
	}

	repo := chatRepository{
		ds:   r.datasource,
		room: *res,
		user: u,
	}

	return &repo, nil
}

type AuthInfoProvidable interface {
	RefreshToken() string
	Name() string
}

type AuthAccessRequestable interface {
	AccessToken() string
}
