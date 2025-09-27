# Go `crypto` (curated)

The top-level [`crypto`](https://pkg.go.dev/crypto) package defines common interfaces (`Signer`, `PublicKey`) and hash IDs, while real primitives live in subpackages.

## Common Functions

- Secure RNG: `crypto/rand.Read`
- Hashing: `crypto.SHA256.New()`, `crypto/sha256`
- HMAC: `crypto/hmac` + `sha256.New`
- Sign/verify: `crypto/ed25519` (shows `crypto.Signer` via `PrivateKey`)


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


### HMAC

Package `crypto/hmac` implements the Keyed-Hash Message Authentication Code (HMAC).

```go
func New(h func() hash.Hash, key []byte) hash.Hash
```

The constructor uses different arguments but the write/read operations are similar to the ones for SHA256.

`hmac.Equal` compares two MACs for equality without leaking timing information.


### ED25519

```go
// GenerateKey generates a public/private key pair using entropy from rand.
func GenerateKey(rand io.Reader) (PublicKey, PrivateKey, error)
// Sign signs the given message with priv. rand is ignored and can be nil.
func (priv PrivateKey) Sign(rand io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error)
// VerifyWithOptions reports whether sig is a valid signature of message by publicKey.
func VerifyWithOptions(publicKey PublicKey, message, sig []byte, opts *Options) error
```


---

[Go Back](../../README.md)