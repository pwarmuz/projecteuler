/*
   @Project:            projecteuler
   @Author:             Phil
   @Date:               2017-05-07 21:17:14
   +Last Modified by:   Phil
   +Last Modified time: 2017-05-08 01:50:50
*/

package projecteuler_test

import (
	"testing"

	"github.com/pwarmuzStvns/projecteuler"
)

func TestProblem0002(t *testing.T) {
	testCases := []struct {
		limits  uint
		answers uint
	}{
		{5, 2},
		{10, 10},
		{300000, 33},
	}

	var got, sum uint
	sum = 0
	fib := projecteuler.Fibonacci(&sum)
	for _, values := range testCases {
		for i := 0; i < 1000; i++ {
			fib(values.limits)
		}
		got = sum
		if got != values.answers {
			t.Error(
				"For:", values.limits,
				"Want:", values.answers,
				"Got:", got)
		}
	}
}
