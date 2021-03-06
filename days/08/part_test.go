package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExamplePart1(t *testing.T) {
	require.Equal(t, 26, part1(util.LoadString("input_example")))
}

func TestExamplePart2(t *testing.T) {
	require.Equal(t, 5353, part2(util.LoadString("input_example_2")))
}

func TestSub(t *testing.T) {
	require.Equal(t, 1, sub("a", "ab"))
	require.Equal(t, 1, sub("ab", "b"))
	require.Equal(t, 2, sub("a", "b"))

	require.Equal(t, 0, sub("abc", "abc"))
	require.Equal(t, 0, sub("bca", "abc"))
	require.Equal(t, 1, sub("bcda", "abc"))

	require.Equal(t, 2, sub("cefabd", "eafb"))
	require.Equal(t, 3, sub("fbcad", "eafb"))

	require.Equal(t, 6, sub("cdfgeb", "ab"))
}

func TestWithin(t *testing.T) {
	require.True(t, within("ab", "abc"))
	require.True(t, within("ba", "abc"))
	require.False(t, within("gcdfa", "cefabd"))
}

func TestRewire(t *testing.T) {
	display, err := NewDisplay(util.LoadString("input_example_2")[0])
	require.Nil(t, err)
	display.Rewire()

	require.Equal(t, 0, display.Connections["cagedb"])
	require.Equal(t, 1, display.Connections["ab"])
	require.Equal(t, 4, display.Connections["eafb"])
	require.Equal(t, 7, display.Connections["dab"])
	require.Equal(t, 8, display.Connections["acedgfb"])

	require.Equal(t, 9, display.Connections["cefabd"])
	require.Equal(t, 6, display.Connections["cdfgeb"])
	require.Equal(t, 3, display.Connections["fbcad"])

	require.Equal(t, 2, display.Connections["gcdfa"])
	require.Equal(t, 5, display.Connections["cdfbe"])

	num, err := display.Result()
	require.Nil(t, err)
	require.Equal(t, 5353, num)
}

func TestRewire2(t *testing.T) {
	display, err := NewDisplay("fdb bgfa cedfg abedcf defgba gbaed dagfbce fbged dgabce fb | bf dfbega fb baged")
	require.Nil(t, err)
	display.Rewire()

	require.Equal(t, 1, display.Connections["fb"])
	require.Equal(t, 4, display.Connections["bgfa"])
	require.Equal(t, 7, display.Connections["fdb"])
	require.Equal(t, 8, display.Connections["dagfbce"])

	require.Equal(t, 5, display.Connections["gbaed"])
}

func BenchmarkRewire(b *testing.B) {
	input := util.LoadString("input_example_2")
	for i := 0; i < b.N; i++ {
		display, _ := NewDisplay(input[0])
		display.Rewire()
	}
}
