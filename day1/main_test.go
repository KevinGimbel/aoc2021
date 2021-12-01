package main

import (
	"testing"
)

func TestCalculateIncreaseCount(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "AoC example inputy",
			input:    []string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"},
			expected: 7,
		}, {
			name:     "Another test example inputy",
			input:    []string{"1", "2", "2", "3", "4", "0"},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countMessurementIncrease(tt.input)

			if got != tt.expected {
				t.Errorf("wanted %v but got %v", tt.expected, got)
			}

		})
	}
}

func TestCalculateThreeMessurementIncreaseCount(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "AoC example inputy",
			input:    []string{"607", "618", "618", "617", "647", "716", "769", "792"},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTripleMessurementIncrease(tt.input)

			if got != tt.expected {
				t.Errorf("wanted %v but got %v", tt.expected, got)
			}

		})
	}
}
