# Go `os` Package (curated)

The [`os`](https://pkg.go.dev/os) package provides platform-independent OS primitives.

## Covered Functions

- Args & env: `os.Args`, `os.Getenv`, `os.LookupEnv`, `os.Setenv`
- Working dir & listing: `os.Getwd`, `os.ReadDir`
- Files: `os.MkdirTemp`, `os.WriteFile`, `os.ReadFile`, `os.Rename`, `os.RemoveAll`, `os.Stat`

```go
// The command-line arguments, starting with the program name.
var Args []string

// To distinguish between an empty value and an unset value, use LookupEnv.
func Getenv(key string) string
func LookupEnv(key string) (string, bool)

func Setenv(key, value string) error
func Getwd() (dir string, err error)

// Reads the named directory, returning all its directory entries sorted by filename.
func ReadDir(name string) ([]DirEntry, error)

func MkdirTemp(dir, pattern string) (string, error)

// WriteFile writes data to the named file, creating it if necessary.
// If the file does not exist, WriteFile creates it with permissions perm.
func WriteFile(name string, data []byte, perm FileMode) error

func ReadFile(name string) ([]byte, error)
func Rename(oldpath, newpath string) error
func RemoveAll(path string) error
func Stat(name string) (FileInfo, error)
```


---

[Go Back](../../README.md)