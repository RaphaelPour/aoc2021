package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

const (
	INFINITY = int(^uint(0) >> 1)
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

func (p Point) AreNeighbors(other Point) bool {
	return (p.x == other.x && util.Abs(p.y-other.y) == 1) ||
		(util.Abs(p.x-other.x) == 1 && p.y == other.y)
}

func (p Point) Neigbors(bounds Point) []Point {
	n := make([]Point, 0)
	if p.x+1 < bounds.x {
		n = append(n, Point{p.x + 1, p.y})
	}

	if p.x-1 >= 0 {
		n = append(n, Point{p.x - 1, p.y})
	}

	if p.y+1 < bounds.y {
		n = append(n, Point{p.x, p.y + 1})
	}

	if p.y-1 >= 0 {
		n = append(n, Point{p.x, p.y - 1})
	}
	return n
}

func (p Point) Area() int {
	return p.x * p.y
}

type Points []Point

func (p Points) contains(other Point) bool {
	for _, point := range p {
		if point == other {
			return true
		}
	}
	return false
}

type Cave struct {
	travelCost map[Point]int
	bounds     Point
}

func NewCave(input []string) *Cave {
	cave := new(Cave)
	cave.travelCost = make(map[Point]int)
	cave.bounds = Point{x: len(input[0]), y: len(input)}
	var err error
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			cave.travelCost[Point{x, y}], err = strconv.Atoi(string(input[y][x]))
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
	visited := Points{workingNode}
	previous := make(map[Point]*Point)

	// add start node to previous list having no predecessor
	previous[workingNode] = nil

	// set cost of start node to zero, as djikstra and AoC description suggests
	costs[workingNode] = 0

	// initialize all costs of the start node's neighbors
	for _, neigh := range workingNode.Neigbors(c.bounds) {
		costs[neigh] = c.travelCost[neigh]
	}
	fmt.Println(costs)

	for len(visited) < c.bounds.Area() {
		min := INFINITY
		for node, newCost := range c.travelCost {
			if visited.contains(node) { // || !workingNode.AreNeighbors(node) {
				// skip already visited
				continue
			}

			if newCost < min {
				min = newCost
				workingNode = node
			}
		}

		visited = append(visited, workingNode)
		fmt.Printf("New working node %s\n", workingNode)

		// update costs
		for _, neigh := range workingNode.Neigbors(c.bounds) {
			if visited.contains(neigh) { // || !workingNode.AreNeighbors(node) {
				// skip already visited
				continue
			}
			cost := costs[workingNode] + c.travelCost[neigh]
			if cost < costs[neigh] {
				costs[neigh] = cost
				previous[neigh] = &workingNode
			}

		}
	}

	fmt.Println(previous)
	c.Dump(previous)

	totalCost := 0
	for n := &workingNode; n != nil; n = previous[*n] {
		fmt.Printf("cost(%s) = %d\n", *n, c.travelCost[*n])
		totalCost += c.travelCost[*n]
	}
	return totalCost
}

func (c Cave) Dump(path map[Point]*Point) {
	for y := 0; y < c.bounds.y; y++ {
		for x := 0; x < c.bounds.x; x++ {
			p := Point{x, y}
			if _, ok := path[p]; ok {
				fmt.Printf("\033[32m%d \033[0m", c.travelCost[p])
			} else {
				fmt.Printf("\033[31m%d \033[0m", c.travelCost[p])
			}
		}
		fmt.Println("")
	}
}

func part1(input []string) int {
	c := NewCave(input)
	return c.LowestPathCost()
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
