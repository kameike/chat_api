package repository

import (
	"testing"

	"github.com/kameike/chat_api/model"
)

func beforeUser() {
	generalBefore()

	user, _, _ := provider.AuthRepository().FindOrCreateUser(token, hash)
	authUser = *user
}

var authUser model.User

func afterUser() {
	generalAfter()
}

func Test_userRepository_UpdateUser(t *testing.T) {
	beforeUser()
	defer afterUser()
	u, _ := provider.UserRepository(authUser)
	user, err := u.UpdateUser(testAuthInfo{})

	if err != nil {
		t.Fatalf(err.Error())
	}
	if user.Url != "url" || user.Name != "name" {
		t.Fatalf("update info is not good")
	}

	user1, _, _ := provider.AuthRepository().FindOrCreateUser(token, hash)

	if user1.Url != "url" || user1.Name != "name" {
		t.Fatalf("db not updated '%s'", user1.Url)
	}
}

func Testチャットルームが存在しなくても作られる(t *testing.T) {
	beforeUser()
	defer afterUser()
	u, _ := provider.UserRepository(authUser)

	var roomSign []string
	roomSign = append(roomSign, `{
			"users": ["hogehohge", "fugafuga"],
			"roomId": "hoge",
			"roomName": "fuga",
		}
	`)

	testData := testChatRoomCreateInfo{
		data: roomSign,
	}

	result, err := u.GetChatRooms(testData)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if result == nil {
		t.Fatalf("chat room has not been created")
	}
}

type testAuthInfo struct{}

func (p testAuthInfo) Name() *string {
	data := "name"
	return &data
}

func (p testAuthInfo) ImageURL() *string {
	data := "url"
	return &data
}

type testChatRoomCreateInfo struct {
	data []string
}

func (t testChatRoomCreateInfo) RoomHashes() []string {
	return t.data
}
