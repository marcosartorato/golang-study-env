package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Uint64
	var wg sync.WaitGroup

	numGoroutines, incrementsPerGoroutine := 1000, 5

	wg.Add(numGoroutines)
	counter.Store(0)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Value after %d goroutines running %d unary increment: %d\n", numGoroutines, incrementsPerGoroutine, counter.Load())

	// Demonstrate atomic store overriding the value.
	counter.Store(uint64(42))
	fmt.Printf("Value after override: %d\n", counter.Load())
}
