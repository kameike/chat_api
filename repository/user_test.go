package repository

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/kameike/chat_api/model"
)

func beforeUser() {
	generalBefore()
	ds.RDB().LogMode(false)

	user, _, _ := provider.AuthRepository().FindOrCreateUser(token, hash)
	user2, _, _ := provider.AuthRepository().FindOrCreateUser("randomToken", "xxxxxxxxhashxxxxxx")
	authUser = *user
	opponentUser = *user2
}

var authUser model.User
var opponentUser model.User

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
	roomSign = append(roomSign, fmt.Sprintf(`{
			"accountHashList": ["%s"],
			"channelName": "fuga"
		}
	`, authUser.UserHash))

	testData := ChatRoomsInfoDescriable{
		RoomHashes: roomSign,
	}

	result, err := u.GetChatRooms(testData)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if result == nil {
		t.Fatalf("chat room has not been created")
	}
}

func Testルーム名が存在しない場合ダメ(t *testing.T) {
	beforeUser()
	defer afterUser()
	u, _ := provider.UserRepository(authUser)

	var roomSign []string
	roomSign = append(roomSign, fmt.Sprintf(`{
			"accountHashList": ["%s"]
		}
	`, authUser.UserHash))

	testData := ChatRoomsInfoDescriable{
		RoomHashes: roomSign,
	}

	_, err := u.GetChatRooms(testData)

	if err == nil {
		t.Fatalf("err should be happen")
	}
}

func Test同一データをフィルターできる(t *testing.T) {

	target := []chatRoomData{
		chatRoomData{
			Accounts: []string{"a", "b"},
			RoomName: "c",
		},
		chatRoomData{
			Accounts: []string{"a", "b"},
			RoomName: "c",
		},
	}

	if len(filterData(target)) != 1 {
		t.Fatal("failed to filter chatroom")
	}
}

func Test同一のチャットルームのリクエストが来たらよしなにマージされる(t *testing.T) {
	beforeUser()
	defer afterUser()
	u, _ := provider.UserRepository(authUser)

	var roomSign []string
	roomSign = append(roomSign, fmt.Sprintf(`{
			"accountHashList": ["%s"],
			"channelName": "fuga"
		}
	`, authUser.UserHash))
	roomSign = append(roomSign, fmt.Sprintf(`{
			"accountHashList": ["%s"],
			"channelName": "fuga"
		}
	`, authUser.UserHash))

	testData := ChatRoomsInfoDescriable{
		RoomHashes: roomSign,
	}

	result, err := u.GetChatRooms(testData)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if 1 != len(result) {
		t.Fatalf("chat room has not been created")
	}
}

func Test部屋名が異なると違うチャットルームができる(t *testing.T) {
	beforeUser()
	defer afterUser()
	u, _ := provider.UserRepository(authUser)

	var roomSign []string
	roomSign = append(roomSign, fmt.Sprintf(`{
			"accountHashList": ["%s"],
			"channelName": "fuga"
		}
	`, authUser.UserHash))
	roomSign = append(roomSign, fmt.Sprintf(`{
			"accountHashList": ["%s"],
			"channelName": "hoge"
		}
	`, authUser.UserHash))

	testData := ChatRoomsInfoDescriable{
		RoomHashes: roomSign,
	}

	result, err := u.GetChatRooms(testData)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if 2 != len(result) {
		t.Fatalf("chat room has not been created")
	}
}

