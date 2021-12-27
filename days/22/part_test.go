package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVolume(t *testing.T) {
	c := Cube{
		from: Vector3{0, 0, 0},
		to:   Vector3{2, 2, 2},
	}
	require.Equal(t, 8, c.Volume())

	c = Cube{
		from: Vector3{-10, -1, 0},
		to:   Vector3{-8, 1, 2},
	}
	require.Equal(t, 8, c.Volume())
}
