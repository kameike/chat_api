package main

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

type envValues struct {
	redisAddr string
	redisPass string
	dbAddr    string
	dbPass    string
	dbUser    string
}

func getEnvs() envValues {
	tet()
	if localEnvCache != nil {
		return *localEnvCache
	}
	target := envValues{
		redisAddr: getEnv("CHAT_REDIS_ADDR"),
		redisPass: getSecretEnv("CHAT_REDIS_PASS"),
		dbAddr:    getEnv("CHAT_RDS_ADDR"),
		dbPass:    getSecretEnv("CHAT_RDS_PASS"),
		dbUser:    getEnv("CHAT_RDS_USER"),
	}
	localEnvCache = &target
	return target
}

func getSecretEnv(target string) string {
	result, hasSet := os.LookupEnv(target)
	if !hasSet {
		panic(target + " has not set")
	}

	println(target + " => xxxxxxxx")

	return result
}

func getEnv(target string) string {
	result, hasSet := os.LookupEnv(target)
	if !hasSet {
		panic(target + " has not set")
	}

	println(target + " => " + result)

	return result
}

var localEnvCache *envValues

func main() {
	// check if all required env has set
	_ = getEnvs()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/health", checkHealth)

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

func checkHealth(c echo.Context) error {
	rdsHealth := true
	redisHealth := true
	code := http.StatusOK

	if err := pingRds(); err != nil {
		rdsHealth = false
	}
	if err := pingRedis(); err != nil {
		redisHealth = false
	}

	if (rdsHealth && redisHealth) == false {
		code = http.StatusServiceUnavailable
	}

	codeDesc := fmt.Sprintf("code: %d\n", code)
	rdsMsg := fmt.Sprintf("Is RDS available \t=> %t\n", rdsHealth)
	redisMsg := fmt.Sprintf("Is redis available\t=> %t\n", rdsHealth)

	c.String(code, codeDesc+rdsMsg+redisMsg)
	return nil
}

func pingRedis() error {
	env := getEnvs()
	client := redis.NewClient(&redis.Options{
		Addr:     env.redisAddr,
		Password: env.redisPass,
		DB:       0, // use default DB
	})
	_, err := client.Ping().Result()
	return err
}

func pingRds() error {
	env := getEnvs()
	dbUrl := fmt.Sprintf("%s:%s@%s", env.dbUser, env.dbPass, env.dbAddr)
	db, err := sql.Open("mysql", dbUrl)
	defer db.Close()

	if err != nil {
		return err
	}
	return db.Ping()
}
