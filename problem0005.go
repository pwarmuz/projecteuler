package main

import "fmt"

// Problem0005 finds the smallest number evenly divisible from 1 to 20
func Problem0005() {
	/*
		https://projecteuler.net/problem=5
		2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
		What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
	*/
	smallest := 0
	for i := 1; i <= 100000000; i++ {
		if smallest != 0 {
			break
		}
		for n := 1; n <= 9; n++ {
			product := i * n

			if smallest != 0 {
				break
			}

			for divisible := 1; divisible <= 20; divisible++ {
				remainder := product % divisible
				if remainder != 0 {
					break
				}
				if divisible == 20 {
					smallest = product
				}
			}
		}
	}

	fmt.Println("smallest number ", smallest)
}
