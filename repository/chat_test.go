package repository

import (
	"testing"

	"github.com/kameike/chat_api/model"
)

func beforeChat() {
	generalBefore()
	user, _, _ := provider.AuthRepository().FindOrCreateUser(token, hash)
	user2, _, _ := provider.AuthRepository().FindOrCreateUser("test", "hoge")

	ur, _ := provider.UserRepository(*user)

	print(user2, ur)

	authUser = *user
}

var chatRoom model.ChatRoom

func afterChat() {
	generalAfter()
}

func Test_hoge(t *testing.T) {

}
