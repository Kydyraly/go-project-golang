package main

import (
	"fmt"
	"sync"
	"time"
)
type RateLimiter struct {
	tokens chan struct{} 
	stop   chan struct{} 
	wg     sync.WaitGroup
}
func NewRateLimiter(n int) *RateLimiter {
	if n <= 0 {
		panic("n must be > 0")
	}
	rl := &RateLimiter{
		tokens: make(chan struct{}, n),
		stop:   make(chan struct{}),
	}
	for i := 0; i < n; i++ {
		rl.tokens <- struct{}{}
	}
	rl.wg.Add(1)
	go func() {
		defer rl.wg.Done()
		ticker := time.NewTicker(time.Second) 
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				for i := 0; i < n; i++ {
					select {
					case rl.tokens <- struct{}{}:
					default:
					}
				}
			case <-rl.stop:
				return
			}
		}
	}()

	return rl
}
func (rl *RateLimiter) AllowTry() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}
func (rl *RateLimiter) Wait() {
	<-rl.tokens
}
func (rl *RateLimiter) Stop() {
	close(rl.stop)
	rl.wg.Wait()
}
func main() {
	rl := NewRateLimiter(5)
	defer rl.Stop()

	var wg sync.WaitGroup
	workers := 10
	requestsPerWorker := 5

	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < requestsPerWorker; i++ {
				rl.Wait()

				fmt.Printf("worker %d processed request %d at %s\n", id, i, time.Now().Format("15:04:05.000"))
			}
		}(w)
	}

	wg.Wait()
}
