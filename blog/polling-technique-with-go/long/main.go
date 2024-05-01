package main

import (
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func upstream() string {
	// Try to get a random name 5 times
	// Non-empty represents a new data, otherwise old data
	for i := 0; i < 5; i++ {
		result := gofakeit.RandomString([]string{"", gofakeit.Name()})
		if result != "" {
			return result
		}
		// Simulate a delay (waiting for new data)
		time.Sleep(time.Second)
	}
	return ""
}

func downstream(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		// Polling upstream data with long polling technique
		default:
			result := upstream()
			if result != "" {
				fmt.Println("Got result:", result)
			} else {
				fmt.Println("Got empty result")
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	downstream(ctx)
}
