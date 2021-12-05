package main

import (
	"fmt"
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	input := util.LoadString("input_example")

	require.Equal(t, 5, part1(input))
}

func TestRealInput(t *testing.T) {
	input := util.LoadString("input")

	require.Equal(t, 6710, part1(input))
}

func TestSeaMapTrivial(t *testing.T) {
	seaMap := make(SeaMap)
	seaMap.MarkPoint(0, 1)
	require.NotNil(t, seaMap[1])
	require.Equal(t, 1, seaMap.Get(0, 1))
	require.Equal(t, 0, seaMap.Overlapping())

	seaMap.MarkPoint(0, 1)
	require.Equal(t, 2, seaMap.Get(0, 1))
	require.Equal(t, 1, seaMap.Overlapping())
}

func TestSeaMapLine(t *testing.T) {
	seaMap := make(SeaMap)
	seaMap.MarkLine(0, 0, 0, 5, true)
	require.Equal(t, 1, seaMap.Get(0, 1))
	require.Equal(t, 0, seaMap.Overlapping())

	seaMap.MarkLine(0, 0, 0, 1, true)
	require.Equal(t, 2, seaMap.Get(0, 0))
	require.Equal(t, 1, seaMap.Get(0, 2))
	require.Equal(t, 2, seaMap.Overlapping())
}

func TestExamplePart2(t *testing.T) {
	input := util.LoadString("input_example")

	require.Equal(t, 12, part2(input))
}

func TestSeaMapLineDontSkipDiagonal(t *testing.T) {
	seaMap := make(SeaMap)
	seaMap.MarkLine(0, 0, 0, 5, false)
	require.Equal(t, 1, seaMap.Get(0, 1))
	require.Equal(t, 0, seaMap.Overlapping())

	seaMap.MarkLine(0, 0, 0, 1, false)
	require.Equal(t, 2, seaMap.Get(0, 0))
	require.Equal(t, 1, seaMap.Get(0, 2))
	require.Equal(t, 2, seaMap.Overlapping())

}

func TestSeaMapDiagonalTrivial(t *testing.T) {
	seaMap := make(SeaMap)
	seaMap.MarkLine(0, 0, 1, 1, false)

	require.Equal(t, 1, seaMap.Get(0, 0))
	require.Equal(t, 0, seaMap.Get(0, 1))
	require.Equal(t, 0, seaMap.Get(1, 0))
	require.Equal(t, 1, seaMap.Get(1, 1))
	require.Equal(t, 0, seaMap.Overlapping())

	seaMap.MarkLine(0, 0, 2, 2, false)
	require.Equal(t, 2, seaMap.Overlapping())
}

func TestSeaMapUpDiagonal(t *testing.T) {
	seaMap := make(SeaMap)
	seaMap.MarkLine(2, 0, 0, 2, false)

	require.Equal(t, 1, seaMap.Get(2, 0))
	require.Equal(t, 1, seaMap.Get(1, 1))
	require.Equal(t, 1, seaMap.Get(0, 2))
	require.Equal(t, 0, seaMap.Overlapping())
}

func TestSeaMapDiagonalOverlapping(t *testing.T) {
	seaMap := make(SeaMap)
	seaMap.MarkLine(0, 0, 2, 2, false)
	fmt.Println("-")
	seaMap.MarkLine(2, 0, 0, 2, false)

	require.Equal(t, 1, seaMap.Get(0, 0))
	require.Equal(t, 2, seaMap.Get(1, 1))
	require.Equal(t, 1, seaMap.Get(2, 2))
	require.Equal(t, 1, seaMap.Get(2, 0))
	require.Equal(t, 1, seaMap.Get(0, 2))

	require.Equal(t, 1, seaMap.Overlapping())
}
