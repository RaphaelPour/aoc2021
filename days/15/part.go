package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

const (
	INFINITY = int(^uint(0) >> 1)
)

var (
	ORIGIN = Point{0, 0}
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
	n := (p.x == other.x && util.Abs(p.y-other.y) == 1) ||
		(util.Abs(p.x-other.x) == 1 && p.y == other.y)
	if n {
		//fmt.Println(p, "and", other, "are neighbors")
	}

	return n
}

func (p Point) Neighbors(bounds Point) []Point {
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

func NewCave(input []string, expand bool) *Cave {
	cave := new(Cave)
	cave.travelCost = make(map[Point]int)

	offset := 1
	if expand {
		offset = 5
	}
	cave.bounds = Point{x: len(input[0]) * offset, y: len(input) * offset}

	for offY := 0; offY < offset; offY++ {
		for offX := 0; offX < offset; offX++ {
			for y := 0; y < len(input); y++ {
				for x := 0; x < len(input[y]); x++ {
					cost, err := strconv.Atoi(string(input[y][x]))
					if err != nil {
						panic(fmt.Sprintf("error converting %s to int", string(input[y][x])))
					}
					cost += offX + offY
					if cost > 9 {
						// add 1, since 0 is not allowed
						cost = (cost % 10) + 1
					}
					cave.travelCost[Point{x + len(input[0])*offX, y + len(input)*offY}] = cost
				}
			}
		}
	}

	// start node has always a cost of zero
	cave.travelCost[ORIGIN] = 0

	return cave
}

func (c *Cave) LowestPathCost() (int, Points) {
	costs := make(map[Point]int)

	// current working node
	workingNode := Point{x: 0, y: 0}

	// stores all working nodes to avoid duplicates
	visited := map[Point]bool{workingNode: true}

	// stores nodes and their predecessor
	previous := make(map[Point]Point)

	// set cost of start node to zero, as djikstra and AoC description suggests
	costs[workingNode] = 0

	// initialize all costs of the start node's neighbors
	for _, neigh := range workingNode.Neighbors(c.bounds) {
		costs[neigh] = INFINITY
	}
	//fmt.Println("costs:", costs)

	for len(visited) < c.bounds.Area() {
		//fmt.Println("------------------")
		// c.Dump(workingNode)
		min := INFINITY
		for node, newCost := range costs {
			if _, ok := visited[node]; ok {
				// skip already visited
				continue
			}

			if newCost < min {
				min = newCost
				workingNode = node
			}
		}

		visited[workingNode] = true
		//fmt.Printf("min[cost]: %s (%d)\n", workingNode, min)

		// update costs
		for _, neigh := range workingNode.Neighbors(c.bounds) {
			if _, ok := visited[neigh]; ok {
				// skip already visited
				continue
			}
			cost := costs[workingNode] + c.travelCost[neigh]
			if _, ok := costs[neigh]; !ok || cost < costs[neigh] {
				costs[neigh] = cost
				previous[neigh] = workingNode
			}

		}
		//fmt.Println("costs:", costs)
		//fmt.Println("previous:", previous)
	}

	//fmt.Println("previous:", previous)

	totalCost := 0
	goal := Point{c.bounds.x - 1, c.bounds.y - 1}
	path := make(Points, 0)
	for n := goal; n != ORIGIN; n = previous[n] {
		//fmt.Printf("cost(%s) = %d\n", n, c.travelCost[n])
		totalCost += c.travelCost[n]
		path = append(path, n)
	}

	// reverse list
	path = append(path, ORIGIN)
	for i := 0; i < int(len(path)/2); i++ {
		path[i], path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}

	// c.Dump(path...)

	return totalCost, path
}

func (c Cave) Dump(path ...Point) {
	pointMap := make(map[Point]bool)
	for _, p := range path {
		pointMap[p] = true
	}
	for y := 0; y < c.bounds.y; y++ {
		for x := 0; x < c.bounds.x; x++ {
			p := Point{x, y}
			if _, ok := pointMap[p]; ok {
				fmt.Printf("\033[32m%d \033[0m", c.travelCost[p])
			} else {
				fmt.Printf("\033[31m%d \033[0m", c.travelCost[p])
			}
		}
		fmt.Println("")
	}
}

func part1(input []string) int {
	c := NewCave(input, false)
	cost, path := c.LowestPathCost()
	c.Dump(path...)
	return cost
}

func part2(input []string) int {
	c := NewCave(input, true)
	cost, path := c.LowestPathCost()
	c.Dump(path...)
	return cost
}

func main() {
	input := "input"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString(input)))
}
