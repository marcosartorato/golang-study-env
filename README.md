# golang-study-env

This repository contains small, self-contained examples showing how to use different Go packages.  
Each package has its own folder with a `main.go` file that demonstrates key functionality.

## Package Examples

- [bufio](./cmd/bufio/README.md): Buffered I/O.
- [context](./cmd/context/README.md): Deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
- [crypto](./cmd/crypto/README.md): Common cryptographic constants and functions.
- [encoding/json](./cmd/encoding-json/README.md): Functions for encoding (marshaling) Go values to JSON and decoding (unmarshaling) JSON data into Go values.
- [errors](./cmd/errors/README.md): Functions to manipulate errors.
- [math](./cmd/math/README.md): Basic constants and mathematical functions.
- [math/rand](./cmd/math-rand/README.md): Pseudorandom numbers. It should not be used for security-sensitive work.
- [net](./cmd/net/README.md): A portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.
- [net/http](./cmd/net-http/README.md): HTTP client and server implementation.
- [os](./cmd/os/README.md): Platform-independent interface to operating system functionality. 
- [reflect](./cmd/reflect/README.md): Implementation of run-time reflection, allowing a program to manipulate objects with arbitrary types.
- [strings](./cmd/strings/README.md): Simple functions to manipulate UTF-8 encoded strings.
- [sync](./cmd/sync/README.md): Synchronization primitives for concurrent programming in Go.
- [sync/atomic](./cmd/sync-atomic/README.md): Low-level atomic memory primitives for safe concurrent programming.
- [testing](./cmd/testing/README.md): Support for automated testing of Go packages. It's intended to be used in concert with the `go test` command.
- [time](./cmd/time/README.md): Functionalities for measuring and displaying time.

## Cites

Lots of the notes and codes is copied or heavily inspired by:

- The [official documentation](https://pkg.go.dev/std).
- [The Go Programming Language](https://www.gopl.io/).
