package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	s := NewSnailfishNumbers([]string{
		"[[[[4,3],4],4],[7,[[8,4],9]]]",
		"[1,1]",
	})

	require.NotNil(t, s[0].root)
	require.NotNil(t, s[1].root)

	// test [1,1]
	// has children and is no literal
	require.NotNil(t, s[1].root.left)
	require.NotNil(t, s[1].root.right)
	require.False(t, s[1].root.literal)

	// children are literals
	require.True(t, s[1].root.left.literal)
	require.True(t, s[1].root.right.literal)
	require.Nil(t, s[1].root.left.left)
	require.Nil(t, s[1].root.left.right)
	require.Nil(t, s[1].root.right.left)
	require.Nil(t, s[1].root.right.right)

	// check neighbors
	require.Nil(t, s[1].root.previous)
	require.Nil(t, s[1].root.next)
}
