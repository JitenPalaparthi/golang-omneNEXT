package main

import (
	"context"
	"fmt"
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
	//ctx, _ := context.WithDeadline(parent, time.Now().Add(time.Second*5))
	ctx, cancel := context.WithCancel(parent)
	go func() {
		time.Sleep(time.Second * 5)
		cancel()
	}()
	//defer cancel()
	for {
		select {
		case <-ctx.Done():
			time.Sleep(time.Second * 5)
			fmt.Println("Closing the system due to control C call")
			return
		default:
			Run(ctx)
			go Cancelled(ctx)
		}
	}

}

func Run(ctx context.Context) {
	go func() {
		<-ctx.Done()
		fmt.Println("I too got to know that the context is cancelled")
	}()
	time.Sleep(time.Second * 1)
	println(time.Now().Unix())
}

func Cancelled(ctx context.Context) {
	<-ctx.Done()
	println("yes It is done")
}
