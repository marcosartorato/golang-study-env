# Go `reflect` Package

The [`reflect`](https://pkg.go.dev/reflect) package provides runtime reflection, allowing a program to inspect and manipulate objects with arbitrary types.  
It is commonly used for generic programming, serialization, and building libraries that need to work with types dynamically.

## Common Types and Functions

- `reflect.TypeOf(v)`: Returns the reflection `Type` that represents the dynamic type of `v`.
- `reflect.ValueOf(v)`: Returns a reflection `Value` representing `v`.
- `(v Value) Kind()`: Returns the `v`'s `Kind`.

---

[Go Back](../../README.md)