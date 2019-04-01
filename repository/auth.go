package repository

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/kameike/chat_api/apierror"
	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/model"
)

type authRepository struct {
	d datasource.DataSourceDescriptor
}

func (r *authRepository) createUser(user model.User) (*model.User, apierror.ChatAPIError) {
	rdb := r.d.RDB()

	err := rdb.Create(&user).Error

	if err != nil {
		return nil, apierror.ErrorLoginAuthFail(err)
	}

	updatedUser := &model.User{}
	err = rdb.Where("user_hash = ?", user.UserHash).First(updatedUser).Error

	if err != nil {
		return nil, apierror.ErrorLoginAuthFail(err)
	}
	return updatedUser, nil
}

func (r *authRepository) FindOrCreateUser(token string, hash string) (*model.User, *model.AccessToken, apierror.ChatAPIError) {
	rdb := r.d.RDB()

	updatedUser := &model.User{}
	err := rdb.Where("user_hash = ?", hash).First(updatedUser).Error

	if gorm.IsRecordNotFoundError(err) {
		updatedUser, err = r.createUser(model.User{
			UserHash:  hash,
			AuthToken: token,
		})
	}

	if err != nil {
		return nil, nil, apierror.ErrorLoginAuthFail(err)
	}

	accessToken, err := r.createOrUpdateAccessToken(*updatedUser)

	if err != nil {
		return nil, nil, apierror.ErrorLoginAuthFail(err)
	}

	return updatedUser, accessToken, nil
}

func (r *authRepository) createOrUpdateAccessToken(u model.User) (*model.AccessToken, apierror.ChatAPIError) {
	rdb := r.d.RDB()

	accestToken := &model.AccessToken{
		UserID: u.ID,
	}

	err := rdb.Where("user_id = ?", u.ID).First(&accestToken).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, apierror.ErrorLoginAuthFail(err)
	}

	accestToken.AccessToken = generateRandomToken()
	err = rdb.Save(accestToken).Error

	return accestToken, nil
}

func (r *authRepository) FindUser(token string) (*model.User, apierror.ChatAPIError) {
	db := r.d.RDB()

	tokenResult := &model.AccessToken{}
	user := &model.User{}
	tokenResult.AccessToken = token

	err := db.Where("access_token = ?", token).First(tokenResult).Error
	if err != nil {
		return nil, apierror.GeneralError(err)
	}
	db.Model(tokenResult).Related(user)

	return user, nil
}

const letterBytes = "abcdefghijk1234567890lmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func generateRandomToken() string {
	rand.Seed(time.Now().UnixNano())
	res := randStringBytes(128)
	return res
}
