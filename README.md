# golang-study-env

This repository contains small, self-contained examples showing how to use different Go packages.  
Each package has its own folder with a `main.go` file that demonstrates key functionality.

## Package Examples

- [sync](./cmd/sync/README.md): Synchronization primitives for concurrent programming in Go.
- [sync/atomic](./cmd/sync-atomic/README.md): Low-level atomic memory primitives for safe concurrent programming.
- [context](./cmd/context/README.md): Deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
- [net/http](./cmd/net-http/README.md): HTTP client and server implementation.
- [encoding/json](./cmd/encoding-json/README.md): Functions for encoding (marshaling) Go values to JSON and decoding (unmarshaling) JSON data into Go values.
- [reflect](./cmd/reflect/README.md): Implements run-time reflection, allowing a program to manipulate objects with arbitrary types.
- [testing](./cmd/testing/README.md): Support for automated testing of Go packages. It's intended to be used in concert with the `go test` command.

## Cites

Lots of the notes and codes is copied or heavily inspired by:

- The [official documentation](https://pkg.go.dev/std).
- [The Go Programming Language](https://www.gopl.io/).
