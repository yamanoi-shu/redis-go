package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	now := time.Now().Unix()

	nowStr := strconv.FormatInt(now, 10)

	tmp := "awg:update_channel_slots.{%s}"
	key := fmt.Sprintf(tmp, "aaaaaaaaaaaaaaaaaa-1")

	expiration := time.Minute * 10

	rClient.SetNX(ctx, key, nowStr, expiration)

	fmt.Println(rClient.MemoryUsage(ctx, key))

}
