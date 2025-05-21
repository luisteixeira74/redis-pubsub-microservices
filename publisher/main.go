package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379", // nome do serviço no docker-compose
	})

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	words := []string{"golang", "docker", "redis", "pubsub", "microservice"}

	log.Println("🚀 Publisher started...")

	// Forma idiomática com for range
	for t := range ticker.C {
		word := words[rand.Intn(len(words))]
		err := rdb.Publish(ctx, "random-words", word).Err()
		if err != nil {
			log.Printf("❌ Failed to publish: %v", err)
			continue
		}
		log.Printf("📢 [%s] Published: %s", t.Format("15:04:05"), word)
	}

	// Se quiser escutar mais de um canal (ex: cancelamento via contexto), use:
	/*
		for {
			select {
			case t := <-ticker.C:
				word := words[rand.Intn(len(words))]
				err := rdb.Publish(ctx, "random-words", word).Err()
				if err != nil {
					log.Printf("❌ Failed to publish: %v", err)
					continue
				}
				log.Printf("📢 [%s] Published: %s", t.Format("15:04:05"), word)

			case <-ctx.Done():
				log.Println("🛑 Context cancelled, exiting publisher...")
				return
			}
		}
	*/
}
