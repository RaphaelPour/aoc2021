package main

import (
	"fmt"
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Equal(t, 17, part1(util.LoadString("input_example")))
}

func TestNewPaper(t *testing.T) {
	p := NewPaper(util.LoadString("input_example"), true)

	require.Equal(t, 1, len(p.folds))
	require.Equal(t, "y", p.folds[0].axis)
	require.Equal(t, 7, p.folds[0].offset)

	require.Equal(t, 9, len(p.fields))
	cells, ok := p.fields[10]
	fmt.Println(cells)
	require.True(t, ok)
	require.Equal(t, 4, len(cells))
}
