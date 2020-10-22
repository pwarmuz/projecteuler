package main

import (
	"testing"
)

func TestProblem0002(t *testing.T) {
	testCases := []struct {
		limits  uint
		answers fibSum
	}{
		{3, 2},
		{4, 2},
		{5, 2},
		{10, 10},
		{300000, 257114},
		{4000000, 4613732},
	}

	for _, value := range testCases {
		var got fibSum
		fib := got.Fibonacci(value.limits)
		fib()
		t.Log("for limit ", value.limits, " want ", value.answers, " and got ", got)
		if got != value.answers {
			t.Error(
				"For:", value.limits,
				"Want:", value.answers,
				"Got:", got)
		}
	}
}
