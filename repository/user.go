package repository

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"sort"

	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/error"
	"github.com/kameike/chat_api/model"
)

type userRepository struct {
	user model.User
	ds   datasource.DataSourceDescriptor
}

type chatRoomData struct {
	Users    []string `json:"users"`
	RoomId   string   `json:"roomId"`
	RoomName string   `json:"roomName"`
}

func (u *userRepository) GetChatRooms(data ChatRoomsInfoDescriable) (*model.ChatRoom, error.ChatAPIError) {
	room := model.ChatRoom{}
	result := room
	hashes := data.RoomHashes()

	rooms := make([]chatRoomData, len(hashes), len(hashes))

	for i, chank := range data.RoomHashes() {
		err := json.Unmarshal([]byte(chank), &rooms[i])
		if err != nil {
			return nil, error.GeneralError(err)
		}
	}

	return &result, nil
}

func (d chatRoomData) hashValue() string {
	res := concatString(d)
	return convertToHash(res)
}

func convertToHash(seed string) string {
	solt := "n4bQgYhMfWWaL-qgxVrQFaO_TxsrC4Is0V1sFbDwCgg"
	hasher := sha256.New()
	hasher.Write([]byte(seed + solt))
	return base64.RawURLEncoding.EncodeToString(hasher.Sum(nil))
}

func concatString(data chatRoomData) string {
	target := ""

	users := data.Users
	sort.Slice(users, func(i, j int) bool { return users[i] < users[j] })

	for _, u := range users {
		target += u
	}

	target += data.RoomId
	return target
}

func (u *userRepository) UpdateUser(data UserUpdateInfoDescriable) (*model.User, error.ChatAPIError) {
	rdb := u.ds.RDB()

	user := u.user

	if data.ImageURL() != nil {
		user.Url = *data.ImageURL()
	}

	if data.ImageURL() != nil {
		user.Name = *data.Name()
	}

	err := rdb.Save(&user).Error

	if err != nil {
		return nil, error.GeneralError(err)
	}

	return &user, nil
}
