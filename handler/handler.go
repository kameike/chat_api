package handler

import (
	"net/http"

	"github.com/kameike/chat_api/apierror"

	"github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/repository"
)

func SetUpHandler() *RequestHandler {
	return &RequestHandler{
		p: repository.CreateAppRepositoryProvider(),
	}
}

type RequestHandler struct {
	p repository.ReposotryProvidable
}

// TODO: いい感じにエラーを整形する
func errorResponseWithCode(code int, message string) middleware.ResponderFunc {
	return func(res http.ResponseWriter, pro runtime.Producer) {
		res.WriteHeader(code)
		res.Write([]byte(message))
	}
}

func errorResponse(err apierror.ChatAPIError) middleware.ResponderFunc {
	return errorResponseWithCode(500, err.ErrorMessage())
}

func notHealthy(message string) middleware.ResponderFunc {
	return func(res http.ResponseWriter, pro runtime.Producer) {
		res.WriteHeader(503)
		res.Write([]byte(message))
	}
}
