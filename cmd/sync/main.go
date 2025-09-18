// Example usage of the sync package in Go.
// Demonstrates sync.WaitGroup and sync.Mutex.

package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func initF() {
	fmt.Println("Init function called - just one time")
}

func main() {
	// --- Example 1: WaitGroup ---
	fmt.Println("---\nExample 1: WaitGroupExample\n---")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(time.Second) // simulate work
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers finished.")

	// --- Example 2: Mutex ---
	fmt.Println("\n---\nExample 2: Mutex\n---")
	var mu sync.Mutex
	counter := 0
	const maxCounter = 1000

	increment := func(workerID int) {
		defer wg.Done()
		contributions := 0
		for i := 0; i < maxCounter; i++ {
			mu.Lock()
			if counter < maxCounter {
				contributions++
				counter++
			}
			mu.Unlock()
			time.Sleep(time.Microsecond) // simulate work. The work is fast enough that the other worker is not able to lock the mutex in time every time.
		}
		fmt.Printf("Worker %d incremented counter %d times. \n", workerID, contributions)
	}

	wg.Add(2)
	go increment(42)
	go increment(666)

	wg.Wait()
	fmt.Printf("Final counter value: %d (expected 2000)\n", counter)

	// --- Example 3: Once ---
	fmt.Println("\n---\nExample 3: Once\n---")
	for i := 0; i < 3; i++ {
		go func() {
			once.Do(initF)
		}()
	}
	time.Sleep(time.Second)

	// --- Example 4: RWMutex ---
	fmt.Println("\n---\nExample 4: RWMutex\n---")
	var rwmu sync.RWMutex
	shared := 0

	wg.Add(3)

	// Reader 1
	go func() {
		defer wg.Done()
		rwmu.RLock()
		fmt.Printf("Reader 1 sees value: %d\n", shared)
		time.Sleep(200 * time.Millisecond)
		rwmu.RUnlock()
		fmt.Println("Reader 1 unlock")
	}()

	// Reader 2
	go func() {
		defer wg.Done()
		rwmu.RLock()
		fmt.Printf("Reader 2 sees value: %d\n", shared)
		time.Sleep(200 * time.Millisecond)
		rwmu.RUnlock()
		fmt.Println("Reader 2 unlock")
	}()

	// Writer
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond) // Let readers start first
		rwmu.Lock()
		shared++
		fmt.Printf("Writer updated value to: %d\n", shared)
		rwmu.Unlock()
		fmt.Println("Writer unlock")
	}()

	wg.Wait()

}
