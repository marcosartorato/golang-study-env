package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Args & environment
	fmt.Println("== args & env ==")
	fmt.Println("Args:", os.Args)
	if v, ok := os.LookupEnv("DEMO_OS_VAR"); ok {
		fmt.Println("DEMO_OS_VAR (existing):", v)
	}
	_ = os.Setenv("DEMO_OS_VAR", "hello")
	fmt.Println("DEMO_OS_VAR (after Setenv):", os.Getenv("DEMO_OS_VAR"))

	// Working directory & list a few entries
	fmt.Println("\n== cwd & list ==")
	if wd, err := os.Getwd(); err == nil {
		fmt.Println("Getwd:", wd)
	}
	if entries, err := os.ReadDir("."); err == nil {
		for _, e := range entries {
			name := e.Name()
			info, err := e.Info()
			if err != nil {
				fmt.Println("info error:", err)
				continue
			}
			fmt.Printf("%s %s\n", info.Mode(), name)
		}
	}

	// Basic file ops in a temp dir
	fmt.Println("\n== files (Write/Read/Stat/Rename) ==")
	tmp, _ := os.MkdirTemp("", "osdemo-*")
	defer func() {
		if err := os.RemoveAll(tmp); err != nil {
			fmt.Println("RemoveAll error:", err)
		}
	}()

	p := filepath.Join(tmp, "hello.txt")
	_ = os.WriteFile(p, []byte("hi\n"), 0o644)

	if b, err := os.ReadFile(p); err == nil {
		if fi, err := os.Stat(p); err == nil {
			fmt.Printf("read %d bytes; mode=%v\n", len(b), fi.Mode())
		}
	}

	newp := filepath.Join(tmp, "renamed.txt")
	_ = os.Rename(p, newp)
	fmt.Println("renamed to:", filepath.Base(newp))
}
