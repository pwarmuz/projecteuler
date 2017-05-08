/*
   @Project:            projecteuler
   @Author:             Phil
   @Date:               2017-05-07 20:26:21
   +Last Modified by:   Phil
   +Last Modified time: 2017-05-07 21:13:10
*/

package main

import (
	"testing"
)

type testvalues struct {
	limits  int
	answers int
}

var tests = []testvalues{
	{1000, 233168},
	{10, 23},
	{10, 33},
}

func TestProblem0001(t *testing.T) {
	var theory, programmatic int
	for _, values := range tests {
		theory, programmatic = problem0001(values.limits)
		if theory != values.answers || programmatic != values.answers {
			t.Error(
				"For", values.limits,
				"Expected", values.answers,
				"got theory:", theory,
				"got programmatic:", programmatic)
		}
	}
}
