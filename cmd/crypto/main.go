package main

import (
	"crypto"
	"crypto/ed25519"
	"crypto/hmac"
	crand "crypto/rand"
	"crypto/sha256"
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

	// --- HMAC-SHA256 ---
	key := make([]byte, 32)
	if _, err := crand.Read(key); err != nil {
		panic(err)
	}
	mac := hmac.New(sha256.New, key)
	_, _ = mac.Write(msg)
	tag := mac.Sum(nil)
	fmt.Printf("HMAC-SHA256(tag) len=%d first8=%s\n", len(tag), hex.EncodeToString(tag[:8]))

	// --- Ed25519 sign/verify (Signer interface) ---
	pub, priv, err := ed25519.GenerateKey(crand.Reader)
	if err != nil {
		panic(err)
	}
	sig, err := priv.Sign(crand.Reader, msg, &ed25519.Options{
		Context: "Example_ed25519ctx",
	})
	if err != nil {
		panic(err)
	}
	if err = ed25519.VerifyWithOptions(pub, msg, sig, &ed25519.Options{
		Context: "Example_ed25519ctx",
	}); err != nil {
		panic(err)
	} else {
		fmt.Println("ed25519 verified")
	}
}
