package main

import (
	"testing"
)

func TestProblem0001(t *testing.T) {
	testCases := []struct {
		limits int
		sums   int
	}{
		{1000, 233168},
		{10, 23},
	}

	var theory, looped int
	for _, values := range testCases {
		theory, looped = Problem0001(values.limits)
		if theory != values.sums || looped != values.sums {
			t.Error(
				"For:", values.limits,
				"Want:", values.sums,
				"Got theory:", theory,
				"Got programmatic:", looped)
		}
	}
}
