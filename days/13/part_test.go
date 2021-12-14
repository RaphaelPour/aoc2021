package main

import (
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
	require.True(t, ok)
	require.Equal(t, 4, len(cells))
}

func TestYFold(t *testing.T) {
	input := []string{
		"2,0", "", "fold along x=1",
	}
	p := NewPaper(input, false)

	require.Equal(t, 1, len(p.folds))
	require.Equal(t, "x", p.folds[0].axis)
	require.Equal(t, 1, p.folds[0].offset)

	require.Equal(t, 1, len(p.fields))

	p.Fold()
	require.Equal(t, 1, len(p.fields))
	require.Equal(t, 1, len(p.fields[0]))
	require.True(t, p.fields[0][0])
}
