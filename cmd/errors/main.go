package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

// ErrNotFound is a sentinel error for demo purposes.
var ErrNotFound = errors.New("not found")

// ParseError is a typed error carrying structured info.
type ParseError struct {
	Line int
	Col  int
	Msg  string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error at %d:%d: %s", e.Line, e.Col, e.Msg)
}

// findWidget simulates a lookup that fails with a sentinel error, wrapped for context.
func findWidget(id int) error {
	return fmt.Errorf("find widget %d: %w", id, ErrNotFound)
}

// load simulates a failure parsing the JSON file and wraps underlying errors with context.
func load(path string) error {
	// Wrap a typed error we can extract with errors.As.
	return fmt.Errorf("load %s: %w", path, &ParseError{Line: 3, Col: 14, Msg: "invalid number"})
}

// demoJoin returns a multi-error using errors.Join.
func demoJoin() error {
	e1 := context.Canceled
	e2 := os.ErrPermission
	return errors.Join(e1, e2)
}

func main() {
	// 1) Wrapping + Is with a sentinel error.
	if err := findWidget(42); err != nil {
		fmt.Printf("findWidget error: %s\n", err.Error())
		fmt.Printf("Is not found? %t\n\n", errors.Is(err, ErrNotFound))
	}

	// 2) As to extract a typed error and read its fields.
	if err := load("bad.json"); err != nil {
		var pe *ParseError
		if errors.As(err, &pe) {
			fmt.Printf("As ParseError: line=%d col=%d msg=%q\n\n", pe.Line, pe.Col, pe.Msg)
		}
	}

	// 3) Join multiple errors and query with Is through join and an extra wrapper.
	if err := demoJoin(); err != nil {
		fmt.Printf("joined error: %s\n", err.Error())
		fmt.Printf("Is context.Canceled? %t\n", errors.Is(err, context.Canceled))
		fmt.Printf("Is os.ErrPermission? %t\n", errors.Is(err, os.ErrPermission))

		wrapped := fmt.Errorf("while doing work: %w", err)
		fmt.Printf("Is the \"context.Canceled\" error information retained through wrap+join operations? %t\n", errors.Is(wrapped, context.Canceled))
	}
}
