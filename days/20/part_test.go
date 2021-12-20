package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Equal(t, 35, part1(util.LoadString("input_example")))
}

func TestCalcIndex1(t *testing.T) {
	tm := NewTrenchMap([]string{
		".#",
		"",
		"...",
		"...",
		"..#",
	})

	require.Equal(t, 0, tm.CalcIndex(0, 0))
	require.Equal(t, 1, tm.CalcIndex(1, 1))
	require.Equal(t, 16, tm.CalcIndex(2, 2))
	require.Equal(t, 256, tm.CalcIndex(3, 3))
}

func TestCalcIndex2(t *testing.T) {
	tm := NewTrenchMap([]string{
		".#",
		"",
		"###",
		"###",
		"###",
	})

	require.Equal(t, 1, tm.CalcIndex(-1, -1))
	require.Equal(t, 27, tm.CalcIndex(0, 0))
	require.Equal(t, 511, tm.CalcIndex(1, 1))
	require.Equal(t, 432, tm.CalcIndex(2, 2))
	require.Equal(t, 256, tm.CalcIndex(3, 3))

	// bottom left corner
	require.Equal(t, 64, tm.CalcIndex(-1, 3))

	// top right corner
	require.Equal(t, 4, tm.CalcIndex(3, -1))
}
