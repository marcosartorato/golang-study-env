# Go `math/rand` Package (curated)

The [`math/rand`](https://pkg.go.dev/math/rand) package provides **pseudo-random** numbers.  
_Not cryptographically secure — use `crypto/rand` for security._

## Common Functions

- Seeding: `rand.New(rand.NewSource(seed))` to get deterministic stream of pseudo-random number.
- Numbers: `rand.Intn(n)`, `rand.Float64()`
- Permutations/shuffle: `rand.Perm(n)`, `rand.Shuffle(n, swap)`

---

[Go Back](../../README.md)