package handler

import (
	"github.com/kameike/chat_api/model"
	"github.com/kameike/chat_api/swggen/apimodel"

	_ "github.com/joho/godotenv/autoload"
)

func mapUsers(users []model.User) []*apimodel.Account {
	result := make([]*apimodel.Account, len(users), len(users))

	for i, u := range users {
		result[i] = mapUser(u)
	}
	return result
}
