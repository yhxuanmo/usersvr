package pg

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"usersvr/utils"
	"usersvr/utils/db/model"
)

var PgClient *gorm.DB

func init(){
	var err error
	pgConf := utils.Config.Postgresql
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		pgConf.Host, pgConf.Port, pgConf.User, pgConf.DbName, pgConf.Password, pgConf.Sslmode)
	PgClient, err =gorm.Open("postgres", args)
	if err != nil {
		log.Fatalln("connect db error: ", err)
	}

	PgClient.LogMode(utils.Config.Server.Debug)
	PgClient.AutoMigrate(&model.User{})
}
