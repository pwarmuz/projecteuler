package main

import "testing"

func TestLowest(t *testing.T) {
	tests := []struct {
		limit int64
		want  uint64
	}{
		{12, 27720},
		{20, 232792560},
	}

	for _, test := range tests {
		result := EvenlyDivisible(test.limit)
		if test.want != result {
			t.Errorf("failed want %d, got %d", test.want, result)
		}
	}
}
