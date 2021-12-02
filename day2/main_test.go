package main

import (
	"testing"
)

func TestFollowPlannedCourse(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "AoC example inputy",
			input:    []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			expected: 150,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FollowPlannedCourse(tt.input)

			if got != tt.expected {
				t.Errorf("wanted %v but got %v", tt.expected, got)
			}

		})
	}
}

func TestFollowPlannedCoursePart2(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "AoC example inputy",
			input:    []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			expected: 900,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FollowPlannedCoursePart2(tt.input)

			if got != tt.expected {
				t.Errorf("wanted %v but got %v", tt.expected, got)
			}

		})
	}
}
