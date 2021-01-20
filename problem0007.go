package main

import (
	"fmt"
	"math"
)

// Problem0007 find the 10001st prime number
func Problem0007() {
	targetPrime := int64(10001)

	targetPrimeValue := sievePrimeFactor(targetPrime)
	fmt.Println("the ", targetPrime, "st prime is ", targetPrimeValue)
}

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?
func isPrime(n int64) bool {
	// bitwise operation to check even
	if n&1 == 0 {
		return false
	}
	fmt.Println("n in isPrime ", n)
	max := int64(math.Ceil(math.Sqrt(float64(n))))
	for i := int64(3); i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countPrimeFactor(targetPrime int64) int64 {
	// elapsed:  324.080372ms
	arrayOfPrimes := []int64{}
	targetPrimeValue := int64(0)
	startPrime := int64(3)
	maxPrime := int64(1100000)

	arrayOfPrimes = append(arrayOfPrimes, 2)

	for p := int64(startPrime); p <= maxPrime; p += 2 {
		if p%int64(math.Sqrt(float64(p))) != 0 || p == 3 {
			if isPrime(p) {
				arrayOfPrimes = append(arrayOfPrimes, p)
			}
		}
	}
	fmt.Println("array", arrayOfPrimes[0:10])
	targetPrimeValue = arrayOfPrimes[targetPrime-1]
	return targetPrimeValue
}

func sievePrimeFactor(targetPrime int64) int64 {
	// elapsed:  7.049759ms
	maxPrime := 1100000
	arrayOfPrimes := make([]bool, maxPrime+1)
	targetPrimeValue := int64(0)
	startPrime := 2 // start prime value at 2

	// create array of integers from 2 to n, defaulting to true
	for i := 2; i < maxPrime+1; i++ {
		arrayOfPrimes[i] = true
	}

	// initial loop from startPrime to maxPrime
	// the square of p should be less than or equal to maxPrime since within the next loop i will be squared and i needs to be less than= maxPrime
	for p := startPrime; p*p <= maxPrime; p++ {
		if arrayOfPrimes[p] == true {
			// this finds the earliest true mark in the array that is greater than p
			for i := p * p; i <= maxPrime; i += p {
				// i is incremented by p because i can now equal the next prime
				arrayOfPrimes[i] = false
			}
		}
	}

	var primeValues []int
	for p := 2; p <= maxPrime; p++ {
		if arrayOfPrimes[p] == true {
			primeValues = append(primeValues, p)
		}
	}

	fmt.Println("array", primeValues[0:10])
	targetPrimeValue = int64(primeValues[targetPrime-1])
	return targetPrimeValue
}
