package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

func initRedis() *redis.Client {
	redisDb := redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	return redisDb
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	redisConn := initRedis()
	redisConn.FlushAll()

	var totalInserted int64

	for x := 0; x < 10000; x++ {

		randID := randString(4)

		inserted := redisConn.SAdd("commit", randID).Val()
		totalInserted += inserted

		redisConn.PFAdd("commitHLL", randID)
	}

	val := redisConn.PFCount("commitHLL").Val()

	fmt.Println("Real:", totalInserted)
	fmt.Println("HLL estimate:", val)
	fmt.Println("Difference:", totalInserted-val)
	fmt.Println("Error (%):", float64(totalInserted-val)*100/float64(totalInserted))

}
