package utils

import (
    // "context"
    "fmt"
    // "log"
    "time"

    "github.com/omniful/go_commons/redis"
)

func InitRedis() {
    // Define the Redis configuration
    cfg := &redis.Config{
        ClusterMode:                    false,
        ServeReadsFromSlaves:           false,
        ServeReadsFromMasterAndSlaves:  false,
        PoolSize:                       50,
        PoolFIFO:                       true,
        MinIdleConn:                    6,
        DB:                             0,
        Hosts:                          []string{"localhost:6379"},
        DialTimeout:                    500 * time.Millisecond,
        ReadTimeout:                    2000 * time.Millisecond,
        WriteTimeout:                   2000 * time.Millisecond,
        IdleTimeout:                    600 * time.Second,
    }

    // Initialize the Redis client
    client := redis.NewClient(cfg)

    // Create a context
    // ctx := context.Background()
    fmt.Println("Connected to Redis!",client)
    // Example of setting a key
    // err := client.Set(ctx, "key", "value", 0).Err()
    // if err != nil {
    //     log.Fatalf("Failed to set key: %v", err)
    // }

    // // Example of getting a key
    // val, err := client.Get(ctx, "key").Result()
    // if err != nil {
    //     log.Fatalf("Failed to get key: %v", err)
    // }
    // fmt.Println("key:", val)
}