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

func (p Point) Area() int {
	return p.x * p.y
}

type Cave struct {
	metrics [][]int
	bounds  Point
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

func (c *Cave) LowestPathCost() int {
	costs := make(map[Point]int)
	workingNode := Point{x: 0, y: 0}
	nodes := []Point{workingNode}
	for _, node := range nodes {
		// TODO: Implement AreNeighbors
		if workingNode.AreNeighbors(workingNode) {
			costs[node] = c.metrics[node.y][node.x]
		}
	}

	for len(nodes) < c.bounds.Area() {
		min := 1000000
		var newNode Point
		for y := range c.metrics {
			for x := range c.metrics[y] {
				// TODO: check if Point{x,y} is in nodes and continue if so
			}
		}
	}
}

func part1(input []string) int {
	c := NewCave(input)
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
