package main

import "fmt"

// Problem0001 finds the sum of all multiples of 3 or 5. Use 0 for the limit argument to default to 1000
func Problem0001(limit int) (int, int) {
	//  If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
	//  Find the sum of all the multiples of 3 or 5 below 1000.
	if limit == 0 {
		limit = 1000
	}
	var m3, m5, m15, theory, looped int
	//   Implementing math theory
	//   sum = n(n+1)/2
	multipleSum := func(limit, i int) int {
		multiples := (limit - 1) / i
		return (multiples * i * (multiples + 1)) / 2
	}
	m3 = multipleSum(limit, 3)
	m5 = multipleSum(limit, 5)
	m15 = multipleSum(limit, 15)
	theory = m3 + m5 - m15

	// Implementing a loop with modulo operator
	for i := 0; i < limit; i++ {
		if i%3 == 0 {
			looped += i
		}
		if i%5 == 0 && i%15 != 0 {
			looped += i
		}
	}

	fmt.Printf("Theory value %d, Looped value %d", theory, looped)
	return theory, looped
}
