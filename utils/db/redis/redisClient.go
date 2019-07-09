package redis

import (
	"github.com/go-redis/redis"
	"log"
)

var RedisClient *redis.Client

func init(){
	//var err error
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password:"",
		DB: 0,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalln("",err)
	}
}
