package main

import (
	"context"
	"fmt"
	"time"
)

// doWork is a worker function that continuously performs work and stops when it receives a cancellation signal.
func doWork(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): // checking if the context is canceled
			fmt.Println(name, "received the cancellation signal:", ctx.Err())
			return // exit the function when the context is canceled
		default:
			fmt.Println(name, "is working..") // simulating work
			time.Sleep(2 * time.Second)       // simulating a time-consuming task
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // creating a cancelable context
	go doWork(ctx, "with cancel")                           // start the with cancel in a goroutine

	time.Sleep(5 * time.Second) // letting the with cancel run for 5 seconds
	cancel()
	time.Sleep(2 * time.Second) // giving some time for the with cancel to exit
}
