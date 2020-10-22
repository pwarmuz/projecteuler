package main

import (
	"fmt"
)

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

func Problem0003a() {
	n := int64(600851475143)
	largest := LargestPrime(n)
	fmt.Printf("Largest: %d\n", largest)
}
