package repository

import (
	"testing"

	"github.com/kameike/chat_api/datasource"
	"github.com/kameike/chat_api/model"
)

func generalBefore() {
	ds = datasource.PrepareDatasource()
	// ds.RDB().LogMode(true)
	println("===start case===")
	ds.Begin()

	provider = &applicationRepositoryProvidable{
		datasource: ds,
	}
}

func generalAfter() {
	ds.Rollback()
}

func authBefore() {
	generalBefore()
	token = generateRandomToken()
	hash = generateRandomToken()
}

func authAfter() {
	generalAfter()
}

var provider ReposotryProvidable
var ds datasource.DataSourceDescriptor
var hash string
var token string

func Test_authRepository_FindOrCreateUser_普通にユーザーが作れる(t *testing.T) {
	authBefore()
	defer authAfter()
	r := provider.AuthRepository()

	beforeCount := 0
	ds.RDB().Model(&model.User{}).Count(&beforeCount)

	user, tokens, err := r.FindOrCreateUser(token, hash)

	afterCount := 0
	ds.RDB().Model(&model.User{}).Count(&afterCount)

	if err != nil {
		t.Fail()
	}

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
	authBefore()
	defer authAfter()
	r := provider.AuthRepository()

	beforeCount := 0
	ds.RDB().Model(&model.User{}).Count(&beforeCount)

	r.FindOrCreateUser(token, hash)

	_, _, err := r.FindOrCreateUser(generateRandomToken(), hash)
	if err != nil {
		t.Fail()
	}
	_, _, err = r.FindOrCreateUser(token, generateRandomToken())
	if err == nil {
		t.Fail()
	}
}

func Test_authRepository_FindOrCreateUser_作ったユーザーをfindできる(t *testing.T) {
	authBefore()
	defer authAfter()
	r := provider.AuthRepository()

	beforeCount := 0
	ds.RDB().Model(&model.User{}).Count(&beforeCount)

	user, _, _ := r.FindOrCreateUser(token, hash)
	user1, _, err := r.FindOrCreateUser(token, hash)

	if err != nil {
		t.Fail()
	}

	if user.ID != user1.ID {
		t.Fail()
	}
}

func Test_authRepository_FindUserユーザーが見つかる(t *testing.T) {
	authBefore()
	defer authAfter()
	r := provider.AuthRepository()

	user, token, _ := r.FindOrCreateUser(token, hash)
	user1, _ := r.FindUser(token.AccessToken)

	if user.ID != user1.ID {
		t.Fail()
	}
}

func Test_authRepository_FindUser存在しないユーザーは見つからない(t *testing.T) {
	authBefore()
	defer authAfter()
	r := provider.AuthRepository()

	user1, _ := r.FindUser(generateRandomToken())

	if user1 != nil {
		t.Fail()
	}
}

func Test_authRepository_FindOrCreateUser_再取得時にtokenが更新される(t *testing.T) {
	authBefore()
	defer authAfter()
	r := provider.AuthRepository()

	_, accessToken1, _ := r.FindOrCreateUser(token, hash)
	_, accessToken2, _ := r.FindOrCreateUser(token, hash)

	if accessToken1.AccessToken == accessToken2.AccessToken {
		t.Fail()
	}
}
