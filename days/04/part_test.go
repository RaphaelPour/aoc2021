package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExampleInput(t *testing.T) {
	input := util.LoadString("input_example")

	boards, turns := LoadBingoInput("input_example")

	require.NotNil(t, boards)
	require.NotNil(t, turns)

	// check if turns are parsed correctly
	turnString := strconv.Itoa(turns[0])
	for _, turn := range turns[1:len(turns)] {
		turnString += fmt.Sprintf(",%d", turn)
	}
	require.Equal(t, input[0], turnString)

	// check boards

	boardsString := ""
	for _, board := range boards {
		boardsString += fmt.Sprintf("%s\n", board)
	}

	fmt.Println(len(boards))

	require.Equal(
		t,
		strings.Join(input[2:len(input)], "\n")+"\n\n",
		boardsString,
	)
}

func TestColBingo(t *testing.T) {
	board := &Board{
		gameField: [][]Field{
			{
				{0, true},
				{0, false},
				{0, false},
			},
			{
				{0, true},
				{0, false},
				{0, false},
			},
			{
				{0, true},
				{0, false},
				{0, false},
			},
		},
	}
	require.True(t, board.IsBingo())
}

func TestRowBingo(t *testing.T) {
	board := &Board{
		gameField: [][]Field{
			{
				{0, true},
				{0, true},
				{0, true},
			},
			{
				{0, false},
				{0, false},
				{0, false},
			},
			{
				{0, false},
				{0, false},
				{0, false},
			},
		},
	}
	require.True(t, board.IsBingo())
}

func TestNoDiagonalBingo(t *testing.T) {
	board := &Board{
		gameField: [][]Field{
			{
				{0, true},
				{0, false},
				{0, false},
			},
			{
				{0, false},
				{0, true},
				{0, false},
			},
			{
				{0, false},
				{0, true},
				{0, false},
			},
		},
	}
	require.False(t, board.IsBingo())
}

func TestPart2Example(t *testing.T) {
	result := part2("input_example")
	require.Equal(t, result, 1924)
}
