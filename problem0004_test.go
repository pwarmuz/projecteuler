package main

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		digits  int
		largest int
	}{
		{1, 9},
		{2, 9009},
		{3, 906609},
	}
	for _, test := range tests {
		var p Palindrome
		p.Digits(test.digits)
		if test.largest != p.Largest {
			t.Errorf("failed to get largest palindrome want %d, got %d", test.largest, p.Largest)
		}
	}
}
