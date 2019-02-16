package repository

import (
	"testing"

	"github.com/kameike/chat_api/datasource"
	. "github.com/kameike/chat_api/model"
)

var repo *userRepotory

func prepareDB() {
	ds := datasource.PrepareInmemoryDatasource()
	repo = &userRepotory{
		datasource: ds,
	}
	ds.MigrateIfNeed()
}

func Testユーザーが作れる(t *testing.T) {
	prepareDB()

	err := repo.createUser(&User{
		AuthToken: "test",
		Name:      "kameike",
		PushToken: "token",
		UserHash:  "hash",
	})

	if err != nil {
		t.Fail()
	}

	count := 0
	repo.datasource.RDB().Model(&User{}).Count(&count)

	if count != 1 {
		t.Fail()
	}
}

func Testユーザーを作った際にアクセストークンも作られる(t *testing.T) {
	prepareDB()

	err := repo.createUser(&User{
		AuthToken: "test33",
		Name:      "kameike",
		PushToken: "token",
		UserHash:  "hashheo",
	})

	if err != nil {
		t.Fail()
	}

	count := 0
	repo.datasource.RDB().Model(&AccessToken{}).Count(&count)

	if count != 1 {
		t.Fail()
	}
}

func Test重複したHashを持ったユーザーは作れない(t *testing.T) {
	prepareDB()

	repo.createUser(&User{
		AuthToken: "test",
		Name:      "kameike",
		PushToken: "token",
		UserHash:  "hash",
	})

	err := repo.createUser(&User{
		AuthToken: "test2",
		Name:      "ttt",
		PushToken: "ooo",
		UserHash:  "hash",
	})

	if err == nil {
		t.Fail()
	}

	count := 0
	repo.datasource.RDB().Model(&User{}).Count(&count)

	if count != 1 {
		t.Fail()
	}
}

func Test重複したAuthTokenを持ったユーザーは作れない(t *testing.T) {
	prepareDB()

	repo.createUser(&User{
		AuthToken: "test",
		Name:      "kameike",
		PushToken: "token",
		UserHash:  "hash",
	})

	err := repo.createUser(&User{
		AuthToken: "test",
		Name:      "ttt",
		PushToken: "ooo",
		UserHash:  "ppp",
	})

	if err == nil {
		t.Fail()
	}

	count := 0
	repo.datasource.RDB().Model(&User{}).Count(&count)

	if count != 1 {
		t.Fail()
	}
}
