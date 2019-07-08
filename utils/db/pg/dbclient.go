package pg

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"usersvr/utils/db/model"
)

var PgClient *gorm.DB

func init(){
	var err error
	PgClient, err =gorm.Open("postgres", "host=127.0.0.1 port=5432 user=iot dbname=iot password=123456 sslmode=disable")
	if err != nil {
		log.Fatalln("connect db error: ", err)
	}

	PgClient.LogMode(true)
	PgClient.AutoMigrate(&model.User{})
}
