package main

import "testing"

func TestProblem0003(t *testing.T) {
	tests := []struct {
		factors int64
		want    int64
	}{
		{1, 0},
		{3, 3},
		{2, 2},
		{165, 11},
		{170, 17},
		{176, 11},
		{190, 19},
		{13195, 29},
		{600851475143, 6857},
	}

	for _, test := range tests {
		got := LargestPrime(test.factors)
		if test.want != got {
			t.Errorf("failed largest prime number want %d, got%d", test.want, got)
		}
	}
}
