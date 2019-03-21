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
	t.Skipf("skipping")
	beforeUser()
	defer afterUser()
	u, _ := provider.UserRepository(authUser)

	var roomSign []string
	roomSign = append(roomSign, `{
			"users": ["hogehohge", "fugafuga"],
			"roomId": "hoge",
			"roomName": "fuga"
		}
	`)

	testData := testChatRoomCreateInfo{
		data: roomSign,
	}

	result, err := u.GetChatRooms(testData)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if result != nil {
		t.Fatalf("chat room has not been created")
	}
}

func Testチャットルームのタイプを渡すといい感じになる(t *testing.T) {
	data := []chatRoomData{
		chatRoomData{
			Users: []string{
				authUser.UserHash,
			},
			RoomId:   "roomid",
			RoomName: "hoge",
		},
	}

	beforeUser()
	defer afterUser()
	repo, _ := provider.UserRepository(authUser)
	app := repo.(*userRepository)
	res, err := app.getChatrooms(data)

	if err != nil {
		t.Fail()
	}

	if len(res) != 1 {
		t.Fail()
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
		Users:    []string{"t", "u"},
		RoomId:   "hoge",
		RoomName: "piyo",
	}
	data1 := chatRoomData{
		Users:    []string{"u", "t"},
		RoomId:   "hoge",
		RoomName: "piyo",
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

func TestConvertToHash(t *testing.T) {
	if convertToHash("test") != "gCoQ4N_Csn3p8pqNz1SrRN2t2mhcj1ZOL2e9Pdn1srs" {
		t.Fatalf("faild to hash %s", convertToHash("test"))
	}
}
