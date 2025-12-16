package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var redisdb *redis.Client
var ctx = context.Background()

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:32769",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping(ctx).Result()
	return
}
func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("Redis conn failed,err: %v", err)
		return
	}
	fmt.Println("连接redis成功!")
	key := "rank"
	items := []redis.Z{
		redis.Z{Score: 90, Member: "PHP"},
		redis.Z{Score: 96, Member: "Golang"},
		redis.Z{Score: 95, Member: "Python"},
		redis.Z{Score: 93, Member: "Java"},
	}
	redisdb.ZAdd(ctx, key, items...)
	newScore := redisdb.ZIncrBy(ctx, "rank", 10, "Golang")
	// if err!=nil{
	// 	fmt.Printf("Incr Golang failed,err: %v",err)
	// 	return
	// }
	fmt.Printf("Golang现在的分数是%d\n", newScore.Val())
}
