package main

import (
	"fmt"
)

// Problem0004 prints the largest palindrome
func Problem0004() {
	/*
		https://projecteuler.net/problem=4
		A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
		Find the largest palindrome made from the product of two 3-digit numbers.
	*/
	fmt.Println("Start Problem0004 palindrome")
	var palindrome, temp int
	largest := 0
	product1 := 0
	product2 := 0

	for i := 1000; i >= 100; i-- {
		for n := 1000; n >= 100; n-- {
			palindrome = i * n
			//fmt.Printf("[ Checking %d ]", palindrome)
			// reorder palindrome
			temp = palindrome
			reverse := 0

			for {
				remainder := palindrome % 10
				reverse = (reverse * 10) + remainder
				palindrome = palindrome / 10
				if palindrome == 0 {
					break
				}
			}
			if largest > temp {
				break
			}

			// compare palindrome
			if temp == reverse {
				largest = temp
				product1 = i
				product2 = n
				// store palindrome to compare if greater
				//fmt.Printf("[ %d is a Palindrome of %d ]", temp, reverse)
			}
		}
	}

	fmt.Println("Largest palindrome", largest, " p1 ", product1, " p2 ", product2)
}
