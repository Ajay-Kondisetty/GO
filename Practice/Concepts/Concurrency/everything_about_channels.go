package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Approach 1 for sending signals from goroutine to main
	completeSignal := make(chan struct{})
	go func(n int, ch chan<- struct{}) {
		for i := 0; i < n; i++ {
			fmt.Printf("Iteration:- %d\n", i+1)
			time.Sleep(1 * time.Second)
		}

		close(ch)
	}(10, completeSignal)
	select {
	case <-completeSignal:
		fmt.Println("completed signal received.")
	}

	// Approach 2 for sending signals from goroutine to main
	ctx, cancel := context.WithCancel(context.Background())
	go func(n int, cancel context.CancelFunc) {
		defer cancel()
		for i := 0; i < n; i++ {
			fmt.Println(fmt.Sprintf("Iteration:- %d", i+1))
			time.Sleep(1 * time.Second)
		}
	}(10, cancel)
	select {
	case <-ctx.Done():
		fmt.Println("completed signal received.")
	}

	// Approach 3 for sending signals from goroutine to main
	completeChannel := make(chan int)
	go func(n int, ch chan<- int) {
		for i := 0; i < n; i++ {
			fmt.Printf("Iteration:- %d\n", i+1)
			time.Sleep(1 * time.Second)
		}

		close(ch)
	}(10, completeChannel)
	<-completeChannel
	fmt.Println("completed signal received")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	fmt.Println("here waiting for interrupt or terminate signal....")
	s := <-ch
	fmt.Println("received signal:", s)
}
