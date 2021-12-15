package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPoint(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{1, 0}
	p3 := Point{0, 1}
	p4 := Point{1, 1}
	p5 := Point{2, 1}
	p6 := Point{-1, -1}
	p7 := Point{0, -1}
	p8 := Point{-1, 0}
	bounds := Point{2, 2}

	// check bounds
	require.True(t, p1.WithinBounds(bounds))
	require.True(t, p2.WithinBounds(bounds))
	require.True(t, p3.WithinBounds(bounds))
	require.True(t, p4.WithinBounds(bounds))
	require.False(t, p5.WithinBounds(bounds))
	require.False(t, p6.WithinBounds(bounds))
	require.False(t, p7.WithinBounds(bounds))
	require.False(t, p8.WithinBounds(bounds))
	require.False(t, bounds.WithinBounds(bounds))

	// check neighbors
	require.True(t, p1.AreNeighbors(p2))
	require.True(t, p2.AreNeighbors(p1))

	require.True(t, p1.AreNeighbors(p3))
	require.False(t, p1.AreNeighbors(p4))
	require.False(t, p1.AreNeighbors(p5))

	require.False(t, Point{0, 2}.AreNeighbors(Point{2, 2}))

	// get neighbors
	require.Equal(t, []Point{p2, p3}, p1.Neighbors(bounds))

	require.Equal(t, []Point{
		{2, 1},
		{0, 1},
		{1, 2},
		{1, 0},
	}, Point{1, 1}.Neighbors(Point{3, 3}))

	require.NotContains(t, Point{0, 2}.Neighbors(Point{5, 5}), Point{2, 2})
}

func TestPoints(t *testing.T) {
	points := Points{{0, 0}}

	require.True(t, points.contains(Point{0, 0}))
	require.False(t, points.contains(Point{0, 1}))
}

func TestSearch(t *testing.T) {
	c := NewCave([]string{
		"131",
		"241",
		"241",
	})

	require.Equal(t, Point{3, 3}, c.bounds)
	require.Equal(t, map[Point]int{
		Point{0, 0}: 0,
		Point{1, 0}: 3,
		Point{2, 0}: 1,
		Point{0, 1}: 2,
		Point{1, 1}: 4,
		Point{2, 1}: 1,
		Point{0, 2}: 2,
		Point{1, 2}: 4,
		Point{2, 2}: 1,
	}, c.travelCost)

	cost, path := c.LowestPathCost()
	require.Equal(t, 6, cost)
	require.NotNil(t, path)
	require.Equal(t, Points{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}, path)
}
