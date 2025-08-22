package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// map1 := make(map[string]map[string]any)
	// map1["user:1"] = map[string]any{"name": "jiten", "age": 42, "email": "jitenp@outlook.com"}

	// ctx := context.Background()
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:17001,localhost:17000,localhost:17002",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	// fmt.Println(rdb)

	// context.TODO()
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGILL)

	for {
		select {
		case <-sigCh:
			cancel()
		case <-ctx.Done():
			time.Sleep(time.Second * 10)
			fmt.Println("Closing the system due to control C call")
			return
		default:
			Run(ctx)
		}
	}

}

func Run(ctx context.Context) {
	time.Sleep(time.Second * 1)
	println(time.Now().Unix())
}
