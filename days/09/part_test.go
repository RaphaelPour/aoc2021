package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Equal(t, 15, part1(util.LoadString("input_example")))
}

func TestExample2(t *testing.T) {
	require.Equal(t, 1134, part2(util.LoadString("input_example")))
}
