/*
   @Project:            projecteuler
   @Author:             Phil
   @Date:               2017-05-06 02:21:19
   +Last Modified by:   Phil
   +Last Modified time: 2017-05-06 21:20:15
*/
package main

func problem1() (int, int) {
	/*  If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
	    Find the sum of all the multiples of 3 or 5 below 1000.
	*/
	var m3, m5, m15, theory, programmatic int
	/*  Implementing math theory
	    sum = n(n+1)/2
	*/
	limit := 1000
	multipleSum := func(limit int, i int) int {
		below := limit - 1
		multiples := below / i
		return (multiples * i * (multiples + 1)) / 2
	}
	m3 = multipleSum(limit, 3)
	m5 = multipleSum(limit, 5)
	m15 = multipleSum(limit, 15)
	theory = m3 + m5 - m15

	/*  Implementing a loop with modulo operator
	 */
	for i := 0; i < limit; i++ {
		if i%3 == 0 {
			programmatic += i
		}
		if i%5 == 0 && i%15 != 0 {
			programmatic += i
		}
	}
	return theory, programmatic
}
