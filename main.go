package main

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisClient.Ping().Result()
	fmt.Println(pong, err)
	pubSub := redisClient.Subscribe(os.Getenv("REUNI_SERVICE") + "_" + os.Getenv("REUNI_NAMESPACE"))

	initContext()
	initGracefulShutdown()
	handleSync(pubSub)
}
