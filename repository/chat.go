package repository

import (
	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/model"
)

type chatRepository struct {
	ds datasource.DataSourceDescriptor
}

func (r chatRepository) PeekMessages([]*model.ChatRoom) map[string]model.Message {
	panic("not implemented")
}

func (r chatRepository) GetUnreadCount([]*model.ChatRoom) map[string]*Unread {
	panic("not implemented")
}
