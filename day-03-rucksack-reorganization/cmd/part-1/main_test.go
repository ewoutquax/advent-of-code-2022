package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, "8085", solvePuzzlePart())
}

func Test_solvePuzzlePart(t *testing.T) {
	tests := []struct {
		name      string
		wantScore string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := solvePuzzlePart(); gotScore != tt.wantScore {
				t.Errorf("solvePuzzlePart() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}