func Testチャットルームの型によってチャットを作ることができる(t *testing.T) {
	data := []chatRoomData{
		chatRoomData{
			Accounts: []string{
				authUser.UserHash,
				opponentUser.UserHash,
			},
			RoomName: "roomid",
		},
	}

	beforeUser()
	defer afterUser()
	repo, _ := provider.UserRepository(authUser)
	app := repo.(*userRepository)

	beforeCount := 0
	beforeRelationCount := 0
	ds.RDB().Model(&model.ChatRoom{}).Count(&beforeCount)
	ds.RDB().Model(&model.UserChatRoom{}).Count(&beforeRelationCount)

	res, err := app.getChatrooms(data)

	afterCount := 0
	afterRelationCount := 0
	ds.RDB().Model(&model.ChatRoom{}).Count(&afterCount)
	ds.RDB().Model(&model.UserChatRoom{}).Count(&afterRelationCount)

	if err != nil {
		t.Fail()
	}

	if len(res) != 1 {
		t.Fail()
	}

	if afterCount-beforeCount != 1 {
		t.Fail()
	}

	if afterRelationCount-beforeRelationCount != 2 {
		t.Fatalf("failed to make relation")
	}

	t.Run("ハッシュからチャットルームを見つけることができる", func(t *testing.T) {
		hash := res[0].RoomHash
		res, err := app.GetChatRoom(GetChatRoomRequest{hash})

		if err != nil {
			t.Fatalf(err.Error())
		}

		err = app.CreateMessage(CreateMessageRequest{
			User:    app.user,
			Room:    *res,
			Message: "test",
		})

		if err != nil {
			t.Fatalf(err.Error())
		}

		res, err = app.GetChatRoom(GetChatRoomRequest{hash})

		if err != nil {
			t.Fatalf(err.Error())
		}

		if res == nil {
			t.Fatalf("result must not be nil")
		}

		if len(res.Users) != 2 {
			t.Fatalf("faild to preload user")
		}

		if len(res.Messages) != 1 {
			t.Fatalf("faild to preload message %d", len(res.Messages))
		}

		found := false
		for _, u := range res.Users {
			if u.ID == authUser.ID {
				found = true
			}
		}

		if !found {
			t.Fatalf("user not preloaded")
		}
	})

	t.Run("チャットルームに含まれないユーザーをもったリポジトリからは発見できない", func(t *testing.T) {
		hash := res[0].RoomHash
		app.user = model.User{}

		_, err := app.GetChatRoom(GetChatRoomRequest{hash})

		if err == nil {
			t.Fatalf("err should not be nil")
		}
	})
}
func Testチャットルームが複数作られる(t *testing.T) {
	data := []chatRoomData{
		chatRoomData{
			Accounts: []string{
				authUser.UserHash,
			},
			RoomName: "roomid1",
		},
		chatRoomData{
			Accounts: []string{
				authUser.UserHash,
			},
			RoomName: "roomid2",
		},
	}

	beforeUser()
	defer afterUser()
	repo, _ := provider.UserRepository(authUser)
	app := repo.(*userRepository)

	beforeCount := 0
	ds.RDB().Model(&model.ChatRoom{}).Count(&beforeCount)
	res, err := app.getChatrooms(data)
	afterCount := 0
	ds.RDB().Model(&model.ChatRoom{}).Count(&afterCount)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(res) != 2 {
		t.Fatalf("count is weard %d", len(res))
	}

	if afterCount-beforeCount != 2 {
		t.Fatalf("dame %d, -> %d", beforeCount, afterCount)
	}
}

func Test条件が一緒であればチャットルームは複数作られない(t *testing.T) {
	data := []chatRoomData{
		chatRoomData{
			Accounts: []string{
				authUser.UserHash,
			},
			RoomName: "roomid",
		},
	}

	beforeUser()
	defer afterUser()
	repo, _ := provider.UserRepository(authUser)
	app := repo.(*userRepository)

	beforeCount := 0
	ds.RDB().Model(&model.ChatRoom{}).Count(&beforeCount)
	app.getChatrooms(data)
	app.getChatrooms(data)
	app.getChatrooms(data)
	app.getChatrooms(data)
	res, err := app.getChatrooms(data)
	afterCount := 0
	ds.RDB().Model(&model.ChatRoom{}).Count(&afterCount)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(res) != 1 {
		t.Fatalf("count is weird %d", len(res))
	}

	if afterCount-beforeCount != 1 {
		t.Fatalf("dame %d, -> %d", beforeCount, afterCount)
	}
}

func Testチャットを一気に読み込む際にはUserとChatRoomUserもプリロードされている(t *testing.T) {
	beforeUser()
	defer afterUser()
	repo, _ := provider.UserRepository(authUser)
	app := repo.(*userRepository)

	data := []chatRoomData{
		chatRoomData{
			Accounts: []string{
				authUser.UserHash,
			},
			RoomName: "roomid",
		},
	}
	app.createChatrooms(data)

	res := app.preloadRooms([]string{data[0].hashValue()})
	target := res[data[0].hashValue()]

	if len(target.Users) != 1 {
		buf := bytes.NewBufferString("")
		for _, u := range target.Users {
			fmt.Fprintf(buf, " %s", u.UserHash)
		}
		t.Fatalf("%d, %s", len(target.Users), buf.String())
	}

	if len(target.UserChatRooms) != 1 {
		t.Fatalf("failed to prelaod user chatroom")
	}
}

