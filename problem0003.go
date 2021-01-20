package main

import (
	"fmt"
)

func Problem0003a() {
	//
	// The prime factors of 13195 are 5, 7, 13 and 29.
	// What is the largest prime factor of the number 600851475143 ?
	//
	n := int64(600851475143)
	largest := LargestPrime(n)
	fmt.Printf("Largest: %d\n", largest)
}

// LargestPrime returns the largest prime of n
func LargestPrime(n int64) int64 {
	var largest int64
	// take care of n = 2
	if (n % 2) == 0 {
		largest = 2
	}

	// all primes larger than 2 are odd which is why i = i + 2
	for i := int64(3); i <= n; i += 2 {
		if (n % i) == 0 {
			n = n / i
			largest = i
		}
	}

	return largest
}
