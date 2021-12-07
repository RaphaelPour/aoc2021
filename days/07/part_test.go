package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExamplePart1(t *testing.T) {
	require.Equal(t, 37, part(util.LoadIntList("input_example"), linearDistance))
}

func TestExamplePart2(t *testing.T) {
	require.Equal(t, 168, part(util.LoadIntList("input_example"), exponentialDistance))
}

func TestRealPart1(t *testing.T) {
	require.Equal(t, 352997, part(util.LoadIntList("input"), linearDistance))
}

func TestRealPart2(t *testing.T) {
	require.Equal(t, 101571302, part(util.LoadIntList("input"), exponentialDistance))
}
