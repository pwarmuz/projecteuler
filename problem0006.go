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
func Problem0006() {
	naturalNumber := 100
	// get the sum of the squares
	sumOfSquares := float64(0)
	for i := 1; i <= naturalNumber; i++ {
		c := math.Pow(float64(i), 2)
		sumOfSquares += c
	}
	// get a square of the sums
	sum := 0
	for i := 1; i <= naturalNumber; i++ {
		sum += i
	}
	squareOfSum := math.Pow(float64(sum), 2)

	// difference = square of the sum - sum of squares
	difference := squareOfSum - sumOfSquares
	fmt.Printf("difference %f\n", difference)
}
