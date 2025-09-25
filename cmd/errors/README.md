# Go `errors` Package

The [`errors`](https://pkg.go.dev/errors) package provides primitives to create and work with errors, including **wrapping**, **type assertions**, and **multi-error joins**.

## Common Functions and Patterns

- `errors.New("msg")`: Creates an `error` whose only content is a text message.
- `fmt.Errorf("context: %w", err)`: Wrap an error.
- `errors.Is(err, target)`: Reports whether `err` or any wrapped error matches `target` (by sentinel equality, e.g. `io.EOF`, or `Is` method).
- `errors.As(err, &target)`: Assigns the first matching error in the chain to `target` (by type or `As` method). `target` must be a pointer.
- `errors.Unwrap(err)`: Returns the next error in the chain (one level).
- `errors.Join(errs...)`: Combines multiple errors; `Is`/`As` traverse all members. `nil` values are ignored; if all are `nil`, result is `nil`.

An error *e* wraps another error if *e*'s type has one of the methods:

```go
Unwrap() error      // single wrapped error
Unwrap() []error    // multiple wrapped errors (introduced with Go 1.20)
```

and `e.Unwrap()` returns a non-nil error *w* or a slice containing *w*.
Before Go 1.20, `fmt.Errorf` allowed only one `%w` in the format string.  
Since Go 1.20, multiple `%w` verbs are supported, and the resulting error implements `Unwrap() []error`.

Both `Is` and `As` recursively call `e.Unwrap()` to walk the chain.


---

[Go Back](../../README.md)