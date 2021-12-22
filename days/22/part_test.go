package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRange(t *testing.T) {
	r1 := Range{10, 20}
	r2 := Range{21, 30}
	r3 := Range{15, 25}

	r, ok := r1.Intersection(r2)
	require.False(t, ok)
	require.Nil(t, r)

	r, ok = r1.Intersection(r3)
	require.True(t, ok)
	require.Equal(t, Range{15, 20}, *r)
}
