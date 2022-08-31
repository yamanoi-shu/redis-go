package main

import (
	"context"
	"log"
	"time"
)

func main() {
	client := newRedisCli()
	ctx := context.Background()
	for {
		result := client.Ping(ctx)
		log.Println("result: ", result.String())
		if err := result.Err(); err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second * 1)
	}
}
