package main

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	f, err := os.Create("polygon-address.txt")
	if err != nil {
		panic(err)
	}

	data, err := rdb.HKeys(context.Background(), ":tokens").Result()
	if err != nil {
		panic(err)
	}

	for i := range data {
		f.WriteString(data[i] + "\n")
	}
}
