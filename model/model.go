package model

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	AuthToken string `gorm:"type:varchar(255);unique;index"`
	UserHash  string `gorm:"type:varchar(255);unique;index"`
	Name      string `gorm:"type:varchar(255)"`
	Url       string `gorm:"type:varchar(255)"`
	PushToken string `gorm:"type:varchar(255)"`
}

type AccessToken struct {
	UserID      uint   `gorm:"index"`
	AccessToken string `gorm:"type:varchar(255);unique;index"`
	User        User   `gorm:"foreignkey:UserID"`
}

type UserChatRoom struct {
	gorm.Model
	UserID     uint `gorm:"index"`
	ChatRoomID uint `gorm:"index"`
	LastReadAt uint
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

func migrate(db *gorm.DB) {
	// 	db.CreateTable(&User{})
	// 	db.CreateTable(&AccessToken{})
	// 	db.CreateTable(&UserChatRoom{})
	// 	db.CreateTable(&ChatRoom{})
	// 	db.CreateTable(&Message{})
	db.Model(&User{}).ModifyColumn("test", "text")
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
