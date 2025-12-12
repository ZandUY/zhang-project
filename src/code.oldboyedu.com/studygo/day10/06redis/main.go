package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:32769",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping(context.Background()).Result()
	return
}
func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("Redis conn failed,err: %v", err)
		return
	}
	fmt.Println("连接redis成功!")
}
