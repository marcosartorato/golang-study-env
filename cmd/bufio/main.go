package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// --- Reader: ReadString + Peek (doesn't advance) ---
	fmt.Println("== bufio.Reader ==")
	src := strings.NewReader("alpha\nbeta\ngamma") // Get the io.Reader out of a string
	r := bufio.NewReader(src)                      // Wrap it in a bufio.Reader

	line1, _ := r.ReadString('\n')   // Read until '\n'
	fmt.Printf("line1: %q\n", line1) // "alpha\n"

	peek, _ := r.Peek(4)                   // look ahead without consuming
	fmt.Printf("peek: %q\n", string(peek)) // "beta"
	fmt.Println("The peeked data is still there. Let's read again...")

	line2, _ := r.ReadString('\n')   // Read again until '\n'
	fmt.Printf("line2: %q\n", line2) // "beta\n"

	// --- Writer: buffer then Flush ---
	fmt.Println("\n== bufio.Writer ==")
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	_, _ = w.WriteString("hello, ")
	_, _ = w.WriteString("world")
	fmt.Println("buffer before flush:", b.String())
	_ = w.Flush()
	fmt.Println("buffer after flush:", b.String())
	fmt.Println("Don't forget to flush!")

	// --- Scanner: split into words ---
	fmt.Println("\n== bufio.Scanner (ScanWords) ==")
	text := "one two  three"
	sc := bufio.NewScanner(strings.NewReader(text))
	sc.Split(bufio.ScanWords)
	var words []string
	for sc.Scan() {
		words = append(words, sc.Text())
	}
	fmt.Println("words:", words)
}
