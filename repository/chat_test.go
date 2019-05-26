package repository

import (
	"fmt"
	"testing"

	"github.com/kameike/chat_api/model"
)

func beforeChat() {
	generalBefore()
	user, _, _ := provider.AuthRepository().FindOrCreateUser(token, hash)
	user2, _, _ := provider.AuthRepository().FindOrCreateUser("test", "kameike")

	ur, _ := provider.UserRepository(*user)

	target1 := fmt.Sprintf(`{
			"accountHashList": ["%s", "%s"],
			"channelName": "room12"
	}`, user.UserHash, user2.UserHash)

	target2 := fmt.Sprintf(`{
			"accountHashList": ["%s", "%s"],
			"channelName": "room21"
	}`, user.UserHash, user2.UserHash)

	roomRequest := ChatRoomsInfoDescriable{
		RoomHashes: []string{target1, target2},
	}

	r, err := ur.GetChatRooms(roomRequest)

	if err != nil {
		panic(err.Error())
	}

	chatroom = r[0]
	otherChatroom = r[1]

	authUser = *user
	chatRepo, err = provider.ChatRepository(*user, chatroom.RoomHash)

	if err != nil {
		panic(err.Error())
	}
}

var chatroom *model.ChatRoom
var otherChatroom *model.ChatRoom
var chatRepo ChatRepositable

func afterChat() {
	generalAfter()
}

func Testメッセージを作成できる(t *testing.T) {
	beforeChat()
	defer afterChat()

	ds.RDB().LogMode(true)
	err := chatRepo.CreateMessage("test")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = chatRepo.CreateMessage("eee")
	if err != nil {
		t.Fatal(err.Error())
	}
	err = chatRepo.CreateMessage("🤗")
	if err != nil {
		t.Fatal(err.Error())
	}
	ds.RDB().LogMode(false)
}

func Test_hoge(t *testing.T) {
	beforeChat()
	defer afterChat()
}
