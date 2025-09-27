# Go `bytes` Package (curated)

The [`bytes`](https://pkg.go.dev/bytes) package provides helpers for working with `[]byte`.  
Like `strings`, but byte-oriented.

## Common Functions

- Search/compare: `Contains`, `HasPrefix`, `HasSuffix`, `Index`, `Equal`, `Compare`
- Split/join/cut: `Split`, `Join`, `Fields`, `Cut`
- Trim/replace: `TrimSpace`, `Trim`, `ReplaceAll`
- Builders/readers: `bytes.Buffer`, `bytes.NewReader`

### Handy signatures

```go
func Contains(b, subslice []byte) bool
func HasPrefix(b, prefix []byte) bool
func HasSuffix(b, suffix []byte) bool
func Index(s, sep []byte) int
func Equal(a, b []byte) bool
func Compare(a, b []byte) int

func Split(s, sep []byte) [][]byte
func Join(s [][]byte, sep []byte) []byte
func Fields(s []byte) [][]byte
func Cut(s, sep []byte) (before, after []byte, found bool)

func TrimSpace(s []byte) []byte
func Trim(s []byte, cutset string) []byte

type Buffer struct { /* ... */ }
func (b *Buffer) Write(p []byte) (int, error)
func (b *Buffer) WriteString(s string) (int, error)
func (b *Buffer) WriteByte(c byte) error
func (b *Buffer) String() string
func (b *Buffer) Bytes() []byte
func (b *Buffer) Grow(n int)

func NewReader(b []byte) *Reader
// Reader implements io.Reader, io.ReaderAt, io.Seeker, io.WriterTo.
```

---

[Go Back](../../README.md)