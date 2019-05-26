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

	db := ds.RDB()

	beforeCount := 0
	db.Model(&model.Message{}).Count(&beforeCount)

	err := chatRepo.CreateMessage("test")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = chatRepo.CreateMessage("😄")
	if err != nil {
		t.Fatal(err.Error())
	}

	afterCount := 0
	db.Model(&model.Message{}).Count(&afterCount)

	if afterCount-beforeCount != 2 {
		t.Fatalf("invalid count before %d after %d", afterCount, beforeCount)
	}
}

func Testメッセージを取得できる(t *testing.T) {
	beforeChat()
	defer afterChat()

	chatRepo.CreateMessage("hey")
	chatRepo.CreateMessage("hey")
	chatRepo.CreateMessage("hey")

	res, err := chatRepo.GetMessageAndReadStatus()

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(res.Messages) != 3 {
		t.Fatalf("count should be 3 bat %d", len(res.Messages))
	}

	u := res.Messages[0].User

	//check user preload
	if u.UserHash != hash {
		t.Fatalf("user has is %s", u.UserHash)
	}
}

func Test違うチャットルームのメッセージには関与しない(t *testing.T) {
	beforeChat()
	defer afterChat()

	chatRepo.CreateMessage("hey1")
	chatRepo.CreateMessage("hey3")
	chatRepo.CreateMessage("hey3")

	//違うの作る
	user, _, _ := provider.AuthRepository().FindOrCreateUser(token, hash)
	otherRepo, err := provider.ChatRepository(*user, otherChatroom.RoomHash)
	otherRepo.CreateMessage("hey")

	res, err := chatRepo.GetMessageAndReadStatus()

	if err != nil {
		t.Fatal(err.Error())
	}

	if len(res.Messages) != 3 {
		t.Fatalf("count should be 3 bat %d", len(res.Messages))
	}
}
