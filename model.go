package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	AuthToken string `gorm:"type:varchar(255);unique;index"`
	UserHash  string `gorm:"type:varchar(255);unique;index"`
	Name      string `gorm:"type:varchar(255)"`
	PushToken string `gorm:"type:varchar(255)"`
}

type AccessToken struct {
	UserID      int
	AccessToken string `gorm:"type:varchar(255);unique;index"`
}

type UserChatRoom struct {
	gorm.Model
	UserID     int
	ChatRoomID int
	LastReadAt int
}

type ChatRoom struct {
	gorm.Model
	Users    []User
	RoomHash string
	Name     string
}

type Message struct {
	gorm.Model
	Text      string
	UserID    int `gorm:"index"`
	RoomID    int `gorm:"index"`
	TimeStamp int64
}

func tet() {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		println(err)
		panic("dame")
	}

	migrate(db)
}

type ChatRoomRedisModel struct {
	RoomHash         string
	LastReadAt       map[string]int64
	UnreadCountCache map[string]int64
	Message          MessageRedisModel
}

type MessageRedisModel struct {
	UserID int
	Text   string
}

type UserRepositable interface {
	createUser(user User)
	updateToken(user User)
	updateAccount(user User)
}

type ChatRepostitable interface {
	createChatRoom()
	getChatRoom()
	postMessage()
	getMessages()
	getUnreads()
}

func migrate(db *gorm.DB) {
	db.CreateTable(&User{})
	db.CreateTable(&AccessToken{})
}
