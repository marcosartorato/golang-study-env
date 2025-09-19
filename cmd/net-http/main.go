package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// CreateServer sets up a simple HTTP server with one endpoint.
// Note that the use of localhost and a fixed port is for simplicity.
func CreateServer() (*http.Server, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, "Hello, World!"); err != nil {
			http.Error(w, "failed to write response", http.StatusInternalServerError)
			return
		}
	})
	return &http.Server{Addr: "localhost:8888", Handler: mux}, nil
}

func main() {
	/*
		Create the server structure and start it in a goroutine.
		Also ensure it gets properly shutdown at the end.
	*/
	srv, err := CreateServer()
	if err != nil {
		fmt.Printf("failed to create server: %v", err)
		os.Exit(1)
		return
	}
	defer func() {
		if err := srv.Shutdown(context.TODO()); err != nil && err != http.ErrServerClosed {
			fmt.Printf("failed to shutdown server: %v", err)
		}
	}()
	go func() {
		fmt.Println("Server listening on http://localhost:8888")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("failed listening/serving by the server: %v", err)
		}
	}()

	// Give the server a moment to start
	time.Sleep(200 * time.Millisecond)

	/*
		Run 20 clients concurrently, each sending a request to the server
		and printing the response.
	*/
	const numOfClients = 20
	var wg sync.WaitGroup
	wg.Add(numOfClients)
	for i := 0; i < numOfClients; i++ {
		go func(id int) {
			resp, err := http.Get("http://localhost:8888/hello")
			if err != nil {
				panic(err)
			}
			defer func() {
				if err = resp.Body.Close(); err != nil {
					panic(err)
				}
			}()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Client %d gets a response from HTTP server: %s", id, string(body))
			wg.Done()
		}(i)
	}

	// Wait for all client goroutines to finish before exiting.
	wg.Wait()
}
