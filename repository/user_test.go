package repository

import (
	"testing"

	"github.com/kameike/chat_api/datasource"
	. "github.com/kameike/chat_api/model"
)

var repo *userRepotory
var accessToken = "hogehoge"

func prepareDB() {
	ds := datasource.PrepareDatasource()
	ds.RDB().LogMode(true)
	println("==== new case ====")
	repo = &userRepotory{
		datasource: ds,
	}
	ds.Begin()
}

func cleanDB() {
	repo.datasource.Rollback()
}

func prepareDBWithUsers() *User {
	prepareDB()
	db := repo.datasource.RDB()

	token := "token"

	user := &User{
		AuthToken: token,
		UserHash:  "test",
	}

	repo.createUser(user, accessToken)
	db.First(user)

	return user
}

func Testユーザーが作れる(t *testing.T) {
	prepareDB()
	defer cleanDB()

	err := repo.createUser(&User{
		AuthToken: "test",
		Name:      "kameike",
		PushToken: "token",
		UserHash:  "hash",
	}, accessToken)

	if err != nil {
		t.Fail()
	}

	count := 0
	repo.datasource.RDB().Model(&User{}).Count(&count)

	if count != 1 {
		t.Fail()
	}
}

func Test名前がなくてもユーザーが作れる(t *testing.T) {
	prepareDB()
	defer cleanDB()

	err := repo.createUser(&User{
		AuthToken: "test2",
		PushToken: "token2",
		UserHash:  "hash2",
	}, accessToken)

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
	defer cleanDB()

	err := repo.createUser(&User{
		AuthToken: "test33",
		Name:      "kameike",
		PushToken: "token",
		UserHash:  "hashheo",
	}, accessToken)

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
	defer cleanDB()

	repo.createUser(&User{
		AuthToken: "test",
		Name:      "kameike",
		PushToken: "token",
		UserHash:  "hash",
	}, accessToken)

	err := repo.createUser(&User{
		AuthToken: "test2",
		Name:      "ttt",
		PushToken: "ooo",
		UserHash:  "hash",
	}, "newAccessToken")

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
	defer cleanDB()

	repo.createUser(&User{
		AuthToken: "test",
		Name:      "kameike",
		PushToken: "token",
		UserHash:  "hash",
	}, accessToken)

	err := repo.createUser(&User{
		AuthToken: "test",
		Name:      "ttt",
		PushToken: "ooo",
		UserHash:  "ppp",
	}, "newAccessToken")

	if err == nil {
		t.Fail()
	}

	count := 0
	repo.datasource.RDB().Model(&User{}).Count(&count)

	if count != 1 {
		t.Fail()
	}
}

func Testアクセストークンで探せる(t *testing.T) {
	prepareDB()
	defer cleanDB()

	accessToken := "hoghogehogeo"
	token := "token"

	repo.createUser(&User{
		AuthToken: token,
		UserHash:  "test",
		PushToken: "test",
	}, accessToken)

	user, err := repo.findUser(accessToken)

	if err != nil {
		t.Fail()
	}

	if user == nil {
		t.Fatalf("user not founded")
	}

	if user.AuthToken != token {
		t.Fatalf("invalid token %s", user.AuthToken)
	}
}

func Testユーザー名が更新される(t *testing.T) {
	user := prepareDBWithUsers()
	defer cleanDB()

	db := repo.datasource.RDB()

	repo.updateName(user, "newName")
	db.First(user)

	if user.Name != "newName" {
		t.Fatalf("user name is %s", user.Name)
	}
}

func Testトークンが更新される(t *testing.T) {
	user := prepareDBWithUsers()
	defer cleanDB()
	db := repo.datasource.RDB()

	oldToken := &AccessToken{}

	db.Model(user).Related(oldToken)

	repo.updateToken(user, "newToken")

	newToken := &AccessToken{}
	db.Model(user).Related(newToken)

	if newToken.AccessToken == oldToken.AccessToken {
		t.Fail()
	}
}
