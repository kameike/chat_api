package repository

import "github.com/kameike/chat_api/datasource"
import "github.com/kameike/chat_api/model"
import "github.com/kameike/chat_api/error"

// import "github.com/jinzhu/gorm"

type userRepotory struct {
	datasource datasource.DataSourceDescriptor
}

func (r *userRepotory) createUser(user *model.User, token string) error.ChatAPIError {
	rdb := r.datasource.RDB()
	err := rdb.Create(user).Error

	if err != nil {
		return error.ErrorLoginAuthFail(err)
	}

	updatedUser := &model.User{}

	err = rdb.Where("auth_token = ?", user.AuthToken).First(updatedUser).Error

	if err != nil {
		return error.ErrorLoginAuthFail(err)
	}

	accestToken := &model.AccessToken{
		AccessToken: token,
		UserID:      updatedUser.ID,
	}

	err = rdb.Create(accestToken).Error

	if err != nil {
		return error.ErrorLoginAuthFail(err)
	}

	return nil
}

func (r *userRepotory) updateToken(user *model.User, newToken string) *model.AccessToken {
	db := r.datasource.RDB()
	token := &model.AccessToken{}
	db.Model(user).Related(token)
	token.AccessToken = newToken
	db.Save(token)

	return token
}

func (r *userRepotory) updateName(user *model.User, name string) {
	user.Name = name
	db := r.datasource.RDB()
	db.Save(user)
}

func (r *userRepotory) findUser(token string) (*model.User, error.ChatAPIError) {
	db := r.datasource.RDB()

	tokenResult := &model.AccessToken{}
	user := &model.User{}

	db.First(tokenResult)
	db.Model(tokenResult).Related(user)

	return user, nil
}

func generateRandomToken() string {
	return "hoghogehogeo"
}
