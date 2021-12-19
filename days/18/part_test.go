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

	// s[0].Add(s[1])
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

	// test 1 is left of 2
	node, ok := twoNode.LeftLiteral()
	require.True(t, ok)
	require.Equal(t, leftNode, node)

	// test 1 is left of 3
	node, ok = threeNode.LeftLiteral()
	require.True(t, ok)
	require.Equal(t, leftNode, node)
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
}

func TestWithoutGrandparent(t *testing.T) {
	s := SnailfishNumber{input: "[1,2]"}
	s.root = new(Node)
	s.parse(s.root)

	node, ok := s.root.left.LeftLiteral()
	require.False(t, ok)
	require.Nil(t, node)

	node, ok = s.root.right.RightLiteral()
	require.False(t, ok)
	require.Nil(t, node)
}

func TestAlreadyOuterMost(t *testing.T) {
	s := SnailfishNumber{input: "[3,[1,2]]"}
	s.root = new(Node)
	s.parse(s.root)

	node, ok := s.root.right.right.RightLiteral()
	require.False(t, ok)
	require.Nil(t, node)

	s = SnailfishNumber{input: "[[1,2],3]"}
	s.root = new(Node)
	s.parse(s.root)

	node, ok = s.root.left.left.LeftLiteral()
	require.False(t, ok)
	require.Nil(t, node)
}

func TestSplit1(t *testing.T) {
	s := SnailfishNumber{input: "[1,0]"}
	s.root = new(Node)
	s.parse(s.root)

	// artificialy increse right side to 10 since the input can only have
	// numbers from 0 to 9
	s.root.right.number = 10

	require.True(t, s.root.right.literal)

	ok := s.root.right.Split()
	require.True(t, ok)
	require.False(t, s.root.right.literal)
	require.NotNil(t, s.root.right.right)
	require.NotNil(t, s.root.right.left)
	require.Equal(t, 5, s.root.right.right.number)
	require.Equal(t, 5, s.root.right.left.number)
}

func TestSplit2(t *testing.T) {
	s := SnailfishNumber{input: "[1,0]"}
	s.root = new(Node)
	s.parse(s.root)

	// artificialy increse right side to 10 since the input can only have
	// numbers from 0 to 9
	s.root.right.number = 11
	require.True(t, s.root.right.Split())
	require.False(t, s.root.right.Split())
	require.Equal(t, 6, s.root.right.right.number)
	require.Equal(t, 5, s.root.right.left.number)
}

func TestSplit3(t *testing.T) {
	s := SnailfishNumber{input: "[1,0]"}
	s.root = new(Node)
	s.parse(s.root)

	// artificialy increse right side to 10 since the input can only have
	// numbers from 0 to 9
	s.root.right.number = 11

	require.True(t, s.root.Split())
}

func TestSplit4(t *testing.T) {
	s := SnailfishNumber{input: "[1,0]"}
	s.root = new(Node)
	s.parse(s.root)

	// artificialy increse right side to 10 since the input can only have
	// numbers from 0 to 9
	s.root.left.number = 11

	require.True(t, s.root.Split())
}

func TestExplode1(t *testing.T) {
	s := SnailfishNumber{input: "[[[[[9,8],1],2],3],4]"}
	s.root = new(Node)
	s.parse(s.root)

	require.True(t, s.root.Explode(0))
	require.Equal(t, "[[[[0,9],2],3],4]", s.root.String())
	require.False(t, s.root.Explode(0))
}

func TestExplode2(t *testing.T) {
	s := SnailfishNumber{input: "[7,[6,[5,[4,[3,2]]]]]"}
	s.root = new(Node)
	s.parse(s.root)

	require.True(t, s.root.Explode(0))
	require.Equal(t, "[7,[6,[5,[7,0]]]]", s.root.String())
	require.False(t, s.root.Explode(0))
}

