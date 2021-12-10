package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExampleBad(t *testing.T) {
	require.Equal(t, parScore["}"], part1([]string{"{([(<{}[<>[]}>{[]{[(<()>"}))
	require.Equal(t, parScore[")"], part1([]string{"[[<[([]))<([[{}[[()]]]"}))
	require.Equal(t, parScore["]"], part1([]string{"[{[{({}]{}}([{[{{{}}([]"}))
	require.Equal(t, parScore[")"], part1([]string{"[<(<(<(<{}))><([]([]()"}))
	require.Equal(t, parScore[">"], part1([]string{"<{([([[(<>()){}]>(<<{{"}))
}

func TestExampleGod(t *testing.T) {
	require.Equal(t, 0, part1([]string{"[({(<(())[]>[[{[]{<()<>>"}))
	require.Equal(t, 0, part1([]string{"[(()[<>])]({[<{<<[]>>("}))
	require.Equal(t, 0, part1([]string{"(((({<>}<{<{<>}{[]{[]{}"}))
	require.Equal(t, 0, part1([]string{"{<[[]]>}<{[{[{[]{()[[[]"}))
	require.Equal(t, 0, part1([]string{"<{([{{}}[<[[[<>{}]]]>[]]"}))
}

func TestExample(t *testing.T) {
	require.Equal(t, 26397, part1(util.LoadString("input_example")))
}

func TestExample2(t *testing.T) {
	require.Equal(t, 288957, part2(util.LoadString("input_example")))
}

func TestReduce1(t *testing.T) {
	l := NewLine("{}")
	ok := l.Reduce()
	require.Equal(t, 0, l.score)
	require.True(t, ok)
}

func TestReduce2(t *testing.T) {
	l := NewLine("<>")
	ok := l.Reduce()
	require.Equal(t, 0, l.score)
	require.True(t, ok)
}

func TestReduce3(t *testing.T) {
	l := NewLine("{}")
	ok := l.Reduce()
	require.Equal(t, 0, l.score)
	require.True(t, ok)
}

func TestReduce4(t *testing.T) {
	l := NewLine("[]")
	ok := l.Reduce()
	require.Equal(t, 0, l.score)
	require.True(t, ok)
}

func TestReduce5(t *testing.T) {
	l := NewLine("{()}")
	ok := l.Reduce()
	require.Equal(t, 0, l.score)
	require.True(t, ok)
}

func TestReduce6(t *testing.T) {
	l := NewLine("{()()}")
	ok := l.Reduce()
	require.Equal(t, 0, l.score)
	require.True(t, ok)
}

func TestReduce7(t *testing.T) {
	l := NewLine("()()")
	ok := l.Reduce()
	require.Equal(t, 0, l.score)
	require.True(t, ok)
}

func TestReduce8(t *testing.T) {
	l := NewLine("()((((()))))")
	ok := l.Reduce()
	require.Equal(t, 0, l.score)
	require.True(t, ok)
}

func TestBad1(t *testing.T) {
	l := NewLine("(")
	require.True(t, l.Reduce())
}

func TestBad2(t *testing.T) {
	l := NewLine("([")
	require.True(t, l.Reduce())
}

func TestBad3(t *testing.T) {
	l := NewLine("([]")
	require.True(t, l.Reduce())
}

func TestBad5(t *testing.T) {
	l := NewLine("(}")
	ok := l.Reduce()
	require.Equal(t, l.input, "}")
	require.False(t, ok)
	require.Equal(t, parScore["}"], l.score)
}

func TestReduce21(t *testing.T) {
	l := NewLine("[()()")
	ok := l.Reduce2()
	require.Equal(t, "", l.input)
	require.Equal(t, "]", l.tail)
	require.True(t, ok)
}
