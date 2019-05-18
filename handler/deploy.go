package handler

import (
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/swggen/restapi/operations/deploy"
)

func (h *RequestHandler) DeployGetHealthHandler() deploy.GetHealthHandlerFunc {
	return func(params deploy.GetHealthParams) middleware.Responder {
		repo := h.p
		message, isHealthy := repo.CheckHealth()
		if isHealthy {
			return deploy.NewGetHealthOK().WithPayload(message)
		} else {
			return notHealthy(message)
		}
	}
}
