# Go `encoding/binary` Package (curated)

The [`encoding/binary`](https://pkg.go.dev/encoding/binary) package encodes/decodes **fixed-size values** and **varints**.
You choose endianness via `binary.LittleEndian` or `binary.BigEndian`.

## Common Functions

- Byte order helpers: `LittleEndian`, `BigEndian` with `PutUint16/32/64`, `Uint16/32/64`
- Stream I/O: `binary.Write(w, order, data)`, `binary.Read(r, order, data)`
- Size: `binary.Size(v)` (returns `-1` for variable-size types like slices/strings)
- Varints: `PutUvarint`/`Uvarint`, `PutVarint`/`Varint`

### Handy signatures

```go
// A ByteOrder specifies how to convert byte slices into 16-, 32-, or 64-bit unsigned integers.
type ByteOrder interface {
    Uint16([]byte) uint16; PutUint16([]byte, uint16)
    Uint32([]byte) uint32; PutUint32([]byte, uint32)
    Uint64([]byte) uint64; PutUint64([]byte, uint64)
}
var LittleEndian ByteOrder
var BigEndian ByteOrder

func Write(w io.Writer, order ByteOrder, data any) error
func Read(r io.Reader, order ByteOrder, data any) error
func Size(data any) int

func PutUvarint(buf []byte, x uint64) int
func Uvarint(buf []byte) (uint64, int)
func PutVarint(buf []byte, x int64) int
func Varint(buf []byte) (int64, int)
```

---

[Go Back](../../README.md)