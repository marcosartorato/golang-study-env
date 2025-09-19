package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	rootCtx := context.Background()
	goroutineProcess := func(
		ctx context.Context, ch chan struct{}, callBackDone, callBackCh func(),
	) {
		select {
		case <-ctx.Done():
			callBackDone()
		case <-ch:
			callBackCh()
		}
	}

	/*
		Example 1: Context with Timeout

		Always release the resources created by "context.WithTimeout", even if the
		timeout.
		To release the resources call the "cancel" function returned by
		context.WithTimeout.
		To be sure it’s always called, use "defer cancel()".
		In this example we want to call it explicitly when the ctx trigger the timeout,
		so we don’t use "defer".
	*/
	fmt.Println("---\nExample 1: Context with timeout and timeout triggered\n---")
	ctx, cancel := context.WithTimeout(rootCtx, 500*time.Millisecond)
	wg.Add(1)

	done := make(chan struct{})
	go goroutineProcess(
		ctx,
		done,
		func() {
			// Here the ctx timeout is triggered and the ctx is cancelled implicitly by the timeout.
			fmt.Println("Operation cancelled or timed out:", ctx.Err())
			cancel() // Release resources
			wg.Done()
		}, func() {
			fmt.Println("Operation completed")
			wg.Done()
		},
	)

	// Simulate work
	time.Sleep(600 * time.Millisecond) // Wait the time needed to trigger the timeout
	wg.Wait()                          // Give goroutine time to print

	/*
		Example 2: Context with Timeout

		Always release the resources created by "context.WithTimeout", even if the
		timeout.
		To release the resources call the "cancel" function returned by
		context.WithTimeout.
		To be sure it’s always called, use "defer cancel()".
		In this example we want to call it explicitly before the timeout (to show
		the difference), so we don’t use "defer".
	*/
	fmt.Println("\n---\nExample 2: Context with timeout and cancel triggered\n---")
	ctx, cancel = context.WithTimeout(rootCtx, 500*time.Millisecond)
	wg.Add(1)

	go goroutineProcess(
		ctx,
		done,
		func() {
			fmt.Println("Operation cancelled or timed out:", ctx.Err())
			wg.Done()
		}, func() {
			fmt.Println("Operation completed")
			wg.Done()
		},
	)

	// Simulate work
	time.Sleep(200 * time.Millisecond)
	cancel()
	wg.Wait()   // Give goroutine time to print
	close(done) // Clean up the channel

	/*
		Example 3: Save a value in the context

		Use context.WithValue to pass request-scoped values, such as
		authorization tokens, user IDs, and other metadata, across API boundaries
		and between processes.
	*/
	fmt.Println("\n---\nExample 3: Context with value\n---")
	f := func(ctx context.Context, key string) {
		if v := ctx.Value(key); v != nil {
			fmt.Printf("Found value for key '%s': %v\n", key, v)
		} else {
			fmt.Printf("No value found for key '%s'\n", key)
		}
	}
	type ctxKey string
	ctx = context.WithValue(rootCtx, ctxKey("key"), "value")
	f(ctx, "key")    // Should find the value
	f(ctx, "no-key") // Should not find the value
}
