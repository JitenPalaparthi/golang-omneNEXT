package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
	defer rdb.Close()

	pubsub := rdb.Subscribe(ctx, "news")
	defer pubsub.Close()

	// Wait for subscription to be created
	if _, err := pubsub.Receive(ctx); err != nil {
		log.Fatal(err)
	}

	ch := pubsub.Channel() // <-chan *redis.Message
	for {
		select {
		case <-ctx.Done():
			fmt.Println("shutting down")
			return
		case m := <-ch:
			if m == nil {
				return
			}
			fmt.Printf("channel=%s payload=%s\n", m.Channel, m.Payload)
		}
	}
}
