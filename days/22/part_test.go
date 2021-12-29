package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestExample1Part1(t *testing.T) {
	require.Equal(t, 39, part1(util.LoadString("input_example1")))
}

func TestExample3Part1(t *testing.T) {
	require.Equal(t, 27, part1(util.LoadString("input_example3")))
}

func TestPart1(t *testing.T) {
	require.Equal(t, 601104, part1(util.LoadString("input")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 1262883317822267, part2(util.LoadString("input")))
}

func TestVolume(t *testing.T) {
	c := Cube{
		on: true,
		x:  Range{0, 2},
		y:  Range{0, 2},
		z:  Range{0, 2},
	}
	require.Equal(t, 27, c.Volume())

	c = Cube{
		on: false,
		x:  Range{-10, -8},
		y:  Range{-1, 1},
		z:  Range{0, 2},
	}
	require.Equal(t, -27, c.Volume())
}

func TestRange(t *testing.T) {
	r1 := Range{0, 1}
	r2 := Range{1, 2}
	r3 := Range{2, 3}
	r4 := Range{0, 3}
	r5 := Range{1, 4}

	/* end of r1 touch start of r2 */
	intersect, ok := r1.Intersect(r2)
	require.True(t, ok)
	require.Equal(t, Range{1, 1}, *intersect)

	/* ...and the other way around */
	intersect, ok = r2.Intersect(r1)
	require.True(t, ok)
	require.Equal(t, Range{1, 1}, *intersect)

	/* r1 is within r2 touching start */
	intersect, ok = r1.Intersect(r4)
	require.True(t, ok)
	require.Equal(t, r1, *intersect)

	/* ...and the other way around */
	intersect, ok = r1.Intersect(r4)
	require.True(t, ok)
	require.Equal(t, r1, *intersect)

	/* r3 is within r4 touching end */
	intersect, ok = r3.Intersect(r4)
	require.True(t, ok)
	require.Equal(t, r3, *intersect)

	/* ...and the other way around */
	intersect, ok = r4.Intersect(r3)
	require.True(t, ok)
	require.Equal(t, r3, *intersect)

	/* both ranges are the same */
	intersect, ok = r1.Intersect(r1)
	require.True(t, ok)
	require.Equal(t, r1, *intersect)

	/* ...and the other way around */
	intersect, ok = r1.Intersect(r1)
	require.True(t, ok)
	require.Equal(t, r1, *intersect)

	/* both neither intersect nor touch */
	intersect, ok = r1.Intersect(r3)
	require.False(t, ok)
	require.Nil(t, intersect)

	/* ...and the other way around */
	intersect, ok = r1.Intersect(r3)
	require.False(t, ok)
	require.Nil(t, intersect)

	/* r2 is completely within r4 */
	intersect, ok = r2.Intersect(r4)
	require.True(t, ok)
	require.Equal(t, r2, *intersect)

	/* ...and the other way around */
	intersect, ok = r4.Intersect(r2)
	require.True(t, ok)
	require.Equal(t, r2, *intersect)

	/* r4 and r5 overlap */
	intersect, ok = r4.Intersect(r5)
	require.True(t, ok)
	require.Equal(t, Range{1, 3}, *intersect)

	/* ...and the other way around */
	intersect, ok = r5.Intersect(r4)
	require.True(t, ok)
	require.Equal(t, Range{1, 3}, *intersect)
}

func TestCubeIntersection(t *testing.T) {
	c1 := Cube{Range{0, 2}, Range{0, 2}, Range{0, 2}, false}
	c2 := Cube{Range{1, 3}, Range{1, 3}, Range{1, 3}, false}
	c3 := Cube{Range{1, 2}, Range{1, 2}, Range{1, 2}, false}
	c4 := Cube{Range{2, 3}, Range{2, 3}, Range{2, 3}, false}
	c5 := Cube{Range{2, 2}, Range{2, 2}, Range{2, 2}, false}

	/* intersect overlapping cubes */
	intersect, ok := c1.Intersection(c2)
	require.True(t, ok)
	require.Equal(t, c3, *intersect)

	/* ... and the other way around */
	intersect, ok = c2.Intersection(c1)
	require.True(t, ok)
	require.Equal(t, c3, *intersect)

	/* intersection with itself is itself */
	intersect, ok = c1.Intersection(c1)
	require.True(t, ok)
	require.Equal(t, c1, *intersect)

	/* touching cubes have no intersection */
	intersect, ok = c3.Intersection(c4)
	require.True(t, ok)
	require.Equal(t, c5, *intersect)

	/* ... and the other way around */
	intersect, ok = c4.Intersection(c3)
	require.True(t, ok)
	require.Equal(t, c5, *intersect)
}
