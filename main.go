package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kameike/chat_api/repository"
	"github.com/labstack/echo"
)

func main() {
	repo := repository.CreateAppRepositoryProvider()
	defer repo.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/health", func(c echo.Context) error {
		msg, ok := repo.CheckHealth()

		if ok {
			return c.String(http.StatusOK, msg)
		} else {
			return c.String(http.StatusServiceUnavailable, msg)
		}
	})

	e.POST("/auth/createAccount", stub)
	e.POST("/auth/login", stub)

	e.POST("/chatrooms", stub) //find or create chatrooms
	e.GET("/chatrooms/{id}?id=hash", stub)
	e.POST("/chatrooms/{id}/message", stub)

	e.Logger.Fatal(e.Start(":1323"))
}

func stub(c echo.Context) error {
	return nil
}
