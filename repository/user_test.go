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
	roomSign = append(roomSign, fmt.Sprintf(`{
			"users": ["%s"],
			"roomId": "hoge",
			"roomName": "fuga"
		}
	`, authUser.UserHash))

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

func Testチャットルームの型によってチャットを作ることができる(t *testing.T) {
	data := []chatRoomData{
		chatRoomData{
			Users: []string{
				authUser.UserHash,
			},
			RoomName:     "roomid",
			RoomMemoText: "hoge",
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

	if afterRelationCount-beforeRelationCount != 1 {
		t.Fatalf("failed to make relation")
	}

	t.Run("ハッシュからチャットルームを見つけることができる", func(t *testing.T) {
		ds.RDB().LogMode(true)
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

		ds.RDB().LogMode(false)
		if res == nil {
			t.Fatalf("result must not be nil")
		}

		if len(res.Users) != 1 {
			t.Fatalf("faild to preload user")
		}

		if len(res.Messages) != 1 {
			t.Fatalf("faild to preload message %d", len(res.Messages))
		}

		for _, u := range res.Users {
			if u.ID != authUser.ID {
				t.Fatalf("invald user")
			}
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
			Users: []string{
				authUser.UserHash,
			},
			RoomName:     "roomid1",
			RoomMemoText: "pppp",
		},
		chatRoomData{
			Users: []string{
				authUser.UserHash,
			},
			RoomName:     "roomid2",
			RoomMemoText: "hohoh",
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
			Users: []string{
				authUser.UserHash,
			},
			RoomName:     "roomid",
			RoomMemoText: "hoge",
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
		t.Fatalf("count is weard %d", len(res))
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
			Users: []string{
				authUser.UserHash,
			},
			RoomName:     "roomid",
			RoomMemoText: "hoge",
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
			Users: []string{
				authUser.UserHash,
			},
			RoomName:     "roomid",
			RoomMemoText: "hoge",
		},
		chatRoomData{
			Users: []string{
				authUser.UserHash,
			},
			RoomName:     "roomid",
			RoomMemoText: "hoge",
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

	if len(res) != 2 {
		t.Fatalf("count is weard %d", len(res))
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

type testChatRoomCreateInfo struct {
	data []string
}

func (t testChatRoomCreateInfo) RoomHashes() []string {
	return t.data
}

func TestChatRoomからハッシュが作れる(t *testing.T) {
	data0 := chatRoomData{
		Users:        []string{"t", "u"},
		RoomName:     "hoge",
		RoomMemoText: "piyo",
	}
	data1 := chatRoomData{
		Users:        []string{"u", "t"},
		RoomName:     "hoge",
		RoomMemoText: "piyo",
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
		Users:        []string{"t", "u"},
		RoomName:     "piyo",
		RoomMemoText: "piyo",
	}
	data1 := chatRoomData{
		Users:        []string{"u", "t"},
		RoomName:     "hoge",
		RoomMemoText: "piyo",
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
			Users:        []string{authUser.UserHash},
			RoomName:     "test",
			RoomMemoText: "hoge",
		},
	})
	return *rooms[0]
}
