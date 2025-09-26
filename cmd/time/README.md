# Go `time` Package

The [`time`](https://pkg.go.dev/time) package provides functionality for measuring and displaying time.

## Common Types and Functions

- `time.Time`, `time.Duration`
- Now & math: `time.Now`, `time.Since`, `Time.Add`, `Time.Sub`
- Comparison: `Time.Before`, `Time.After`, and `Time.Equal`
- Format/parse: `Time.Format(layout)`, `time.Parse(layout, s)`, `time.RFC3339`
- Durations: `time.ParseDuration("2h45m")`
- Rounding: `Time.Truncate`, `Time.Round`

A `Time` represents an instant in time with nanosecond precision. A `Duration` represents the elapsed time between two instants as an int64 nanosecond count.

`time.Now` returns the current local time. `time.Since` returns the time elapsed (i.e. a `Duration`) since the `Time` *t* in the argument. It is shorthand for `time.Now().Sub(t)`.

The `Time.Sub` method subtracts two instants, producing a `Duration`. The `Time.Add` method adds a `Time` and a `Duration`, producing a `Time`. `time.Parse(layout, s)`

Time instants can be compared using the `Time.Before`, `Time.After`, and `Time.Equal` methods. 

`Time.Format(layout)` returns a textual representation of the time value formatted according to the *layout* defined by the argument. `time.Parse(layout, s)` parses a formatted string according to the *layout* defined by the argument and returns the time value it represents. `time.RFC3339` is the string describing the RFC3339 format. This format is to be preferred to others being a well established standard.

`ParseDuration` parses a duration string. A duration string is a possibly signed sequence of decimal numbers, each with a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

```go
func (t Time) Truncate(d Duration) Time
```

`Truncate` returns the result of rounding *t* down to a multiple of *d*.

```go
func (t Time) Round(d Duration) Time
```

`Round` returns the result of rounding *t* to the nearest multiple of *d*. 

---

[Go Back](../../README.md)
