package main

import (
	"context" // Package used for managing deadlines, cancellation and request scoped values
	"fmt"     // Package useed for formatted inputs and outputs
	"time"    // Package used for time related operations
)

// DoWork performs a continuous task until the context signals cancellation.
func doWork(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): // Checking if the context is canceled
			fmt.Println(name, "received the cancellation signal:", ctx.Err())
			return // Exits the function when context is canceled
		default:
			fmt.Println(name, "is working..") // Simulating work
			time.Sleep(5 * time.Second)       // Simulating a time consuming task
		}
	}
}

func main() {
	rootCtx := context.Background() // Creating the root context and using for only initialization and it does no work by itself

	// Part 1: By Using cancelable context
	cancelCtx, cancel := context.WithCancel(rootCtx) // Creating a cancelable context that can be manually canceled
	go doWork(cancelCtx, "With Cancel")              // Working with goroutine that performs work using the cancelable context
	time.Sleep(5 * time.Second)                      // Letting the worker run for 5 seconds
	cancel()                                         // Manually sending a cancellation signal
	time.Sleep(2 * time.Second)                      // Giving the worker 2 seconds time to stop

	// Part 2: By Using timeout context
	timeoutCtx, cancelTimeout := context.WithTimeout(rootCtx, 5*time.Second) // Creating a context that will automatically be canceled after 5 seconds
	defer cancelTimeout()                                                    // Ensuring that the resources are cleaned up
	go doWork(timeoutCtx, "With Timeout")                                    // Working with goroutine that performs work by using timeout context
	time.Sleep(2 * time.Second)                                              // Waiting for 2 seconds to observe the worker

	// Part 3: By Using deadline context
	deadline := time.Now().Add(5 * time.Second)                            // Setting a deadline 5 seconds from now
	deadlineCtx, cancelDeadline := context.WithDeadline(rootCtx, deadline) // Creating a context that will cancel at specific given deadline
	defer cancelDeadline()                                                 // Ensuring the cleanup of resources
	go doWork(deadlineCtx, "With Deadlines")                               // Launching the goroutine that performs tasks using deadline context
	time.Sleep(2 * time.Second)                                            // Waiting for 2 seconds time to observe the worker

	// Part 4: By Using context with values
	valueCtx := context.WithValue(rootCtx, "userID", 0310)         // Creating the context to hold a key-value pair
	fmt.Println("User ID from context:", valueCtx.Value("userID")) // Used for retreving the value and printing the value from the context
}
