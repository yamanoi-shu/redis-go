package main

import "github.com/go-redis/redis/v8"

func newRedisCli() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "10.0.224.4:6379",
	})
}
