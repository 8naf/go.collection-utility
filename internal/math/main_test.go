package math

import (
	"testing"
)

func TestSign(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 0, expected: 0},
		{input: 1, expected: 1},
		{input: 5, expected: 1},
		{input: -1, expected: -1},
		{input: -10, expected: -1},
	}

	for _, test := range tests {
		output := Sign(test.input)
		if output != test.expected {
			t.Errorf(
				"Sign(%d): expected %d, got %d",
				test.input, test.expected, output,
			)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 0, expected: 0},
		{input: 1, expected: 1},
		{input: 5, expected: 5},
		{input: -1, expected: 1},
		{input: -10, expected: 10},
	}
	for _, test := range tests {
		output := Abs(test.input)
		if output != test.expected {
			t.Errorf(
				"Abs(%d): expected %d, got %d",
				test.input, test.expected, output,
			)
		}
	}
}
