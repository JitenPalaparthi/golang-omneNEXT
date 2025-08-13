# Kafka Producer Acks — Latency vs Durability

| Acks mode | franz-go option                         | Broker waits for…                   | Latency | Durability | Typical use |
|---|---|---|---|---|---|
| **0**     | `kgo.RequiredAcks(kgo.NoAck())`        | **No one** (fire-and-forget)        | ⭐ **Lowest** | 🔴 **Weak** (loss possible) | Ultra-low latency, non-critical telemetry you can lose |
| **1**     | `kgo.RequiredAcks(kgo.LeaderAck())`    | **Leader only**                     | ⭐⭐ Low | 🟡 Medium (loss if leader dies before replication) | High-throughput, acceptable small loss |
| **all**   | `kgo.RequiredAcks(kgo.AllISRAcks())`   | **All in-sync replicas (ISR)**      | ⭐⭐⭐ Highest | 🟢 **Strong** | Payments, orders, anything you mustn’t lose |

## Visual intuition

```
Producer --> Leader --> Followers
          (acks=1)     (replicate async)

Producer --> Leader --wait--> Followers (ISR) --ack back-->
               (acks=all, stronger durability)
```

## Recommended combos

- **Critical events** (orders, payments, ledger)  
  `acks=all` + `kgo.IdempotentProducer(true)` + sensible retries  
  Consider keys for per-entity ordering.

- **Throughput-heavy but okay to lose some** (clickstream, noisy logs)  
  `acks=1` (or `all` if durability matters more)  
  Keep idempotence if retries can happen.

- **Fire-and-forget spikes** (best-effort metrics with backup path)  
  `acks=0`, bounded channel backpressure, drop policy or local buffer.

## franz-go snippet switcher

```go
opts := []kgo.Opt{
  kgo.SeedBrokers(brokers...),
  // pick ONE:
  // kgo.RequiredAcks(kgo.NoAck()),
  // kgo.RequiredAcks(kgo.LeaderAck()),
  kgo.RequiredAcks(kgo.AllISRAcks()),
  kgo.IdempotentProducer(true),     // strongly recommended for prod
  kgo.RecordRetries(5),
}
client, _ := kgo.NewClient(opts...)
```
