package main

import (
	"fmt"
	"testing"
)

/*
eq is a tiny test helper to compare expected/actual values.
Marking it as a helper makes failure lines point at the caller.
*/
func eq[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}

/*
BenchmarkFib benchmarks the Fib function providing:
- The average time per iteration.
- Memory allocation stats (allocs, bytes allocated).
*/
func BenchmarkFib(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_ = Fib(30)
	}
}

/*
Basic example test.
Check if the output matches the comment.
*/
func ExampleSum() {
	fmt.Printf("%d", Sum([]int{4, 5, 6}))
	// Output: 15
}

/*
This example test is not a good test for the Reverse function but it shows how to
Check if the output matches the comment without accounting for the order.
*/
func ExampleReverse() {
	for _, c := range Reverse("!oG") {
		fmt.Printf("%s\n", string(c))
	}
	// Unordered output: o
	// !
	// G
}

/*
Basic test.
Run a function and check the result.
*/
func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3})
	eq(t, got, 6)
}

// Table-driven test.
func TestIsPalindrome_Table(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"ab", false},
		{"racecar", true},
		{"åäå", true},    // Unicode
		{"こんにちは", false}, // not a palindrome
		{"\xff", false},  // invalid UTF-8
	}

	for _, tc := range cases {
		tc := tc // capture
		t.Run(tc.input, func(t *testing.T) {
			t.Parallel() // subtests can run in parallel
			got := IsPalindrome(tc.input)
			if got != tc.want {
				t.Errorf("IsPalindrome(%q) = %v; want %v", tc.input, got, tc.want)
			}
		})
	}
}

// Sub-tests for grouped scenarios
func TestReverse_Subtests(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		eq(t, Reverse(""), "")
	})
	t.Run("ascii", func(t *testing.T) {
		eq(t, Reverse("Go!"), "!oG")
	})
	t.Run("unicode", func(t *testing.T) {
		eq(t, Reverse("åäå"), "åäå")
	})
}
