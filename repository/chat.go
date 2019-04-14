package repository

import (
	"time"

	"github.com/kameike/chat_api/apierror"
	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/model"
)

type chatRepository struct {
	ds   datasource.DataSourceDescriptor
	room model.ChatRoom
	user model.User
}

func (c *chatRepository) CreateMessage(t string) apierror.ChatAPIError {
	req := CreateMessageRequest{
		Message: t,
		Room:    c.room,
		User:    c.user,
	}

	db := c.ds.RDB()
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
