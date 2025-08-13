// cmd/consumer/main.go
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

var brokers = []string{
	"localhost:19092", "localhost:29092", "localhost:39092",
}

const (
	topic   = "demo"
	groupID = "demo-group"
)

func main() {
	// Build client
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumeTopics(topic),   // subscribe to topics
		kgo.ConsumerGroup(groupID), // join consumer group
		kgo.DisableAutoCommit(),    // we'll commit after processing
		// kgo.AutoCommitMarks(),            // (alternative) enable auto commit of marked offsets
	)
	if err != nil {
		log.Fatalf("new client: %v", err)
	}
	defer cl.Close()

	// Shutdown on SIGINT/SIGTERM
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	log.Printf("consumer started: group=%s topic=%s", groupID, topic)

	for {
		select {
		case <-ctx.Done():
			log.Println("shutting down…")
			return
		default:
		}

		fetches := cl.PollFetches(ctx)

		// If the context is canceled during Poll, stop cleanly
		if err := fetches.Err(); err != nil && ctx.Err() != nil {
			break
		}

		// Handle any partition-level errors (authorization, unknown topic, etc.)
		for _, fe := range fetches.Errors() {
			log.Printf("fetch error on %s[%d]: %v", fe.Topic, fe.Partition, fe.Err)
		}

		// Process records
		var toCommit []*kgo.Record
		fetches.EachRecord(func(r *kgo.Record) {
			// Access key/value
			key := string(r.Key)
			val := string(r.Value)

			// Read headers (optional)
			hdrs := map[string]string{}
			for _, h := range r.Headers {
				hdrs[h.Key] = string(h.Value)
			}

			// ---- Your business logic ----
			if err := handleMessage(r.Topic, r.Partition, r.Offset, key, val, hdrs); err != nil {
				// If processing fails, DO NOT commit this record.
				// Add retry/DLQ logic here if needed.
				log.Printf("process failed @ %s[%d]@%d: %v", r.Topic, r.Partition, r.Offset, err)
				return
			}
			// Mark on success
			toCommit = append(toCommit, r)
		})

		// Synchronous commit of all successfully processed records
		if len(toCommit) > 0 {
			cl.MarkCommitRecords(toCommit...)
			cctx, ccancel := context.WithTimeout(ctx, 5*time.Second)
			if err := cl.CommitMarkedOffsets(cctx); err != nil {
				log.Printf("commit failed: %v", err)
			}
			ccancel()
		}
	}
}

func handleMessage(topic string, partition int32, offset int64, key, val string, headers map[string]string) error {
	// TODO: implement your logic. Keep it idempotent.
	log.Printf("got %s[%d]@%d key=%q val=%q headers=%v", topic, partition, offset, key, val, headers)
	return nil
}
