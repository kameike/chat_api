package repository

import (
	"math/rand"

	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/error"
	"github.com/kameike/chat_api/model"
)

type authRepository struct {
	d datasource.DataSourceDescriptor
}

func (r *authRepository) FindOrCreateUser(token string, hash string) (*model.User, *model.AccessToken, error.ChatAPIError) {
	rdb := r.d.RDB()
	err := rdb.Create(&model.User{
		AuthToken: token,
		UserHash:  hash,
	}).Error

	if err != nil {
		return nil, nil, error.ErrorLoginAuthFail(err)
	}

	updatedUser := &model.User{}

	err = rdb.Where("auth_token = ?", token).First(updatedUser).Error

	if err != nil {
		return nil, nil, error.ErrorLoginAuthFail(err)
	}

	accestToken := &model.AccessToken{
		AccessToken: generateRandomToken(),
		UserID:      updatedUser.ID,
	}

	err = rdb.Create(accestToken).Error

	if err != nil {
		return nil, nil, error.ErrorLoginAuthFail(err)
	}

	return updatedUser, accestToken, nil
}

func (r *authRepository) FindUser(token string) (*model.User, error.ChatAPIError) {
	db := r.d.RDB()

	tokenResult := &model.AccessToken{}
	user := &model.User{}
	tokenResult.AccessToken = token

	err := db.Where("access_token = ?", token).First(tokenResult).Error
	if err != nil {
		return nil, error.GeneralError(err)
	}
	db.Model(tokenResult).Related(user)

	return user, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func generateRandomToken() string {
	return randStringBytes(128)
}
