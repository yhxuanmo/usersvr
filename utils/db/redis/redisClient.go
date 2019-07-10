package redis

import (
	"github.com/go-redis/redis"
	"log"
	"usersvr/utils"
)

var RedisClient *redis.Client

func init(){
	//var err error
	redisConf := utils.Config.Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisConf.Addr,
		Password: redisConf.Password,
		DB: redisConf.Db,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalln("",err)
	}
}