func TestExplode3(t *testing.T) {
	s := SnailfishNumber{input: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"}
	s.root = new(Node)
	s.parse(s.root)

	require.True(t, s.root.Explode(0))
	require.Equal(t, "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", s.root.String())
	require.True(t, s.root.Explode(0))
	require.Equal(t, "[[3,[2,[8,0]]],[9,[5,[7,0]]]]", s.root.String())
	require.False(t, s.root.Explode(0))
}

func TestExplode4(t *testing.T) {
	s := SnailfishNumber{input: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"}
	s.root = new(Node)
	s.parse(s.root)

	require.True(t, s.root.Explode(0))
	require.Equal(t, "[[3,[2,[8,0]]],[9,[5,[7,0]]]]", s.root.String())
	require.False(t, s.root.Explode(0))
}

func TestExplode5(t *testing.T) {
	s := SnailfishNumber{input: "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"}
	s.root = new(Node)
	s.parse(s.root)

	require.True(t, s.root.Explode(0))
	require.Equal(t, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", s.root.String())
	require.False(t, s.root.Explode(0))
}

func TestAdditionManually(t *testing.T) {
	s := SnailfishNumber{input: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"}
	s.root = new(Node)
	s.parse(s.root)

	// first two explodes
	require.True(t, s.root.Explode(0))
	require.Equal(t, "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]", s.root.String())
	require.True(t, s.root.Explode(0))
	require.Equal(t, "[[[[0,7],4],[15,[0,13]]],[1,1]]", s.root.String())
	require.False(t, s.root.Explode(0))

	// two splits
	require.True(t, s.root.Split())
	require.Equal(t, "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", s.root.String())
	require.True(t, s.root.Split())
	require.Equal(t, "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", s.root.String())
	require.False(t, s.root.Split())

	// last explode
	require.True(t, s.root.Explode(0))
	require.Equal(t, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", s.root.String())

	// nothing to do anymore
	require.False(t, s.root.Explode(0))
	require.False(t, s.root.Split())
}

func TestAddition1(t *testing.T) {
	input := []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
	}
	s := NewSnailfishNumbers(input)
	s[0].Add(s[1:]...)

	require.Equal(t, "[[[[1,1],[2,2]],[3,3]],[4,4]]", s[0].root.String())
}

func TestAddition2(t *testing.T) {
	input := []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
		"[5,5]",
	}
	s := NewSnailfishNumbers(input)
	s[0].Add(s[1:]...)

	require.Equal(t, "[[[[3,0],[5,3]],[4,4]],[5,5]]", s[0].root.String())
}

func TestAddition3(t *testing.T) {
	input := []string{
		"[1,1]",
		"[2,2]",
		"[3,3]",
		"[4,4]",
		"[5,5]",
		"[6,6]",
	}
	s := NewSnailfishNumbers(input)
	s[0].Add(s[1:]...)

	require.Equal(t, "[[[[5,0],[7,4]],[5,5]],[6,6]]", s[0].root.String())
}

func TestAddition4(t *testing.T) {
	input := []string{
		"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
		"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
		"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
		"[7,[5,[[3,8],[1,4]]]]",
		"[[2,[2,2]],[8,[8,1]]]",
		"[2,9]",
		"[1,[[[9,3],9],[[9,0],[0,7]]]]",
		"[[[5,[7,4]],7],1]",
		"[[[[4,2],2],6],[8,7]]",
	}
	s := NewSnailfishNumbers(input)
	s[0].Add(s[1:]...)

	require.Equal(t, "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", s[0].root.String())
}

func TestMagnitude1(t *testing.T) {
	s := SnailfishNumber{input: "[9,1]"}
	s.root = new(Node)
	s.parse(s.root)

	require.Equal(t, 29, s.root.Magnitude())
}

func TestMagnitude2(t *testing.T) {
	s := SnailfishNumber{input: "[[1,2],[[3,4],5]]"}
	s.root = new(Node)
	s.parse(s.root)

	require.Equal(t, 143, s.root.Magnitude())
}

func TestMagnitude3(t *testing.T) {
	s := SnailfishNumber{input: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"}
	s.root = new(Node)
	s.parse(s.root)

	require.Equal(t, 1384, s.root.Magnitude())
}

func TestMagnitude4(t *testing.T) {
	s := SnailfishNumber{input: "[[[[1,1],[2,2]],[3,3]],[4,4]]"}
	s.root = new(Node)
	s.parse(s.root)

	require.Equal(t, 445, s.root.Magnitude())
}

func TestMagnitude5(t *testing.T) {
	s := SnailfishNumber{input: "[[[[3,0],[5,3]],[4,4]],[5,5]]"}
	s.root = new(Node)
	s.parse(s.root)

	require.Equal(t, 791, s.root.Magnitude())
}

func TestMagnitude6(t *testing.T) {
	s := SnailfishNumber{input: "[[[[5,0],[7,4]],[5,5]],[6,6]]"}
	s.root = new(Node)
	s.parse(s.root)

	require.Equal(t, 1137, s.root.Magnitude())
}

func TestMagnitude7(t *testing.T) {
	s := SnailfishNumber{input: "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"}
	s.root = new(Node)
	s.parse(s.root)

	require.Equal(t, 3488, s.root.Magnitude())
}

func TestExampleHomeworkPart1(t *testing.T) {
	input := []string{
		"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		"[[[5,[2,8]],4],[5,[[9,9],0]]]",
		"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
		"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
		"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
		"[[[[5,4],[7,7]],8],[[8,3],8]]",
		"[[9,3],[[9,9],[6,[4,9]]]]",
		"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
		"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
	}

	require.Equal(t, 4140, part1(input))
}
