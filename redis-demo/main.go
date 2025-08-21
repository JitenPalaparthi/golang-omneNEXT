package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	// map1 := make(map[string]map[string]any)
	// map1["user:1"] = map[string]any{"name": "jiten", "age": 42, "email": "jitenp@outlook.com"}

	ctx := context.Background()

	// cluster := redis.NewClusterClient(&redis.ClusterOptions{
	// 	Addrs: []string{"localhost:7000", "localhost:7001", "localhost:7002"},
	// 	// Username:  "", Password: "", // if needed
	// 	// TLSConfig: &tls.Config{},    // if you use TLS
	// })

	//fmt.Println(cluster)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // or your VM IP if not forwarded
		// Password:  "",                // set if you configured one
		// DB:        0,
		DialTimeout:  2 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		PoolSize:     20,
		MinIdleConns: 4,
	})
	defer rdb.Close()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal(err)
	} else {
		println("Connected")
	}

	err := rdb.Set(ctx, "company-name", "OmneNEXT Tecnhologies", 0).Err()
	if err != nil {
		println(err.Error(), "key could not be set up")
	}

	sc, err := rdb.Get(ctx, "company-name").Result()
	if err != nil {
		println(err.Error())
	} else {
		println("---", sc)
	}

	rdb.LPush(ctx, "tasks", "golang", "kafka", "redis")
	vals, _ := rdb.LRange(ctx, "tasks", 0, -1).Result()
	fmt.Println(vals)

	//ch := make(chan string)
	sig := make(chan struct{})
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(time.Second * 10)
			rdb.Publish(ctx, "news", "Hello World-->"+fmt.Sprint(i))
		}
		sig <- struct{}{}
	}()
	<-sig

	//rdb.Close()

}