func Testメッセージのプリロード(t *testing.T) {
	beforeUser()
	defer afterUser()

	repo, _ := provider.UserRepository(authUser)
	app := repo.(*userRepository)
	room := createStubChatRoom(app)

	app.CreateMessage(CreateMessageRequest{
		Message: "test",
		Room:    room,
		User:    authUser,
	})

	res := app.preloadRooms([]string{room.RoomHash})

	target := res[room.RoomHash]

	if len(target.Messages) != 1 {
		t.Fatalf("%d", len(target.Messages))
	}
}

func Test条件が一緒であればチャットルームはたとえ同時リクエスであっても作られない(t *testing.T) {
	data := []chatRoomData{
		chatRoomData{
			Accounts: []string{
				authUser.UserHash,
			},
			RoomName: "roomid",
		},
		chatRoomData{
			Accounts: []string{
				authUser.UserHash,
			},
			RoomName: "roomid",
		},
	}

	beforeUser()
	defer afterUser()
	repo, _ := provider.UserRepository(authUser)
	app := repo.(*userRepository)

	beforeCount := 0
	ds.RDB().Model(&model.ChatRoom{}).Count(&beforeCount)
	app.getChatrooms(data)
	res, err := app.getChatrooms(data)
	afterCount := 0
	ds.RDB().Model(&model.ChatRoom{}).Count(&afterCount)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(res) != 1 {
		t.Fatalf("count is wierd %d", len(res))
	}

	if afterCount-beforeCount != 1 {
		t.Fatalf("dame %d, -> %d", beforeCount, afterCount)
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

func TestChatRoomからハッシュが作れる(t *testing.T) {
	data0 := chatRoomData{
		Accounts: []string{"t", "u"},
		RoomName: "hoge",
	}
	data1 := chatRoomData{
		Accounts: []string{"u", "t"},
		RoomName: "hoge",
	}

	hash0 := concatString(data0)
	hash1 := concatString(data1)

	if hash1 != "tuhoge" {
		t.Fail()
	}

	if hash1 != hash0 {
		t.Fatalf("hash should be same %s <=> %s", hash0, hash1)
	}
}

func TestChatRoomからハッシュが作れるてroomが違うと色々違う(t *testing.T) {
	data0 := chatRoomData{
		Accounts: []string{"t", "u"},
		RoomName: "piyo",
	}
	data1 := chatRoomData{
		Accounts: []string{"u", "t"},
		RoomName: "hoge",
	}

	hash0 := concatString(data0)
	hash1 := concatString(data1)

	if hash1 == hash0 {
		t.Fatalf("hash should be same %s <=> %s", hash0, hash1)
	}
}

func TestConvertToHash(t *testing.T) {
	if convertToHash("test") != "gCoQ4N_Csn3p8pqNz1SrRN2t2mhcj1ZOL2e9Pdn1srs" {
		t.Fatalf("faild to hash %s", convertToHash("test"))
	}
}

func TestPeekMessages(t *testing.T) {
	beforeUser()
	defer afterUser()
}

func Test未読カウントの取得(t *testing.T) {

}

func Testメッセージの作成(t *testing.T) {
	beforeUser()
	defer afterUser()

	repo, _ := provider.UserRepository(authUser)
	app := repo.(*userRepository)
	room := createStubChatRoom(app)
	rdb := app.ds.RDB()

	beforeCount := 0
	rdb.Model(&model.Message{}).Count(&beforeCount)

	err := app.CreateMessage(CreateMessageRequest{
		Message: "test",
		Room:    room,
		User:    authUser,
	})

	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	afterCount := 0
	rdb.Model(&model.Message{}).Count(&afterCount)

	if afterCount-beforeCount != 1 {
		t.Fatalf("bad count %d, %d", beforeCount, afterCount)
	}

	t.Run("ユーザーデータがアップデートされている", func(t *testing.T) {
		target := &model.UserChatRoom{}
		if rdb.Where(&model.UserChatRoom{
			ChatRoomID: room.ID,
			UserID:     authUser.ID,
		}).First(target).Error != nil {
			t.Fatalf("db error")
		}

		if target.UpdatedAt.Unix() < time.Now().Unix()-100 {
			t.Fail()
		}
	})
}

func createStubChatRoom(app *userRepository) model.ChatRoom {
	rooms, _ := app.getChatrooms([]chatRoomData{
		chatRoomData{
			Accounts: []string{authUser.UserHash},
			RoomName: "test",
		},
	})
	return *rooms[0]
}
