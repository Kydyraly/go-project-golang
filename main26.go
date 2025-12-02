package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int, tasks <-chan int) {
	for {
		select {
		case t := <-tasks:
			fmt.Printf("Worker %d processed task %d\n", id, t)
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	tasks := make(chan int)

	for i := 1; i <= 3; i++ {
		go worker(ctx, i, tasks)
	}
	for i := 1; i <= 5; i++ {
		tasks <- i
		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("STOP signal...")
	cancel()

	time.Sleep(1 * time.Second)
}
