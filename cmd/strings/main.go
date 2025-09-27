package main

import (
	"fmt"
	"strings"
)

func main() {
	// Search
	s := "  hello, Go!  "
	fmt.Println("Contains 'Go':", strings.Contains(s, "Go"))
	fmt.Println("HasPrefix '  he':", strings.HasPrefix(s, "  he"))
	fmt.Println("HasSuffix '!  ':", strings.HasSuffix(s, "!  "))
	fmt.Println("Index of ',':", strings.Index(s, ","))

	// Split / Join / Fields / Cut
	fmt.Println("\nSplit:", strings.Split("a,b,,c", ","))
	fmt.Println("Join:", strings.Join([]string{"aa", "bb", "cc"}, "-"))
	fmt.Println("Fields:", strings.Fields(" one   two\tthree\n"))
	before, after, found := strings.Cut("key=value", "=")
	fmt.Printf("Cut: before=%q after=%q found=%v\n", before, after, found)

	// Trim
	fmt.Printf("\nTrimSpace(%q) = %q\n", s, strings.TrimSpace(s))
	fmt.Println("Trim cutset \" ,!\":", strings.Trim(s, " ,!"))

	// Replace / Repeat
	fmt.Println("\nReplaceAll:", strings.ReplaceAll("foo bar foo", "foo", "baz"))
	fmt.Println("Repeat:", strings.Repeat("go", 3))

	// Case & compare
	fmt.Println("\nToUpper:", strings.ToUpper("Gopher"))
	fmt.Println("EqualFold(\"Go\", \"go\"):", strings.EqualFold("Go", "go"))

	// Efficient building
	var b strings.Builder
	b.WriteString("\nhello")
	b.WriteByte(',')
	b.WriteString(" world")
	fmt.Println("Builder:", b.String())
}
