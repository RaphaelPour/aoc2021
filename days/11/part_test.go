package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Equal(t, 1656, part1(util.LoadString("input_example")))
}

func TestRealInput1(t *testing.T) {
	require.Equal(t, 1743, part1(util.LoadString("input")))
}

func TestRealInput2(t *testing.T) {
	require.Equal(t, 364, part2(util.LoadString("input")))
}

func TestStrPos(t *testing.T) {
	require.Equal(t, "0,0", strPos(0, 0))
}

func TestCavern(t *testing.T) {
	c := NewCavern([]string{"11", "11"})
	require.Equal(t, 2, c.Width())
	require.Equal(t, 2, c.Height())

	c.Flash(0, 0)
	require.Equal(t, [][]int{{1, 2}, {2, 2}}, c.Field)

	c = NewCavern([]string{"999", "919", "999"})
	require.True(t, c.NextGen())
	require.Equal(t, [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, c.Field)
	require.Equal(t, 9, len(c.Visited))
}
