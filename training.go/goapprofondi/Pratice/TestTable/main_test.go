package main

import (
	"testing"
)

// go test .
func TestClampValues(t *testing.T) {

	var tests = []struct {
		in       int
		min      int
		max      int
		expected int
	}{
		{-1, 0, 10, 0},
		{0, 0, 10, 0},
		{3, 0, 10, 3},
		{10, 0, 10, 10},
		{15, 0, 10, 10},
	}

	for _, tt := range tests {

		v := clamp(tt.in, tt.min, tt.max)

		if v != tt.expected {
			t.Errorf("Invalid clamp(%v) result. expected=%v, got=%v", tt.in, tt.expected, v)
		}
	}
}
