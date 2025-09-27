package main

import (
	"crypto"
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
)

func main() {
	msg := []byte("hello crypto")

	// --- Secure random bytes ---
	rb := make([]byte, 16)
	if _, err := crand.Read(rb); err != nil {
		panic(err)
	}
	fmt.Printf("rand(16): %s\n", hex.EncodeToString(rb))

	// --- SHA-256 (via crypto.Hash) ---
	h := crypto.SHA256.New()
	_, _ = h.Write(msg)
	sum := h.Sum(nil)
	fmt.Printf("SHA256(%q): %s\n", msg, hex.EncodeToString(sum))
}
