package main

import "testing"

func TestSumOfSquares(t *testing.T) {
	tests := []struct {
		natural int
		want    float64
	}{
		{10, 2640.0},
		{100, 25164150.0},
	}

	for _, test := range tests {
		result := SumOfSquares(test.natural)
		if test.want != result {
			t.Errorf("failed equivalent sum of squares; want %f, got %f\n", test.want, result)
		}
	}
}
