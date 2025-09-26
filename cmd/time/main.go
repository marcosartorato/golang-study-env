package main

import (
	"fmt"
	"time"
)

func main() {
	// Examples of Now, Format, and Parse
	// Example of time.Now, Time.Format, and time.RFC3339.
	fmt.Println("== Now / Format / Parse ==")
	now := time.Now()
	fmt.Println("now RFC3339:", now.Format(time.RFC3339))

	// Example of Time.Format and custom layout.
	layout := "2006-01-02 15:04:05 MST"
	fmt.Println("now custom :", now.Format(layout))

	// Example of Time.Parse
	parsedRFC, err := time.Parse(time.RFC3339, "2024-12-31T23:59:00Z")
	if err != nil {
		panic(err)
	}
	fmt.Println("parsed RFC3339:", parsedRFC)

	fmt.Println()

	// Examples of Since, Date, Add, and Sub
	fmt.Println("== Since / Add / Sub ==")
	// Example of time.Since
	start := time.Now()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("time.Since(start):", time.Since(start))

	// Example of time.Date, package constants, Time.Add, and Time.Sub
	base := time.Date(2025, time.September, 26, 12, 0, 0, 0, time.UTC)
	fmt.Println("base:", base)
	fmt.Println("base.Add(90s):", base.Add(90*time.Second))
	fmt.Println("now.Sub(base):", now.Sub(base)) // Duration (may be negative or positive)

	fmt.Println()

	// Examples of Before, After, and Equal
	fmt.Println("== Comparisons (Before / After / Equal) ==")
	t1 := time.Date(2025, 9, 26, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, 9, 26, 12, 0, 1, 0, time.UTC)
	fmt.Println("t1.Before(t2):", t1.Before(t2))
	fmt.Println("t2.After(t1):", t2.After(t1))
	fmt.Println("t1.Equal(t1):", t1.Equal(t1))

	fmt.Println()

	// Examples of ParseDuration and Add
	fmt.Println("== Durations (ParseDuration) ==")
	dur, err := time.ParseDuration("2h45m")
	if err != nil {
		panic(err)
	}
	fmt.Println(`ParseDuration("2h45m"):`, dur)
	fmt.Println("base.Add(dur):", base.Add(dur))

	fmt.Println()

	// Examples of Truncate and Round
	fmt.Println("== Rounding (Truncate / Round) ==")
	t := time.Date(2025, 9, 26, 12, 0, 0, 123_456_789, time.UTC)
	fmt.Println("t:", t)
	fmt.Println("t.Truncate(1ms):", t.Truncate(time.Millisecond))
	fmt.Println("t.Round(1ms):   ", t.Round(time.Millisecond))
	fmt.Println("t.Truncate(10ms):", t.Truncate(10*time.Millisecond))
	fmt.Println("t.Round(10ms):   ", t.Round(10*time.Millisecond))
}
