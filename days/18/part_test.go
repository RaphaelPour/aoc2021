package main

import (
	"fmt"
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
}

func TestLeftNeighbor(t *testing.T) {
	line := "[[1,[2,3]],4]"
	s := SnailfishNumber{input: line}
	s.root = new(Node)
	s.parse(s.root)

	// get all leafs as nodes
	leftNode := s.root.left.left
	require.True(t, leftNode.literal)
	require.Equal(t, 1, leftNode.number)

	twoNode := s.root.left.right.left
	require.True(t, twoNode.literal)
	require.Equal(t, 2, twoNode.number)

	threeNode := s.root.left.right.right
	require.True(t, threeNode.literal)
	require.Equal(t, 3, threeNode.number)

	rightNode := s.root.right
	require.True(t, rightNode.literal)
	require.Equal(t, 4, rightNode.number)

	// test left most
	node, ok := leftNode.LeftLiteral()
	require.Nil(t, node)
	require.False(t, ok)

	// test 1 is left of 2
	node, ok = twoNode.LeftLiteral()
	require.True(t, ok)
	require.Equal(t, leftNode, node)

	// test 1 is left of 3
	node, ok = threeNode.LeftLiteral()
	require.True(t, ok)
	require.Equal(t, leftNode, node)

	// test 1 is left of [2,3]
	node, ok = s.root.left.right.LeftLiteral()
	require.Equal(t, leftNode, node)
	require.True(t, ok)
}

func TestRightNeighbor(t *testing.T) {
	line := "[[1,[2,3]],4]"
	s := SnailfishNumber{input: line}
	s.root = new(Node)
	s.parse(s.root)

	// get all leafs as nodes
	leftNode := s.root.left.left
	require.True(t, leftNode.literal)
	require.Equal(t, 1, leftNode.number)

	twoNode := s.root.left.right.left
	require.True(t, twoNode.literal)
	require.Equal(t, 2, twoNode.number)

	threeNode := s.root.left.right.right
	require.True(t, threeNode.literal)
	require.Equal(t, 3, threeNode.number)

	rightNode := s.root.right
	require.True(t, rightNode.literal)
	require.Equal(t, 4, rightNode.number)

	node, ok := rightNode.RightLiteral()
	require.False(t, ok)
	require.Nil(t, node)

	node, ok = threeNode.RightLiteral()
	require.True(t, ok)
	require.Equal(t, rightNode, node)

	node, ok = twoNode.RightLiteral()
	require.True(t, ok)
	require.Equal(t, rightNode, node)

	fmt.Println("--------")

	node, ok = leftNode.RightLiteral()
	require.True(t, ok)
	require.Equal(t, twoNode, node)
}
