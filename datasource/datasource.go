package datasource

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "github.com/kameike/chat_api/apierror"
	. "github.com/kameike/chat_api/model"
)

type DataSourceDescriptor interface {
	RDB() *gorm.DB
	Begin()
	Commit()
	Rollback()
	MigrateIfNeed() ChatAPIError
	CheckHealth() (string, bool)
	Close()
}

type appDatasourceDescriptor struct {
	db *gorm.DB
	tx *gorm.DB
}

func (d *appDatasourceDescriptor) RDB() *gorm.DB {
	if d.tx != nil {
		return d.tx
	}
	return d.db
}

func (d *appDatasourceDescriptor) Close() {
	d.RDB().Close()
}

func (d *appDatasourceDescriptor) Rollback() {
	if d.tx != nil {
		d.tx.Rollback()
	} else {
		log.Fatalf("try rollback even tx not exist")
	}
	d.tx = nil
}

func (d *appDatasourceDescriptor) Commit() {
	if d.tx != nil {
		d.tx.Commit()
	} else {
		log.Fatalf("try rollback even tx not exist")
	}
	d.tx = nil
}

func (d *appDatasourceDescriptor) Begin() {
	if d.tx != nil {
		log.Fatalf("dose not support nested transaction")
	}
	d.tx = d.db.Begin()
}

func (d *appDatasourceDescriptor) CheckHealth() (string, bool) {
	rdsHealth := true
	redisHealth := true
	code := http.StatusOK

	if err := d.pingRDB(); err != nil {
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
	redisMsg := fmt.Sprintf("Is redis available\t=> %t\n", redisHealth)

	msg := codeDesc + rdsMsg + redisMsg

	return msg, (redisHealth && rdsHealth)
}

func pingRedis() error {
	env := GetEnvs()
	client := redis.NewClient(&redis.Options{
		Addr:     env.RedisAddr,
		Password: env.RedisPass,
		DB:       0,
	})
	_, err := client.Ping().Result()

	return err
}

func (d *appDatasourceDescriptor) MigrateIfNeed() ChatAPIError {
	d.db.CreateTable(&User{})
	d.db.CreateTable(&AccessToken{})
	d.db.CreateTable(&UserChatRoom{})
	d.db.CreateTable(&ChatRoom{})
	d.db.CreateTable(&Message{})

	d.db.Model(&User{}).ModifyColumn("url", "text")

	return nil
}

var mysqlDB *gorm.DB

func PrepareDatasource() DataSourceDescriptor {
	if mysqlDB == nil {
		db, err := gorm.Open("mysql", GetEnvs().DbAddr)
		if err != nil {
			panic(err.Error())
		}
		mysqlDB = db
	}
	// mysqlDB.LogMode(true)

	return &appDatasourceDescriptor{
		db: mysqlDB,
	}
}

func (d *appDatasourceDescriptor) pingRDB() ChatAPIError {
	if d.db == nil {
		return Error(SERVICE_DOWN, nil)
	}
	if err := d.db.DB().Ping(); err != nil {
		return Error(SERVICE_DOWN, err)
	}
	return nil
}
