package main

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/kameike/chat_api/swggen/apimodel"
	apiclient "github.com/kameike/chat_api/swggen/client"
	"github.com/kameike/chat_api/swggen/client/account"
)

var count = 100

func main() {
	rand.Seed(time.Now().UnixNano())
	createAccounts()

	wg := sync.WaitGroup{}
	for _, u := range users {
		wg.Add(1)
		time.Sleep(1000 * 1000 * 10)
		uc := u
		go func() {
			updateProfiles(uc)
			wg.Done()
		}()
	}
	wg.Wait()

	for i := 0; i < 1000; i++ {
		r := randomRoomInfo()
		println(r.user1.user.ID, r.user2.user.ID, r.name)
	}
}

var counter = 0
var users = make([]basicAccount, count, count)

type basicAccount struct {
	accessToken string
	user        apimodel.User
}

type roomInfo struct {
	user1 basicAccount
	user2 basicAccount
	name  string
	hash  string
}

func randomRoomInfo() roomInfo {
	u1 := users[rand.Intn(len(users))]
	u2 := users[rand.Intn(len(users))]
	name := randStringBytes(10)

	return roomInfo{
		user1: u1,
		user2: u2,
		name:  name,
	}
}

func authFor(a basicAccount) runtime.ClientAuthInfoWriter {
	return httptransport.APIKeyAuth("x_chat_access_token", "header", a.accessToken)
}

func updateProfiles(a basicAccount) {
	newName := randStringBytes(30)

	res, err := client.Account.PostProfile(&account.PostProfileParams{
		Context: cxt,
		Name:    &newName,
	}, authFor(a))

	if err != nil {
		println(err.Error())
	} else {
		println(newName, "->", res.Payload.Name)
	}
}

func createAccounts() {
	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		time.Sleep(1000 * 1000 * 1)
		c := i
		go func() {
			testCreateAcoount(c)
			wg.Done()
		}()
	}
	wg.Wait()
}

var cxt = context.Background()
var host = "localhost:1323"
var transport = httptransport.New(host, "", nil)
var client = apiclient.New(transport, strfmt.Default)

func testCreateAcoount(index int) {
	c := counter
	counter++

	before := time.Now()

	res, err := client.Account.PostAuth(&account.PostAuthParams{
		Context:   cxt,
		AuthToken: generateRandomToken(),
		UserHash:  generateRandomToken(),
	})

	after := time.Now()
	ms := (after.UnixNano() - before.UnixNano()) / 1000 / 1000

	println(c, ms)

	if err != nil {
		println(err.Error())
	} else {
		users[index] = basicAccount{
			user:        *res.Payload.User,
			accessToken: res.Payload.AccessToken,
		}
	}
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
	res := randStringBytes(128)
	return res
}
