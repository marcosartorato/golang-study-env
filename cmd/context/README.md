# Go `context` Package

The [`context`](https://pkg.go.dev/context) package provides context propagation, cancellation, and deadlines for processes and goroutines in Go.  
It is commonly used to control the lifetime of operations, especially in concurrent or networked programs.

## Common Types and Functions

- `context.Context`: The structure carries deadlines, cancellation signals, and other request-scoped values.
- `context.Background()`: Returns a non-nil, empty root context.
- `context.WithCancel(parent)`: Returns a context that points to the parent one but has a new `Done` channel. This `Done` channel is closed when the returned `cancel` function is called or when the parent context's `Done` channel is closed, whichever happens first.
- `context.WithTimeout(parent, timeout)`: Returns a copy of the parent context with a deadline. So code should call `cancel` as soon as the operations running in this `Context` complete.
- `context.WithValue(parent, key, val)`: Returns a copy of the parent context. In the derived context, the value associated with key is val.

---

[Go Back](../../README.md)