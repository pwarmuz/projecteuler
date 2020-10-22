package main

import (
	"testing"
)

func TestProblem0002(t *testing.T) {
	testCases := []struct {
		limits  uint
		answers uint
	}{
		{3, 2},
		{4, 2},
		{5, 2},
		{10, 10},
		{300000, 257114},
		{4000000, 4613732},
	}

	for _, value := range testCases {
		got := SumTotal(value.limits, EVEN)
		t.Log("for limit ", value.limits, " want ", value.answers, " and got ", got)
		if got != value.answers {
			t.Error(
				"For:", value.limits,
				"Want:", value.answers,
				"Got:", got)
		}
	}
}

func TestProblem0002Seq(t *testing.T) {
	tc := struct {
		limits  uint
		answers []uint
	}{
		13, []uint{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144},
	}

	got := Sequence(tc.limits)
	t.Log("for limit ", tc.limits, " want ", tc.answers, " and got ", got)
	for i, num := range tc.answers {
		if got[i] != num {
			t.Error(
				"For:", tc.limits,
				"Want:", tc.answers,
				"Got", got[i])
		}
	}
}
