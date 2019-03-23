package model

import (
	"time"

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
	ID         uint `gorm:"index"`
	UserID     uint `gorm:"index"`
	ChatRoomID uint `gorm:"index"`
	LastReadAt uint
	UpdatedAt  time.Time
}

type ChatRoom struct {
	gorm.Model
	Users         []User    `gorm:"many2many:user_chat_rooms;"`
	Messages      []Message `gorm:"foreignkey:RoomID"`
	UserChatRooms []UserChatRoom
	RoomHash      string
	Name          string
}

type Message struct {
	gorm.Model
	Text   string
	UserID uint `gorm:"index"`
	RoomID uint `gorm:"index"`
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
