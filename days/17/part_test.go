package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Equal(t, 45, part1(util.LoadString("input_example")[0]))
}

func TestPart1(t *testing.T) {
	require.Equal(t, 8256, part1(util.LoadString("input")[0]))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 2326, part2(util.LoadString("input")[0]))
}
