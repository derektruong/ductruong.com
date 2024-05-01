package main

import (
	"context"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func upstream() string {
	return gofakeit.RandomString([]string{"", gofakeit.Name()})
}

func downstream(ctx context.Context) {
	// Create a timer with 1 second interval
	interval := 1 * time.Second
	timer := time.NewTimer(interval)
	defer timer.Stop()
	// Start polling until the context is done
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			result := upstream()
			if result != "" {
				fmt.Println("Got result:", result)
			} else {
				fmt.Println("Got empty result")
			}
			timer.Reset(interval)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	downstream(ctx)
}
