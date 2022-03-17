// Task 05
// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	breakWithTimeout(3)
}

func breakWithContext(seconds int) {
	if seconds < 1 {
		panic("seconds parameter must be greater than zero")
	}

	// Create context with Timeout.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(seconds))
	defer cancel()

	// Create a channel for messages.
	messages := make(chan string)

	// Start goroutine to publish messages.
	go func() {
		for i := 0; ; i++ {
			messages <- fmt.Sprintf("Message: %v", i)
		}
	}()

Loop:
	for {
		select {
		case msg := <-messages: // Got a message? Print it!
			fmt.Println(msg)
		case <-ctx.Done(): // call for the cancel function or timeout end? Here we are!
			break Loop
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Practically the same. Here we make handmade timeout.
func breakWithTimeout(seconds int) {
	if seconds < 1 {
		panic("seconds parameter must be greater than zero")
	}

	messages := make(chan string)

	go func() {
		for i := 0; ; i++ {
			messages <- fmt.Sprintf("Message: %v", i)
		}
	}()

	// For loop until time.Since() is less than defined.
	for start := time.Now(); time.Since(start) < time.Second*time.Duration(seconds); {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
