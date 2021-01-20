package main

import (
	"fmt"
	"math"
)

type Palindrome struct {
	Largest, Product1, Product2 int
}

// Problem0004 prints the largest palindrome
func Problem0004() {
	/*
		https://projecteuler.net/problem=4
		A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
		Find the largest palindrome made from the product of two 3-digit numbers.
	*/
	var p Palindrome
	p.Digits(3)
	fmt.Println("Largest palindrome", p.Largest, " p1 ", p.Product1, " p2 ", p.Product2)
}

// Digits is a palindrome method that returns palindrome data of a digits limit
func (p *Palindrome) Digits(digits int) {
	var palindrome, temp int
	high := int(math.Pow10(digits))
	low := int(math.Pow10(digits - 1))
	for i := high; i >= low; i-- {
		for n := high; n >= low; n-- {
			palindrome = i * n
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
			if p.Largest > temp {
				break
			}

			// compare palindrome and store largest palidrome
			if temp == reverse {
				p.Largest = temp
				p.Product1 = i
				p.Product2 = n
			}
		}
	}
}
