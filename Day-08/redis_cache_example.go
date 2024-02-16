package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Create a new Redis client.
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Test the connection to Redis
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Connected to Redis:", pong)

	// Set a key-value pair in Redis
	err = client.Set(ctx, "mykey", "myvalue", 0).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
		return
	}

	// Get the value for the key from Redis
	val, err := client.Get(ctx, "mykey").Result()
	if err != nil {
		fmt.Println("Error getting key:", err)
		return
	}
	fmt.Println("Value for 'mykey':", val)
}
