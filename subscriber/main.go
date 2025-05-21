package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379", // nome do serviço no docker-compose
	})

	pubsub := rdb.Subscribe(ctx, "random-words")
	defer pubsub.Close()

	log.Println("📥 Subscriber listening on channel: random-words")

	ch := pubsub.Channel()

	for msg := range ch {
		log.Printf("📝 Received: %s", msg.Payload)
	}
}
