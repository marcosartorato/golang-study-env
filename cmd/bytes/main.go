package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// --- Search / compare ---
	s := []byte("  hello, Go!  ")
	fmt.Println("Contains 'Go':", bytes.Contains(s, []byte("Go")))
	fmt.Println("HasPrefix '  he':", bytes.HasPrefix(s, []byte("  he")))
	fmt.Println("HasSuffix '!  ':", bytes.HasSuffix(s, []byte("!  ")))
	fmt.Println("Index of ',':", bytes.Index(s, []byte(",")))
	fmt.Println("Equal('abc','abc'):", bytes.Equal([]byte("abc"), []byte("abc")))
	fmt.Println("Compare('a','b'):", bytes.Compare([]byte("a"), []byte("b"))) // -1,0,1

	// --- Split / Join / Fields / Cut ---
	fmt.Println("\nSplit:", bytes.Split([]byte("a,b,,c"), []byte(",")))
	fmt.Println("Join:", string(bytes.Join([][]byte{[]byte("aa"), []byte("bb"), []byte("cc")}, []byte("-"))))
	fmt.Println("Fields:", bytes.Fields([]byte(" one   two\tthree\n")))
	before, after, found := bytes.Cut([]byte("key=value"), []byte("="))
	fmt.Printf("Cut: before=%q after=%q found=%v\n", before, after, found)

	// --- Trim / ReplaceAll ---
	fmt.Printf("\nTrimSpace(%q) = %q\n", s, bytes.TrimSpace(s))
	fmt.Println("Trim cutset \" ,!\":", string(bytes.Trim(s, " ,!")))
	fmt.Println("ReplaceAll:", string(bytes.ReplaceAll([]byte("foo bar foo"), []byte("foo"), []byte("baz"))))

	// --- Efficient building: bytes.Buffer ---
	var b bytes.Buffer
	b.Grow(16)
	_, _ = b.WriteString("hello")
	_ = b.WriteByte(',')
	_, _ = b.Write([]byte(" world"))
	fmt.Println("\nBuffer.String():", b.String())

	// --- Reading with bytes.Reader (seek/read/read-at) ---
	r := bytes.NewReader([]byte("alpha beta"))
	buf := make([]byte, 5)
	_, _ = r.Read(buf) // reads "alpha"
	fmt.Printf("Read: %q\n", buf)

	_, _ = r.Seek(6, io.SeekStart) // point at 'b' in "beta"
	_, _ = r.Read(buf[:4])
	fmt.Printf("Seek+Read: %q\n", buf[:4])

	buf2 := make([]byte, 4)
	_, _ = r.ReadAt(buf2, 6) // read "beta" at offset 6
	fmt.Printf("ReadAt@6: %q\n", buf2)
}
