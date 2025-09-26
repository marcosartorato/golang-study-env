# Go `math` Package (curated)

The [`math`](https://pkg.go.dev/math) package provides floating-point math and constants.

## Common Constants and Functions

- Constants: `math.Pi`, `math.E`
- Roots & powers: `math.Sqrt`, `math.Pow`
- Trig: `math.Sin`, `math.Cos`
- Rounding: `math.Ceil`, `math.Floor`, `math.Round`, `math.Modf`
- Special values: `math.IsNaN`, `math.IsInf`

```go
func Modf(f float64) (int float64, frac float64)
```

`Modf` returns integer and fractional floating-point numbers that sum to *f*. Both values have the same sign as *f*.

`IsNaN` reports whether f is an IEEE 754 “not-a-number” value. `IsInf` reports whether f is an infinity, according to sign.

---

[Go Back](../../README.md)