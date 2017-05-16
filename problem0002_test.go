/*
   @Project:            projecteuler
   @Date:               2017-05-07 21:17:14
   +Last Modified time: 2017-05-16 03:26:33
*/

package main

import (
	"testing"
)

func TestProblem0002(t *testing.T) {
	testCases := []struct {
		limits  uint
		answers uint
	}{
		{5, 2},
		{10, 10},
		{300000, 257114},
	}

	var got uint
	fib := Fibonacci(&got)
	for _, values := range testCases {
		for i := 0; i < 1000; i++ {
			fib(values.limits)
		}
		if got != values.answers {
			t.Error(
				"For:", values.limits,
				"Want:", values.answers,
				"Got:", got)
		}
	}
}
