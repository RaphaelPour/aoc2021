package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	input := util.LoadInt("input_example")
	require.Equal(t, 7, part1(input))
	require.Equal(t, 5, part2(input))
}

func TestPart1(t *testing.T) {
	require.Equal(t, 0, part1([]int{1, 1}))
	require.Equal(t, 0, part1([]int{1, 0}))
	require.Equal(t, 1, part1([]int{0, 1}))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 0, part2([]int{1, 1, 1}))
	require.Equal(t, 0, part2([]int{1, 1, 1, 1}))
	require.Equal(t, 1, part2([]int{0, 1, 1, 1}))
}

func TestRealInput(t *testing.T) {
	input := util.LoadDefaultInt()
	require.Equal(t, 1342, part1(input))
	require.Equal(t, 1378, part2(input))
}
