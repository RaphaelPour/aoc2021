package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExamplePart1(t *testing.T) {
	a := NewALU(util.LoadString("input"))
	a.Run(5)
	require.Equal(t, 0, a.variables["z"])
}
func TestExample1(t *testing.T) {
	a := NewALU([]string{"inp x", "mul x -1"})
	a.Run(1)
	require.Equal(t, -1, a.variables["x"])
}

func TestExample2(t *testing.T) {
	a := NewALU([]string{
		"inp z",
		"inp x",
		"mul z 3",
		"eql z x",
	})
	a.Run(13)
	require.Equal(t, 1, a.variables["z"])
}

func TestExample3(t *testing.T) {
	a := NewALU([]string{
		"inp w",
		"add z w",
		"mod z 2",
		"div w 2",
		"add y w",
		"mod y 2",
		"div w 2",
		"add x w",
		"mod x 2",
		"div w 2",
		"mod w 2",
	})
	a.Run(4)
	require.Equal(t, 0, a.variables["z"])
	require.Equal(t, 0, a.variables["y"])
	require.Equal(t, 1, a.variables["x"])
	require.Equal(t, 0, a.variables["w"])
}

func TestExample4(t *testing.T) {
	a := NewALU([]string{"inp x", "inp y", "inp z", "inp w"})
	a.Run(1234)
	require.Equal(t, 1, a.variables["x"])
	require.Equal(t, 2, a.variables["y"])
	require.Equal(t, 3, a.variables["z"])
	require.Equal(t, 4, a.variables["w"])
}

func TestReal1(t *testing.T) {
	a := NewALU([]string{
		"inp w",
		"mul x 0",
		"add x z",
		"mod x 26",
		"div z 1",
		"add x 14",
		"eql x w",
		"eql x 0",
		"mul y 0",
		"add y 25",
		"mul y x",
		"add y 1",
		"mul z y",
		"mul y 0",
		"add y w",
		"add y 8",
		"mul y x",
		"add z y",
	})
	a.Run(5)
	require.Equal(t, 13, a.variables["z"])
}
