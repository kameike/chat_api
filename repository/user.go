package repository

import (
	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/error"
	"github.com/kameike/chat_api/model"
)

type userRepository struct {
	user model.User
	ds   datasource.DataSourceDescriptor
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
