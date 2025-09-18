# Go `sync/atomic` Package

The [`sync/atomic`](https://pkg.go.dev/sync/atomic) package provides low-level atomic memory primitives for safe concurrent programming in Go.  
It allows you to perform atomic operations on variables, such as incrementing counters or swapping values, without using locks.

## Common Functions

- `atomic.Uint64`: An atomic unsigned 64-bit integer type.
    - `.Add(delta)`: Atomically adds `delta` to the value.
    - `.Load()`: Atomically loads (reads) the value.
    - `.Store(val)`: Atomically stores (writes) the value.

---

[Go Back](../../README.md)