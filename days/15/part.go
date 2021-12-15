package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d|%d", p.x, p.y)
}

func (p Point) WithinBounds(bounds Point) bool {
	return p.x >= 0 && p.x < bounds.x && p.y >= 0 && p.y < bounds.y
}

type Cave struct {
	fields       [][]int
	markedFields [][]int
	bounds       Point
	visited      [][]bool
}

func NewCave(input []string) *Cave {
	cave := new(Cave)
	cave.fields = make([][]int, len(input))
	cave.bounds = Point{x: len(input[0]), y: len(input)}
	var err error
	for y := 0; y < len(input); y++ {
		cave.fields[y] = make([]int, len(input[y]))
		for x := 0; x < len(input[y]); x++ {
			cave.fields[y][x], err = strconv.Atoi(string(input[y][x]))
			if err != nil {
				panic(fmt.Sprintf("error converting %s to int", string(input[y][x])))
			}
		}
	}
	return cave
}

func (c *Cave) resetVisited() {
	c.visited = make([][]bool, len(c.fields))
	for y := 0; y < len(c.fields); y++ {
		c.visited[y] = make([]bool, len(c.fields[y]))
	}
}

func (c *Cave) StartMark() {
	c.resetVisited()
	c.markedFields = make([][]int, len(c.fields))
	for y := 0; y < len(c.fields); y++ {
		c.markedFields[y] = make([]int, len(c.fields[y]))
	}
	c.Mark(Point{0, 0}, 0)
}

func (c *Cave) Mark(start Point, score int) {
	if !start.WithinBounds(c.bounds) || c.visited[start.y][start.x] {
		return
	}

	// mark as visited to avoid double summing
	c.visited[start.y][start.x] = true

	// add current score to field
	c.markedFields[start.y][start.x] = c.fields[start.y][start.x] + score

	// mark adjacent fields only in pos. directions (don't do diagonals)
	c.Mark(Point{start.x + 1, start.y}, c.markedFields[start.y][start.x])
	c.Mark(Point{start.x, start.y + 1}, c.markedFields[start.y][start.x])
}

func (c *Cave) StartSearch() int {
	c.resetVisited()
	score, path := c.Search(Point{0, 0}, Point{c.bounds.x - 1, c.bounds.y - 1})
	c.DumpPath(path)
	return score
}

func (c *Cave) Search(current, goal Point) (int, []Point) {
	if current == goal {
		return c.fields[current.y][current.x], []Point{current}
	}

	if !current.WithinBounds(c.bounds) {
		return 0, nil
	}

	currentScore := c.fields[current.y][current.x]

	right := Point{current.x + 1, current.y}
	down := Point{current.x, current.y + 1}

	// This should never happen
	if !down.WithinBounds(c.bounds) && !right.WithinBounds(c.bounds) {
		return 0, nil
		// panic(fmt.Sprintf("Down and right are out-of-bounds from pos. %s. Shouldn't this be the goal %s?", current, goal))
	}

	// if one of both neighbors is out-of-bounds, use the other
	var visited []Point
	var neighborScore int
	if !down.WithinBounds(c.bounds) {
		neighborScore, visited = c.Search(right, goal)
	} else if !right.WithinBounds(c.bounds) {
		neighborScore, visited = c.Search(down, goal)
	} else if c.markedFields[right.y][right.x] < c.markedFields[down.y][down.x] {
		// choose the neighbor with the lowest score, it's okay to be greedy
		// since we summed up everything from the start to the goal so this
		// should work
		neighborScore, visited = c.Search(right, goal)
	} else {
		neighborScore, visited = c.Search(down, goal)
	}

	return currentScore + neighborScore, append(visited, current)
}

func (c Cave) Dump() {
	for y := 0; y < len(c.fields); y++ {
		for x := 0; x < len(c.fields[y]); x++ {
			fmt.Printf("%2d ", c.markedFields[y][x])
		}
		fmt.Println("")
	}
}

func (c Cave) DumpPath(points []Point) {
	pointMap := make(map[Point]bool)
	for _, p := range points {
		pointMap[p] = true
	}

	fmt.Println("---")
	for y := 0; y < len(c.fields); y++ {
		for x := 0; x < len(c.fields[y]); x++ {
			if _, ok := pointMap[Point{x, y}]; ok {
				fmt.Printf("\033[32m%2d \033[0m", c.markedFields[y][x])
			} else {
				fmt.Printf("\033[31m%2d \033[0m", c.markedFields[y][x])
			}
		}
		fmt.Println("")
	}
}

func part1(input []string) int {
	c := NewCave(input)
	c.StartMark()
	c.Dump()

	return c.StartSearch()
}

func part2() {

}

func main() {
	input := "input_example"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	part2()
}
