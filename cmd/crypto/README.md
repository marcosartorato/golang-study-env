# Go `crypto` (curated)

Package [`crypto`](https://pkg.go.dev/crypto)collects common cryptographic constants.

## Common Functions

- Secure RNG: `crypto/rand.Read`
- Hashing: `crypto.SHA256.New()`, `crypto/sha256`


### Secure Random Number Generator

Package `crypto/rand` implements a cryptographically secure random number generator. The functions provided by the sub-package are `Int`, `Prime`, `Read`, and `Text`.

```go
// Int and Prime uses crypto/rand.Reader.
// Reader is a global, shared instance of a cryptographically secure random number generator.
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
func Prime(rand io.Reader, bits int) (*big.Int, error)

// Read fills b with cryptographically secure random bytes.
func Read(b []byte) (n int, err error)

// Return a random string.
func Text() string
```


### SHA256

Package `crypto/sha256` implements the SHA224 and SHA256 hash algorithms. `crypto.SHA256.New()` is the equivalent of running `import crypto/sha256` first followed by `sha256.New()`.

```go
func New() hash.Hash
```

A common use case for `Hash` is:

```go
h := sha256.New()
h.Write([]byte("hello world\n"))    // add a message with Write([]byte{...})
fmt.Printf("%x", h.Sum(nil))        // read the hash with Sum(nil)
```


---

[Go Back](../../README.md)