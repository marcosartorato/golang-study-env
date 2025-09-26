package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("== basic numbers (package-level RNG) ==")
	fmt.Println("Intn(10):", rand.Intn(10))
	fmt.Println("Float64():", rand.Float64())

	fmt.Println("\n== independent deterministic stream (fixed seed) ==")
	// fixed seed -> deterministic stream
	r := rand.New(rand.NewSource(42))
	fmt.Println("r.Intn(100):", r.Intn(100))
	fmt.Println("r.Intn(100):", r.Intn(100))

	fmt.Println("\n== independent non-deterministic stream (time-based seed) ==")
	// the seed is different at each run -> different stream
	// unless you need a custom RNG, you can use rand.Intn() directly
	// which uses a global rand.Source seeded by the current time
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println("r.Intn(100):", r.Intn(100))
	fmt.Println("r.Intn(100):", r.Intn(100))

	// permutations and shuffling
	fmt.Println("\n== permutations & shuffle ==")
	fmt.Println("Perm(5):", rand.Perm(5))
	s := []string{"a", "b", "c", "d"}
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	fmt.Println("Shuffled:", s)
}
