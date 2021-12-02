package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Equal(t, part1(util.LoadString("input_example")), 150)
	require.Equal(t, part2(util.LoadString("input_example")), 900)
}

func TestResult(t *testing.T) {
	require.Equal(t, part1(util.LoadString("input")), 1990000)
	require.Equal(t, part2(util.LoadString("input")), 1975421260)
}
