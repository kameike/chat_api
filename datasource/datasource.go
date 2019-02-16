package datasource

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	. "github.com/kameike/chat_api/error"
	. "github.com/kameike/chat_api/model"
)

type DataSourceDescriptor interface {
	RDB() *gorm.DB
	MigrateIfNeed() ChatAPIError
	CheckHealth() (string, bool)
	Close()
}

type appDatasourceDescriptor struct {
	db *gorm.DB
}

func (d *appDatasourceDescriptor) RDB() *gorm.DB {
	return d.db
}

func (d *appDatasourceDescriptor) Close() {
	d.RDB().Close()
}

func (d *appDatasourceDescriptor) CheckHealth() (string, bool) {
	rdsHealth := true
	redisHealth := true
	code := http.StatusOK

	if err := d.pingRDB(); err != nil {
		rdsHealth = false
	}
	// if err := pingRedis(); err != nil {
	// 	redisHealth = false
	// }

	if (rdsHealth && redisHealth) == false {
		code = http.StatusServiceUnavailable
	}

	codeDesc := fmt.Sprintf("code: %d\n", code)
	rdsMsg := fmt.Sprintf("Is RDS available \t=> %t\n", rdsHealth)
	redisMsg := fmt.Sprintf("Is redis available\t=> %t\n", rdsHealth)

	msg := codeDesc + rdsMsg + redisMsg

	return msg, (redisHealth || rdsHealth)
}

func (d *appDatasourceDescriptor) MigrateIfNeed() ChatAPIError {
	d.db.CreateTable(&User{})
	d.db.CreateTable(&AccessToken{})
	d.db.CreateTable(&UserChatRoom{})
	d.db.CreateTable(&ChatRoom{})
	d.db.CreateTable(&Message{})
	return nil
}

func PrepareDatasource() DataSourceDescriptor {
	env := GetEnvs()
	dbUrl := env.dbAddr
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}

	return &appDatasourceDescriptor{
		db: db,
	}
}

func PrepareInmemoryDatasource() DataSourceDescriptor {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err.Error())
	}

	return &appDatasourceDescriptor{
		db: db,
	}
}

func (d *appDatasourceDescriptor) pingRDB() ChatAPIError {
	if d.db == nil {
		return ErroRDBConnection(nil)
	}
	if err := d.db.DB().Ping(); err != nil {
		return ErroRDBConnection(err)
	}
	return nil
}
