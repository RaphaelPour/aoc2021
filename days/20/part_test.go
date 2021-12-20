package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample1(t *testing.T) {
	require.Equal(t, 35, part1(util.LoadString("input_example")))
}

func TestExample2(t *testing.T) {
	require.Equal(t, 3351, part2(util.LoadString("input_example")))
}

func TestRealInput1(t *testing.T) {
	require.Equal(t, 5306, part1(util.LoadString("input")))
}

func TestRealInput2(t *testing.T) {
	require.Equal(t, 17497, part2(util.LoadString("input")))
}

func BenchmarkPart1(b *testing.B) {
	input := util.LoadString("input")
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := util.LoadString("input")
	for i := 0; i < b.N; i++ {
		part2(input)
	}
}
