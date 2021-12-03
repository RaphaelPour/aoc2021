package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	hist := NewHistogram(util.LoadString("input_example"))

	require.Equal(t, []int{7, 5, 8, 7, 5}, hist.ones)
	require.Equal(t, []int{5, 7, 4, 5, 7}, hist.zeroes)
	x, y := hist.numbers()
	require.Equal(t, 22, x)
	require.Equal(t, 9, y)
	require.Equal(t, 198, hist.powerConsumption())

	require.Equal(t, 23, hist.oxygen())
	require.Equal(t, 10, hist.carbon())
}
