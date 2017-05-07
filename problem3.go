/*
   @Project:            projecteuler
   @Author:             Phil
   @Date:               2017-05-06 21:22:38
   +Last Modified by:   Phil
   +Last Modified time: 2017-05-07 01:50:29
*/

package main

import (
	"fmt"
	"sort"
)

func problem3() {
	/*  The prime factors of 13195 are 5, 7, 13 and 29.
	    What is the largest prime factor of the number 600851475143 ?
	*/
	number := 600851475143
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	sort.Sort(sort.Reverse(sort.IntSlice(primes)))
	fmt.Println("primes", primes)

	var makePrimes []int
	getPrimes := func() {
		for _, v := range primes {
			if number%v == 0 {
				makePrimes = append(makePrimes, v)
				number = number % v
			}
		}
	}
	getPrimes()
	fmt.Println("primes", makePrimes)
}
