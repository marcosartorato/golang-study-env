# Go `strings` Package (curated)

The [`strings`](https://pkg.go.dev/strings) package provides helpers for working with UTF-8 text.

## Common Functions

- Search: `Contains`, `HasPrefix`, `HasSuffix`, `Index`
- Split/join: `Split`, `Join`, `Fields`, `Cut`
- Trim: `TrimSpace`, `Trim`
- Replace/repeat: `ReplaceAll`, `Repeat`
- Case & compare: `ToLower`, `ToUpper`, `EqualFold`
- Efficient building: `strings.Builder` (`WriteString`, `WriteByte`, `String`)

### Handy signatures

```go
func Contains(s, substr string) bool
func HasPrefix(s, prefix string) bool
func HasSuffix(s, suffix string) bool
func Index(s, substr string) int

func Split(s, sep string) []string
func Join(elems []string, sep string) string
func Fields(s string) []string
func Cut(s, sep string) (before, after string, found bool)

func TrimSpace(s string) string
func Trim(s, cutset string) string

func ReplaceAll(s, old, new string) string
func Repeat(s string, count int) string

func ToLower(s string) string
func ToUpper(s string) string
func EqualFold(s, t string) bool
```

---

[Go Back](../../README.md)