/*
   @Project:            projecteuler
   @Author:             Phil
   @Date:               2017-05-07 21:17:14
   +Last Modified by:   Phil
   +Last Modified time: 2017-05-07 21:30:43
*/

package main

import (
	"testing"
)

type testvaluesuint struct {
	limits  uint
	answers uint
}

var tests0002 = []testvaluesuint{
	{5, 2},
	{10, 10},
	{300000, 33},
}

func TestProblem0002(t *testing.T) {
	var got uint
	sum = 0
	fib := fibonacci()
	for _, values := range tests0002 {
		for i := 0; i < 1000; i++ {
			fib(values.limits)
		}
		got = sum
		if got != values.answers {
			t.Error(
				"For", values.limits,
				"Expected", values.answers,
				"Got:", got)
		}
	}
}
