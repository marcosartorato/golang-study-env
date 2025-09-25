package main

import (
	"fmt"
	"unicode/utf8"
)

// Sum returns the sum of a slice of ints.
func Sum(xs []int) int {
	total := 0
	for _, v := range xs {
		total += v
	}
	return total
}

// Reverse returns s reversed by runes (handles Unicode).
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// IsPalindrome reports whether s reads the same forwards and backwards by runes.
func IsPalindrome(s string) bool {
	// Early exit for invalid UTF-8; tests will exercise this path too.
	if !utf8.ValidString(s) {
		return false
	}
	return s == Reverse(s)
}

// Fib returns the n-th Fibonacci number (iterative, avoids recursion overhead).
func Fib(n int) int {
	if n < 2 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println("Sum([]int{1,2,3}) =", Sum([]int{1, 2, 3}))
	fmt.Println(`Reverse("Go!") =`, Reverse("Go!"))
	fmt.Println(`IsPalindrome("racecar") =`, IsPalindrome("racecar"))
	fmt.Println("Fib(10) =", Fib(10))
}
