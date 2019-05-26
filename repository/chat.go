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
		return apierror.Error(apierror.POST_MESSAGE_FAILD, err)
	}

	rel := &model.UserChatRoom{}

	err = db.Where(&model.UserChatRoom{
		UserID:     req.User.ID,
		ChatRoomID: req.Room.ID,
	}).First(rel).Error

	if err != nil {
		return apierror.Error(apierror.RELOAD_AFTER_POST_FALID, err)
	}

	db.Model(&rel).Update("UpdatedAt", time.Now())

	return nil
}

func (r *chatRepository) GetMessageAndReadStatus() (*MessageAndReadState, apierror.ChatAPIError) {
	db := r.ds.RDB()
	messages := []*model.Message{}

	err := db.Where("room_id = ?", r.room.ID).Find(&messages).Error

	if err != nil {
		return nil, apierror.Error(apierror.GET_MESSAGE_FAIL, err)
	}

	return &MessageAndReadState{
		Messages: messages,
	}, nil
}
