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
	"github.com/kameike/chat_api/swggen/client/chatrooms"
	"github.com/kameike/chat_api/swggen/client/messages"
)

var userCount = 3
var reqPerSec = 5
var wtime = time.Duration(1000 / reqPerSec)

var users = make([]basicAccount, userCount, userCount)
var rooms = make(map[string]roomInfo)
var userRooms = make(map[int64][]roomInfo)

var cxt = context.Background()

// var host = "dev-chat.taimee.co.jp"
// var transport = httpclient.New(host, "", []string{"https"})

var host = "localhost"
var transport = httpclient.New(host, "", []string{"http"})

// var host = "13.231.204.249"
// var transport = httpclient.New(host, "", []string{"http"})

// var host = "chat-stg-aagktp6bnbxvfrw8.stg-taimee.com"
// var transport = httpclient.New(host, "", []string{"http"})

var client = apiclient.New(transport, strfmt.Default)

func main() {
	rand.Seed(time.Now().UnixNano())
	println("create user")
	createAccounts()

	println("udpate user")
	wg := sync.WaitGroup{}
	for i, u := range users {
		uc := u
		id := i
		wg.Add(1)
		go func() {
			time.Sleep(1000 * 1000 * wtime * time.Duration(id))
			updateProfiles(uc)
			wg.Done()
		}()
	}
	wg.Wait()

	for _, u := range users {
		count := 3
		room := make([]roomInfo, count, count)
		for i := 0; i < count; i++ {
			r := randomRoomInfo(u)
			room[i] = r
			rooms[r.name] = r
		}
		userRooms[u.user.ID] = room
	}

	for k := 0; k < 1; k++ {
		for i, u := range users {
			wg.Add(1)
			id := i
			us := u
			go func() {
				time.Sleep(1000 * 1000 * 30 * time.Duration(id*k))
				benchmark("get rooms", func() {
					postRoomRequest(userRooms[us.user.ID], us)
				})
				wg.Done()
			}()
		}
	}
	wg.Wait()

	for k := 0; k < 1; k++ {
		for i, u := range users {
			wg.Add(1)
			id := i
			us := u
			go func() {
				time.Sleep(1000 * 1000 * 30 * time.Duration(id*k))
				benchmark("get rooms", func() {
					postRoomRequest(userRooms[us.user.ID], us)
				})
				wg.Done()
			}()
		}
	}
	wg.Wait()

	for k := 0; k < 1; k++ {
		for i, u := range users {
			wg.Add(1)
			id := i
			us := u
			go func() {
				time.Sleep(1000 * 1000 * 30 * time.Duration(id*k))
				benchmark("get rooms", func() {
					postRoomRequest(userRooms[us.user.ID], us)
				})
				wg.Done()
			}()
		}
	}
	wg.Wait()

	for k := 0; k < 1; k++ {
		for i, u := range users {
			wg.Add(1)
			id := i
			us := u
			go func() {
				time.Sleep(1000 * 1000 * wtime * time.Duration(id*k))
				benchmark("get rooms", func() {
					postRoomRequest(userRooms[us.user.ID], us)
				})
				wg.Done()
			}()
		}
	}
	wg.Wait()

	counter := 0
	for _, r := range rooms {
		counter++
		c := counter
		wg.Add(1)
		rm := r
		if rm.hash == "" {
			continue
		}
		go func() {
			time.Sleep(1000 * 1000 * wtime * time.Duration(c))
			println(rm.hash, rm.name)
			benchmark("post chat", func() {
				postMessage(rm)
			})
			wg.Done()
		}()
	}
	wg.Wait()
}

func benchmark(name string, proc func()) {
	before := time.Now().UnixNano()
	proc()
	after := time.Now().UnixNano()
	ms := (after - before) / 1000 / 1000
	println(ms, "ms:", name)
}

func postMessage(r roomInfo) {
	if r.hash == "" {
		println("not effective room")
		return
	}
	_, err := client.Messages.PostChatroomsChatroomHashMessages(&messages.PostChatroomsChatroomHashMessagesParams{
		Body: &apimodel.MessageRequest{
			Content: "test",
		},
		ChatroomHash: r.hash,
		Context:      cxt,
	}, authFor(r.user1))

	if err != nil {
		println("err", err.Error())
	}
}

type basicAccount struct {
	accessToken string
	user        apimodel.Account
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
	return httpclient.APIKeyAuth("X-CHAT-ACCESS-TOKEN", "header", a.accessToken)
}

func roomString(r roomInfo) string {
	data := struct {
		Accounts []string `json:"accountHashList"`
		RoomName string   `json:"channelName"`
	}{
		Accounts: []string{
			r.user1.user.Hash,
			r.user2.user.Hash,
		},
		RoomName: r.name,
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
	res, err := client.Chatrooms.PostChatrooms(&chatrooms.PostChatroomsParams{
		Body: &apimodel.ChatroomRequest{
			Chatrooms: roomsString(r),
		},
		Context: cxt,
	}, authFor(u))

	if err != nil {
		println(err.Error())
	} else {
		for _, v := range res.Payload.Chatrooms {
			r := rooms[v.Name]
			r.hash = v.Hash
			rooms[v.Name] = r
			println("room ->", rooms[v.Name].hash)
		}
	}
}

func updateProfiles(a basicAccount) {
	newName := randStringBytes(30)

	res, err := client.Account.PostProfile(&account.PostProfileParams{
		Context: cxt,
		Body: account.PostProfileBody{
			Name: newName,
		},
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
		c := i
		go func() {
			time.Sleep(1000 * 1000 * wtime * time.Duration(c))
			testCreateAcoount(c)
			wg.Done()
		}()
	}
	wg.Wait()
}

func testCreateAcoount(index int) {
	res, err := client.Account.PostAuth(&account.PostAuthParams{
		Context: cxt,
		Body: account.PostAuthBody{
			AuthToken:   generateRandomToken(),
			AccountHash: generateRandomToken(),
		},
	})

	if err != nil {
		println(err.Error())
	} else {
		users[index] = basicAccount{
			user:        *res.Payload.Account,
			accessToken: res.Payload.AccessToken,
		}
		println(users[index].accessToken)
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
