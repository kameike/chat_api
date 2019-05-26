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
			"roomName": "room12"
	}`, user.UserHash, user2.UserHash)

	target2 := fmt.Sprintf(`{
			"accountHashList": ["%s", "%s"],
			"roomName": "room21"
	}`, user.UserHash, user2.UserHash)

	print(target1, target2)

	roomRequest := ChatRoomsInfoDescriable{
		RoomHashes: []string{target1, target2},
	}

	r, err := ur.GetChatRooms(roomRequest)

	if err != nil {
		panic(err.Error())
	}

	chatroom = r[0]

	print(user2, ur)

	authUser = *user
}

var chatroom *model.ChatRoom
var otherChatroom *model.ChatRoom

func afterChat() {
	generalAfter()
}

func Test_hoge(t *testing.T) {
	// beforeChat()
	// defer afterChat()
}
