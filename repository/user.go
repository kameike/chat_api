package repository

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/kameike/chat_api/apierror"
	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/model"
)

type userRepository struct {
	user model.User
	ds   datasource.DataSourceDescriptor
}

type chatRoomData struct {
	Users        []string `json:"users"`
	RoomName     string   `json:"roomId"`
	RoomMemoText string   `json:"roomName"`
}

func (u *userRepository) GetChatRooms(data ChatRoomsInfoDescriable) ([]*model.ChatRoom, apierror.ChatAPIError) {
	hashes := data.RoomHashes()
	roomsInfo := make([]chatRoomData, len(hashes), len(hashes))

	for i, chank := range data.RoomHashes() {
		err := json.Unmarshal([]byte(chank), &roomsInfo[i])
		if err != nil {
			return nil, apierror.GeneralError(err)
		}
	}
	return u.getChatrooms(roomsInfo)
}

func (u *userRepository) getChatrooms(data []chatRoomData) ([]*model.ChatRoom, apierror.ChatAPIError) {
	result := u.findChatrooms(data)
	currentChatrooms := result.found

	if len(result.notFound) != 0 {
		err := u.createChatrooms(result.notFound)
		if err != nil {
			return currentChatrooms, apierror.GeneralError(err)
		}
		result := u.findChatrooms(result.notFound)

		if len(result.notFound) != 0 {
			panic("cant find chatrooms even create succeeded")
		}
		currentChatrooms = append(currentChatrooms, result.found...)
	}
	return currentChatrooms, nil
}

type findChatRoomInfo struct {
	found    []*model.ChatRoom
	notFound []chatRoomData
}

func (u *userRepository) createChatrooms(data []chatRoomData) apierror.ChatAPIError {
	userMaps := u.preloadUser(extractUserHashes(data))
	errors := []apierror.ChatAPIError{}
	rooms := []*model.ChatRoom{}

NEXT_CHAT_ROOM:
	for _, d := range data {
		users := make([]model.User, len(d.Users), len(d.Users))

		if len(users) == 0 {
			println("cant create empty room")
			continue NEXT_CHAT_ROOM
		}

		for i, u := range d.Users {
			t := userMaps[u]
			if t == nil {
				errors = append(errors, apierror.FailToCreateChatRooom(d.description()))
				println("faild to create due to user")
				continue NEXT_CHAT_ROOM
			}
			users[i] = *t
		}

		room := &model.ChatRoom{
			RoomHash: d.hashValue(),
			Name:     d.RoomName,
			Users:    users,
		}

		rooms = append(rooms, room)
	}

	if len(rooms) == 0 {
		return apierror.NestedError(errors)
	}

	db := u.ds.RDB()
	for _, r := range rooms {
		e := db.Save(r).Error
		if e != nil {
			println("faild to save")
			errors = append(errors, apierror.GeneralError(e))
		}
	}

	if len(errors) != 0 {
		return apierror.NestedError(errors)
	}

	return nil
}
func extractUserHashes(data []chatRoomData) []string {
	result := make([]string, 0, len(data))

	for _, d := range data {
		result = append(result, d.Users...)
	}

	return result
}

func (u *userRepository) findChatrooms(data []chatRoomData) findChatRoomInfo {
	var roomHashes = make([]string, len(data), len(data))
	roomInfoMap := map[string]*chatRoomData{}

	for i, d := range data {
		roomHashes[i] = d.hashValue()
		data := d
		roomInfoMap[d.hashValue()] = &data
	}

	rooms := u.preloadRooms(roomHashes)
	notfoundRooms := make([]chatRoomData, 0, len(data))
	foundRooms := make([]*model.ChatRoom, 0, len(data))

	for _, r := range roomHashes {
		if rooms[r] == nil {
			notfoundRooms = append(notfoundRooms, *roomInfoMap[r])
		} else {
			foundRooms = append(foundRooms, rooms[r])
		}
	}

	return findChatRoomInfo{
		found:    foundRooms,
		notFound: notfoundRooms,
	}
}

func (u *userRepository) preloadRooms(hashes []string) map[string]*model.ChatRoom {
	db := u.ds.RDB()
	target := map[string]*model.ChatRoom{}
	rooms := make([]*model.ChatRoom, 0, len(hashes))

	db.Preload("UserChatRooms").Preload("Users").Preload("Messages").Where("room_hash in (?)", hashes).Find(&rooms)

	for _, r := range rooms {
		target[r.RoomHash] = r
	}

	return target
}

func (u *userRepository) preloadUser(hashes []string) map[string]*model.User {
	l := len(hashes)
	target := make(map[string]*model.User)
	users := make([]model.User, l, l)
	db := u.ds.RDB()

	db.Where("user_hash in (?)", hashes).Find(&users)

	for _, u := range users {
		target[u.UserHash] = &u
	}

	return target
}

func (d chatRoomData) description() string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "faild to create %s", d.RoomMemoText)
	return buf.String()
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

	target += data.RoomName
	return target
}

func (u *userRepository) UpdateUser(data UserUpdateInfoDescriable) (*model.User, apierror.ChatAPIError) {
	rdb := u.ds.RDB()

	user := u.user

	if data.ImageURL() != nil {
		user.Url = *data.ImageURL()
	}

	if data.Name() != nil {
		user.Name = *data.Name()
	}

	err := rdb.Save(&user).Error

	if err != nil {
		return nil, apierror.GeneralError(err)
	}

	return &user, nil
}

type CreateMessageRequest struct {
	Message string
	User    model.User
	Room    model.ChatRoom
}

type GetMessageRequest struct {
}

func (u *userRepository) CreateMessage(req CreateMessageRequest) apierror.ChatAPIError {
	db := u.ds.RDB()
	message := &model.Message{
		UserID: req.User.ID,
		RoomID: req.Room.ID,
		Text:   req.Message,
	}

	err := db.Save(message).Error

	if err != nil {
		return apierror.GeneralError(err)
	}

	rel := &model.UserChatRoom{}

	err = db.Where(&model.UserChatRoom{
		UserID:     req.User.ID,
		ChatRoomID: req.Room.ID,
	}).First(rel).Error

	if err != nil {
		return apierror.GeneralError(err)
	}

	db.Model(&rel).Update("UpdatedAt", time.Now())

	return nil
}

func (u *userRepository) GetMessages(GetMessageRequest) ([]*model.Message, apierror.ChatAPIError) {
	return nil, nil
}

func (u *userRepository) GetChatRoom(req GetChatRoomRequest) (*model.ChatRoom, apierror.ChatAPIError) {
	db := u.ds.RDB()
	result := model.ChatRoom{}

	err := db.Preload("Users").Preload("Messages").Model(&model.ChatRoom{}).Where("room_hash = ?", req.Hash).Find(&result).Error

	if err != nil {
		return nil, apierror.GeneralError(err)
	}

	isContain := false

	for _, us := range result.Users {
		if us.ID == u.user.ID {
			isContain = true
		}
	}

	if !isContain {
		return nil, apierror.ErrorNoPermission()
	}

	return &result, nil
}
