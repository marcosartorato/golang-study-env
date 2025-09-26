package main

import (
	"fmt"
	"math"
)

func main() {
	// Example of constants
	fmt.Println("== constants ==")
	fmt.Println("Pi:", math.Pi, "E:", math.E)

	// Example of basic math functions
	fmt.Println("\n== roots & powers ==")
	fmt.Println("Sqrt(2):", math.Sqrt(2))
	fmt.Println("Pow(2,10):", math.Pow(2, 10))
	fmt.Println("Hypot(3,4):", math.Hypot(3, 4)) // 5

	// Example of trig functions
	fmt.Println("\n== trig ==")
	theta := math.Pi
	fmt.Println("Sin(pi):", math.Sin(theta))
	fmt.Println("Cos(pi):", math.Cos(theta))
	fmt.Println("Note that the values are not exact due to floating point precision limitations.")

	// Example of rounding functions
	fmt.Println("\n== rounding ==")
	x := -2.7
	y := 2.3
	fmt.Println("Floor(2.3):", math.Floor(y))
	fmt.Println("Ceil(-2.7):", math.Ceil(x))
	fmt.Println("Round(2.3):", math.Round(y))

	// Example of Modf function
	intPart, frac := math.Modf(-2.7)
	fmt.Println("Modf(-2.7) -> int:", intPart, "frac:", frac)
	fmt.Println("Note that the fractional part is not precisely -0.7.")
	fmt.Println("This issue is due to floating point precision limitations.")

	// Example of special values
	fmt.Println("\n== special values ==")
	nan := math.Log(-1) // NaN
	inf := math.Inf(1)  // +Inf
	fmt.Println("IsNaN(Log(-1)):", math.IsNaN(nan))
	fmt.Println("IsInf(+Inf,+1):", math.IsInf(inf, +1))
}
