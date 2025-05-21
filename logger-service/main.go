package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379", // nome do serviço Redis no docker-compose
	})

	subscriber := rdb.Subscribe(ctx, "random-words")
	ch := subscriber.Channel()

	log.Println("📝 Logger-service started and listening on channel: random-words")

	for msg := range ch {
		log.Printf("🗂️ Logger received: %s\n", msg.Payload)
	}
}
