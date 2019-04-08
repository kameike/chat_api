package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"sync"
	"time"

	"github.com/go-openapi/runtime"
	httpclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/kameike/chat_api/swggen/apimodel"
	apiclient "github.com/kameike/chat_api/swggen/client"
	"github.com/kameike/chat_api/swggen/client/account"
	"github.com/kameike/chat_api/swggen/client/chat_rooms"
)

var userCount = 100

func main() {
	rand.Seed(time.Now().UnixNano())
	createAccounts()

	wg := sync.WaitGroup{}
	for i, u := range users {
		wg.Add(1)
		time.Sleep(1000 * 1000 * 10 * time.Duration(i))
		uc := u
		go func() {
			updateProfiles(uc)
			wg.Done()
		}()
	}
	wg.Wait()

	for i, u := range users {
		wg.Add(1)
		go func() {
			time.Sleep(1000 * 1000 * 1000 * time.Duration(i))
			room := make([]roomInfo, 10, 10)
			for i := 0; i < 10; i++ {
				r := randomRoomInfo(u)
				room[i] = r
			}
			before := time.Now().UnixNano()
			postRoomRequest(room, u)
			after := time.Now().UnixNano()
			ms := (after - before) / 1000 / 1000
			println(ms)
			wg.Done()
		}()
	}
	wg.Wait()
}

var users = make([]basicAccount, userCount, userCount)
var rooms = make(map[string]roomInfo)

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

func randomRoomInfo(u1 basicAccount) roomInfo {
	u2 := users[rand.Intn(len(users))]
	name := randStringBytes(10)

	return roomInfo{
		user1: u1,
		user2: u2,
		name:  name,
	}
}

func authFor(a basicAccount) runtime.ClientAuthInfoWriter {
	return httpclient.APIKeyAuth("x_chat_access_token", "header", a.accessToken)
}

func roomString(r roomInfo) string {
	data := struct {
		Users    []string `json:"users"`
		RoomId   string   `json:"roomId"`
		RoomName string   `json:"roomName"`
	}{
		Users: []string{
			r.user1.user.Hash,
			r.user2.user.Hash,
		},
		RoomId: r.name,
	}

	d, err := json.Marshal(data)
	if err != nil {
		panic(err.Error())
	}

	return string(d)
}

func roomsString(r []roomInfo) []string {
	s := make([]string, len(r), len(r))

	for i, rr := range r {
		s[i] = roomString(rr)
	}

	return s
}

func postRoomRequest(r []roomInfo, u basicAccount) {
	res, err := client.ChatRooms.PostChatrooms(&chat_rooms.PostChatroomsParams{
		Body: &apimodel.ChatroomRequest{
			Request: roomsString(r),
		},
		Context: cxt,
	}, authFor(u))

	if err != nil {
		println(err.Error())
	} else {
		println("====")
		println(len(r))
		for _, v := range res.Payload {
			r := rooms[v.Name]
			r.hash = v.ID
			rooms[r.name] = r
			println(v.ID)
		}
	}
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
	for i := 0; i < userCount; i++ {
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
var transport = httpclient.New(host, "", nil)
var client = apiclient.New(transport, strfmt.Default)

func testCreateAcoount(index int) {
	res, err := client.Account.PostAuth(&account.PostAuthParams{
		Context:   cxt,
		AuthToken: generateRandomToken(),
		UserHash:  generateRandomToken(),
	})

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
