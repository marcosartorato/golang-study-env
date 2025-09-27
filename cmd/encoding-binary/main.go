package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	// --- Fixed-width numbers with explicit endianness ---
	fmt.Println("== fixed-width numbers ==")
	var v uint32 = 0xDEADBEEF
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	fmt.Printf("LE bytes: % X\n", b)
	fmt.Printf("LE -> uint32: 0x%X\n", binary.LittleEndian.Uint32(b))

	// --- Stream write/read (struct & slice) ---
	fmt.Println("\n== binary.Write / binary.Read ==")
	type Header struct {
		Magic   uint32
		Version uint16
		Flags   uint16
	}
	hw := Header{Magic: 0xABCD1234, Version: 3, Flags: 0x01}

	var buf bytes.Buffer
	// write struct then a slice of uint16
	_ = binary.Write(&buf, binary.BigEndian, hw)
	payload := []uint16{10, 20, 30}
	_ = binary.Write(&buf, binary.BigEndian, payload)

	// read them back
	var hr Header
	_ = binary.Read(&buf, binary.BigEndian, &hr)
	dst := make([]uint16, len(payload))
	_ = binary.Read(&buf, binary.BigEndian, dst)
	fmt.Printf("Header round-trip: %+v\n", hr)
	fmt.Printf("Slice round-trip:  %v\n", dst)
	fmt.Println("binary.Size(Header):", binary.Size(hw)) // fixed-size → bytes

	// --- Varints (compact integers) ---
	fmt.Println("\n== varints ==")
	var u uint64 = 300 // encodes in 2 bytes (0xAC 0x02)
	vbuf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(vbuf, u)
	fmt.Printf("PutUvarint(300) -> % X (len=%d)\n", vbuf[:n], n)

	uu, m := binary.Uvarint(vbuf[:n])
	fmt.Printf("Uvarint decode -> %d (read=%d)\n", uu, m)

	// signed variant
	var s int64 = -123
	n = binary.PutVarint(vbuf, s)
	ss, _ := binary.Varint(vbuf[:n])
	fmt.Printf("Varint(-123) -> % X; decode -> %d\n", vbuf[:n], ss)
}
