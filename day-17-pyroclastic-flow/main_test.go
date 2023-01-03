package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1Solve(t *testing.T) {
	assert.Equal(t, 3119, solvePart1())
}

func TestPart2Solve(t *testing.T) {
	assert.Equal(t, 1536994219669, solvePart2())
}

// Explanation:
// cave.relativeHeights: [
//   {-14,0,0,-6,-4,-10,-29 1740 2693}
//   {-14,0,0,-6,-4,-10,-29 3470 5352}
//   {-14,0,0,-6,-4,-10,-29 5200 8011}
//   {-14,0,0,-6,-4,-10,-29 6930 10670}
//   {-14,0,0,-6,-4,-10,-29 8660 13329}
//   {-14,0,0,-6,-4,-10,-29 10390 15988}
//   {-14,0,0,-6,-4,-10,-29 12120 18647}
// ]
//
// Every 1.730 steps (with an offset of 10), the stack grows with 2659 rows (with an offset of 34)
//
// Task:
// We can do 1000000000000/1730 = 578034682 full rounds, each giving 2659 extra rows
// After all these rounds, since we have an offset of 140, we need to drop 1000000000000%1730-10=130 more blocks
//
// The same height will be gotten by doing 130 blocks after block 1740
//
// Calculate:
// - Do (1740+140-10) blocks, and remember the height
// - Calculate (1000000000000-(1740+140-10))/1730*2659+rememberedHeight
