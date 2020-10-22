package main

import (
	"fmt"
	"math"
)

// Problem0006 find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
// The sum of the squares of the first ten natural numbers is, 385
// The square of the sum of the first ten natural numbers is, 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is
// 3025-385=2640
func SumOfSquares(natural int) float64 {
	// get the sum of the squares
	var sumOfSquares float64
	for i := 1; i <= natural; i++ {
		c := math.Pow(float64(i), 2)
		sumOfSquares += c
	}
	// get a square of the sums
	sum := 0
	for i := 1; i <= natural; i++ {
		sum += i
	}
	squareOfSum := math.Pow(float64(sum), 2)

	// difference = square of the sum - sum of squares
	return squareOfSum - sumOfSquares
}
func Problem0006() {
	sum := SumOfSquares(100)
	fmt.Printf("difference %f\n", sum)
}
