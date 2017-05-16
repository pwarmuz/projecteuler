/*
   @Project:            projecteuler
   @Date:               2017-05-07 20:26:21
   +Last Modified time: 2017-05-16 03:26:19
*/

package main

import (
	"testing"
)

func TestProblem0001(t *testing.T) {
	testCases := []struct {
		limits  int
		answers int
	}{
		{1000, 233168},
		{10, 23},
	}
	var theory, programmatic int
	for _, values := range testCases {
		theory, programmatic = Problem0001(values.limits)
		if theory != values.answers || programmatic != values.answers {
			t.Error(
				"For:", values.limits,
				"Want:", values.answers,
				"Got theory:", theory,
				"Got programmatic:", programmatic)
		}
	}
}
