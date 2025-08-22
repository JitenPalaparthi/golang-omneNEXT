# 🚀 Redis Concepts with Examples

Redis is an **in-memory key-value store**. But unlike a plain cache, it supports rich **data structures**.

---

## 🔹 1. Strings (basic key-value)
- Simplest type: holds text or binary (up to 512 MB).
- Common commands: `SET`, `GET`, `INCR`, `APPEND`.

### Example (CLI)
```bash
SET name "Alice"
GET name
INCR counter
GET counter
```

### Example (Go)
```go
rdb.Set(ctx, "name", "Alice", 0)
val, _ := rdb.Get(ctx, "name").Result()
fmt.Println(val) // Alice
```

---

## 🔹 2. Lists (ordered collection, like linked list)
- Commands: `LPUSH`, `RPUSH`, `LPOP`, `LRANGE`.

### Example (CLI)
```bash
LPUSH tasks "task1"
LPUSH tasks "task2"
RPUSH tasks "task3"
LRANGE tasks 0 -1
```
👉 returns: `["task2", "task1", "task3"]`

### Example (Go)
```go
rdb.LPush(ctx, "tasks", "task1", "task2")
vals, _ := rdb.LRange(ctx, "tasks", 0, -1).Result()
fmt.Println(vals)
```

---

## 🔹 3. Hashes (like a map/dictionary)
- Good for storing objects (key → field-value pairs).
- Commands: `HSET`, `HGET`, `HGETALL`.

### Example (CLI)
```bash
HSET user:1 name "Alice" age "30" city "Hyderabad"
HGET user:1 name
HGETALL user:1
```

### Example (Go)
```go
rdb.HSet(ctx, "user:1", "name", "Alice", "age", "30", "city", "Hyderabad")
val, _ := rdb.HGet(ctx, "user:1", "name").Result()
fmt.Println(val)
```

---

## 🔹 4. Sets (unique unordered values)
- Commands: `SADD`, `SMEMBERS`, `SISMEMBER`.

### Example (CLI)
```bash
SADD tags "golang" "redis" "docker"
SMEMBERS tags
SISMEMBER tags "golang"
```

### Example (Go)
```go
rdb.SAdd(ctx, "tags", "golang", "redis", "docker")
members, _ := rdb.SMembers(ctx, "tags").Result()
fmt.Println(members)
```

---

## 🔹 5. Sorted Sets (Set + score → ordered)
- Good for leaderboards, rankings.
- Commands: `ZADD`, `ZRANGE`, `ZREVRANGE`, `ZSCORE`.

### Example (CLI)
```bash
ZADD leaderboard 100 "Alice" 200 "Bob" 150 "Charlie"
ZRANGE leaderboard 0 -1 WITHSCORES
ZREVRANGE leaderboard 0 -1 WITHSCORES
```

### Example (Go)
```go
rdb.ZAdd(ctx, "leaderboard", 
    redis.Z{Score: 100, Member: "Alice"},
    redis.Z{Score: 200, Member: "Bob"},
)
vals, _ := rdb.ZRevRangeWithScores(ctx, "leaderboard", 0, -1).Result()
fmt.Println(vals)
```

---

## 🔹 6. Streams (append-only log)
- Great for event sourcing, message queues.
- Commands: `XADD`, `XRANGE`, `XREAD`.

### Example (CLI)
```bash
XADD mystream * user "Alice" action "login"
XRANGE mystream - +
XREAD COUNT 1 STREAMS mystream 0
```

---

## 🔹 7. Pub/Sub (publish-subscribe messaging)
- Commands: `PUBLISH`, `SUBSCRIBE`.

### Example (CLI)
Terminal 1:
```bash
SUBSCRIBE news
```
Terminal 2:
```bash
PUBLISH news "Redis 7 released!"
```

---

## 🔹 8. Bitmaps (bit-level operations on strings)
- Commands: `SETBIT`, `GETBIT`, `BITCOUNT`.

### Example (CLI)
```bash
SETBIT userlogins 7 1
GETBIT userlogins 7
BITCOUNT userlogins
```

---

## 🔹 9. HyperLogLog (approximate cardinality counter)
- Estimates unique items.
- Commands: `PFADD`, `PFCOUNT`.

### Example (CLI)
```bash
PFADD visitors "alice" "bob" "charlie"
PFADD visitors "alice"
PFCOUNT visitors
```
👉 returns `3`

---

## 🔹 10. Geospatial Indexes
- Store locations with latitude/longitude.
- Commands: `GEOADD`, `GEORADIUS`, `GEODIST`.

### Example (CLI)
```bash
GEOADD places 77.5946 12.9716 "Bangalore" 72.8777 19.0760 "Mumbai"
GEODIST places Bangalore Mumbai km
```

---

# ✅ Summary Table

| Type       | Key Commands                          | Use Case                          |
|------------|---------------------------------------|-----------------------------------|
| String     | `SET`, `GET`, `INCR`                  | Caching, counters                 |
| List       | `LPUSH`, `LRANGE`                     | Queues, task lists                |
| Hash       | `HSET`, `HGETALL`                     | Objects (user profiles)           |
| Set        | `SADD`, `SMEMBERS`                    | Tags, unique values               |
| Sorted Set | `ZADD`, `ZREVRANGE`                   | Leaderboards, rankings            |
| Stream     | `XADD`, `XREAD`                       | Event streams, logs               |
| Pub/Sub    | `PUBLISH`, `SUBSCRIBE`                | Real-time messaging               |
| Bitmap     | `SETBIT`, `BITCOUNT`                  | Tracking logins, attendance       |
| HyperLogLog| `PFADD`, `PFCOUNT`                    | Unique visitor counts             |
| Geo        | `GEOADD`, `GEORADIUS`                 | Location-based queries            |
