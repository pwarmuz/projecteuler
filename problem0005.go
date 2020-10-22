package main

import (
	"fmt"
	"math"
)

func EvenlyDivisible(limit int64) uint64 {
	/*
		https://projecteuler.net/problem=5
		2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
		What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
	*/
	var smallest uint64
	if 43 < limit || limit < 12 {
		return 0
	}
	// this max is a guesstimate
	max := int64(math.Pow10(((int(limit) / 2) - 2)))
	var i, n, product, remainder, divisible int64
	for i = 1; i <= max; i++ {
		if smallest != 0 {
			break
		}
		for n = 1; n <= 9; n++ {
			product = i * n

			if smallest != 0 {
				break
			}

			for divisible = 1; divisible <= limit; divisible++ {
				remainder = product % divisible
				if remainder != 0 {
					break
				}
				if divisible == limit {
					smallest = uint64(product)
				}
			}
		}
	}

	return smallest
}

// Problem0005 finds the smallest number evenly divisible from 1 to 20
func Problem0005() {
	limit := int64(20)
	result := EvenlyDivisible(limit)
	fmt.Printf("Evenly divisible from 1 to %d is %d\n", limit, result)
}
