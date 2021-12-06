package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Equal(t, 26, part1(util.LoadString("input_example")[0], 18))
	require.Equal(t, 5934, part1(util.LoadString("input_example")[0], 80))
}

func TestExample2(t *testing.T) {
	require.Equal(t, 256, part1(util.LoadString("input_example")[0], 26984457539))
}
