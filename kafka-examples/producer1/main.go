// cmd/producer/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

var brokers = []string{"localhost:19092", "localhost:29092", "localhost:39092"}

const topic = "demo"

func main() {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.RequiredAcks(kgo.AllISRAcks()),
		kgo.DisableIdempotentWrite(),
		//kgo.IdempotentProducer(true),            // enable idempotence
		kgo.RecordRetries(5),                    // retry sends
		kgo.ProducerBatchMaxBytes(64<<10),       // 64 KiB
		kgo.ProducerLinger(10*time.Millisecond), // small linger to batch
		// TLS/SASL options go here if you need them
	)
	if err != nil {
		log.Fatalf("new client: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	for i := 0; i < 5; i++ {
		r := &kgo.Record{
			Topic: topic,
			Key:   []byte(fmt.Sprintf("key-%02d", i)),
			Value: []byte(fmt.Sprintf("hello %d", i)),
			Headers: []kgo.RecordHeader{
				{Key: "source", Value: []byte("franz-go")},
			},
		}
		// async produce with a completion callback
		client.Produce(ctx, r, func(rec *kgo.Record, err error) {
			if err != nil {
				log.Printf("produce failed: %v", err)
				return
			}
			log.Printf("produced to %s[%d]@%d", rec.Topic, rec.Partition, rec.Offset)
		})
	}

	// flush pending records before exit
	client.Flush(ctx)
	log.Println("done")
}
