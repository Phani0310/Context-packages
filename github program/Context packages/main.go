package main

import (
	"context" // Package for managing deadlines, cancellation, and request-scoped values.
	"fmt"     // Package used for formatted input and output.
	"time"    // Package used for time-related operations.
)

// Task represents a task that performs operations with a given context.
type Task struct {
	Name string
}

// NewTask initializes and returns a new Task instance.
func NewTask(name string) *Task {
	return &Task{Name: name}
}

// DoWork performs a continuous task until the context signals cancellation.
// It prints a message every 5 seconds to simulate work.
func (t *Task) DoWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Checks if the context is canceled.
			fmt.Println(t.Name, "received cancellation signal:", ctx.Err())
			return // Exits the function when context is canceled.
		default:
			fmt.Println(t.Name, "is working...") // Simulating work.
			time.Sleep(5 * time.Second)          // Simulating a long-running task.
		}
	}
}

func main() {
	rootCtx := context.Background()

	// Part 1: By Using a cancelable context.
	cancelCtx, cancel := context.WithCancel(rootCtx) // Creates a cancelable context.
	task1 := NewTask("With Cancel")
	go task1.DoWork(cancelCtx)  // Runs the task in a goroutine.
	time.Sleep(5 * time.Second) // Lets the task run for 5 seconds.
	cancel()                    // Sends a cancellation signal.
	time.Sleep(2 * time.Second) // Allows time for cleanup.

	// Part 2: By Using a timeout context.
	timeoutCtx, cancelTimeout := context.WithTimeout(rootCtx, 5*time.Second) // Creates a context that cancels after 5 seconds.
	defer cancelTimeout()                                                    // Ensures resource cleanup.
	task2 := NewTask("With Timeout")
	go task2.DoWork(timeoutCtx) // Runs the task with a timeout context.
	time.Sleep(2 * time.Second) // Observes the task for 2 seconds.

	// Part 3: By Using a deadline context.
	deadline := time.Now().Add(5 * time.Second)                            // Sets a deadline 5 seconds from now.
	deadlineCtx, cancelDeadline := context.WithDeadline(rootCtx, deadline) // Creates a context that cancels at a specific deadline.
	defer cancelDeadline()                                                 // Ensures cleanup.
	task3 := NewTask("With Deadline")
	go task3.DoWork(deadlineCtx) // Runs the task with a deadline context.
	time.Sleep(2 * time.Second)  // Observes the task for 2 seconds.

	// Part 4: By Using context with values.
	valueCtx := context.WithValue(rootCtx, "userID", 310)          // Creates a context to hold a key-value pair.
	fmt.Println("User ID from context:", valueCtx.Value("userID")) // Retrieves and prints the value from context.
}
