# Go `testing` Package

The [`testing`](https://pkg.go.dev/testing) package provides support for automated testing. It's used in concert with the `go test` command, which automates execution of any function with signature `func TestXxx(*testing.T)`, where `Xxx` does not start with a lowercase letter.

The test file can be in the same package as the one being tested, or in a corresponding package with the suffix `_test`.

It provides primitives for **unit tests**, **benchmarks**, **examples**, and **fuzz tests**.

## Common Types and Functions

- `type T`: Test state and helpers, e.g. `t.Error`, `t.Fatalf`, `t.Helper`, `t.Run`, `t.Parallel`.
- `type B`: Benchmark state (`b.Loop` loop, `b.ReportAllocs`).
- `type F`: Fuzzing entry (`f.Add`, `f.Fuzz`).
- `go test`: Discovers `*_test.go` files and runs `Test*`, `Benchmark*`, `Example*`, and `Fuzz*`.

About `type T`:

- `t.Log`/`t.Logf` formats its argument using default/specified formatting, like `fmt.Println`/`fmt.Printf`, and records the text in the error log.
- `t.Fail`/`t.FailNow` marks the function as having failed but continues/stop its execution. 
- `t.Error` is equivalent to `Log` followed by `Fail`.
- `t.Fatalf` is equivalent to `Logf` followed by `FailNow`.
- `t.Helper` marks the calling function as a test helper function. When printing file and line information, that function will be skipped. Helper may be called simultaneously from multiple goroutines.
- `t.Run(name string, f func(t *T))` runs *f* as a sub-test of *t* called *name*. It runs *f* in a separate goroutine and blocks until *f* returns or calls `t.Parallel` to become a parallel test.
- `t.Parallel` signals that this test is to be run in parallel with (and only with) other parallel tests.


### Benchmarks

Benchmarks are functions of the form `func BenchmarkXxx(*testing.B)`. They are executed by `go test` when the `-bench` flag is provided. Benchmarks are run sequentially.

A sample benchmark function looks like this:

```go
func BenchmarkSomething(b *testing.B) {
    for b.Loop() {
        ... // Body of the loop
    }
}
```

Only the body of the loop is timed, so benchmarks should do expensive setup before calling `b.Loop`. The output:

```
Benchmark<name>     <number of iterations>      <avg speed of the body of the loop> ns/op
```

### Examples

The package also runs and verifies *example* code. *Example* functions may include a concluding line comment that begins with `Output:` and is compared with the standard output of the function when the tests are run. The comparison ignores leading and trailing space.

These are examples of an example:

```go
func ExampleHello() {
    fmt.Println("hello")
    // Output: hello
}

func ExampleSalutations() {
    fmt.Println("hello, and")
    fmt.Println("goodbye")
    // Output:
    // hello, and
    // goodbye
}

func ExamplePerm() {
    for _, value := range Perm(5) {
        fmt.Println(value)
    }
    // Unordered output: 4
    // 2
    // 1
    // 3
    // 0
}
```

The comment prefix `Unordered output:` is like `Output:`, but matches any line order. Example functions without output comments are compiled but not executed.

### Fuzzing

The package supports *fuzzing*, a testing technique where a function is called with randomly generated inputs to find bugs not anticipated by unit tests.

Fuzz tests come in the form of `func FuzzXxx(*testing.F)`.

### Sub-tests and sub-benchmarks

The `T.Run` and `B.Run` methods allow defining sub-tests and sub-benchmarks enabling table-driven benchmarks, hierarchical tests, etc. It also provides a way to share common setup and tear-down code:

```go
func TestFoo(t *testing.T) {
    // <setup code>
    t.Run("A=1", func(t *testing.T) { ... })
    t.Run("A=2", func(t *testing.T) { ... })
    t.Run("B=1", func(t *testing.T) { ... })
    // <tear-down code>
}
```

## How to Run

- All tests: `go test ./cmd/testing`
- Filter tests: `go test -run TestIsPalindrome ./cmd/testing`
- Benchmarks: `go test -bench . -benchmem ./cmd/testing`
- Examples (show output): `go test -run Example -v ./cmd/testing`
- Coverage: `go test -cover ./cmd/testing`
- Fuzzing (Go ≥1.18): `go test -fuzz=Fuzz -fuzztime=2s ./cmd/testing`

The argument to the `-run`, `-bench`, and `-fuzz` flags is an unanchored regular expression that matches the test's name. An empty expression matches any string. For example:

```
go test -run ''        # Run all tests.
go test -run Foo       # Run top-level tests matching "Foo", such as "TestFooBar".
go test -run Foo/A=    # For top-level tests matching "Foo", run subtests matching "A=".
go test -run /A=1      # For all top-level tests, run subtests matching "A=1".
go test -fuzz FuzzFoo  # Fuzz the target matching "FuzzFoo"
```

---

[Go Back](../../README.md)
