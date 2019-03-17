package repository

import (
	"testing"

	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/model"
)

func before() {
	d = datasource.PrepareDatasource()
	d.Begin()
	r = &authRepository{d}
	token = generateRandomToken()
	hash = generateRandomToken()
}

func after() {
	d.Rollback()
}

var r *authRepository
var d datasource.DataSourceDescriptor
var hash string
var token string

func Test_authRepository_FindOrCreateUser_普通にユーザーが作れる(t *testing.T) {
	before()
	defer after()

	beforeCount := 0
	d.RDB().Model(&model.User{}).Count(&beforeCount)

	user, tokens, _ := r.FindOrCreateUser(token, hash)

	afterCount := 0
	d.RDB().Model(&model.User{}).Count(&afterCount)

	if afterCount-beforeCount != 1 {
		t.Fail()
	}

	if user.AuthToken != token {
		t.Fail()
	}

	if user.UserHash != hash {
		t.Fail()
	}

	if tokens.UserID != user.ID {
		t.Fail()
	}
}
func Test_authRepository_FindOrCreateUser_ユーザーが重複して作れない(t *testing.T) {
	before()
	defer after()

	beforeCount := 0
	d.RDB().Model(&model.User{}).Count(&beforeCount)

	r.FindOrCreateUser(token, hash)

	_, _, err := r.FindOrCreateUser("anther token", hash)
	if err == nil {
		t.Fail()
	}
	_, _, err = r.FindOrCreateUser(token, "another hash")
	if err == nil {
		t.Fail()
	}
}

func Test_authRepository_FindUserユーザーが見つかる(t *testing.T) {
	before()
	defer after()

	user, token, _ := r.FindOrCreateUser(token, hash)
	user1, _ := r.FindUser(token.AccessToken)

	if user.ID != user1.ID {
		t.Fail()
	}
}

func Test_authRepository_FindUser存在しないユーザーは見つからない(t *testing.T) {
	before()
	defer after()

	user1, _ := r.FindUser(generateRandomToken())

	if user1 != nil {
		t.Fail()
	}
}
