package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

type Point struct {
	x, y int
}

type Heightmap struct {
	cave         [][]int
	lowestPoints []Point
	basins       map[int]int
}

func NewHeightmap(input []string) *Heightmap {
	h := new(Heightmap)
	h.cave = make([][]int, len(input))
	h.basins = make(map[int]int)
	for y, row := range input {
		h.cave[y] = make([]int, len(row))
		for x := 0; x < len(row); x++ {
			var err error
			h.cave[y][x], err = strconv.Atoi(string(input[y][x]))
			if err != nil {
				panic(fmt.Sprintf("error converting %d", input[y][x]))
			}
		}
	}

	return h
}

func (h *Heightmap) RiskLevel() int {
	level := 0

	for y, row := range h.cave {
		for x, cell := range row {
			lowest := true
			if y-1 >= 0 {
				// top neighbor
				if h.cave[y-1][x] <= cell {
					lowest = false
				}
			}
			if y+1 < len(h.cave) {
				// bottom neighbor
				if h.cave[y+1][x] <= cell {
					lowest = false
				}
			}

			if x-1 >= 0 {
				// left neighbor
				if h.cave[y][x-1] <= cell {
					lowest = false
				}
			}

			if x+1 < len(row) {
				// right neighbor
				if h.cave[y][x+1] <= cell {
					lowest = false
				}
			}

			if lowest {
				// fmt.Printf("\033[32m%d\033[0m", cell)
				h.lowestPoints = append(h.lowestPoints, Point{x, y})
				level += cell + 1
			} else {
				// fmt.Printf("\033[31m%d\033[0m", cell)
			}
		}
		// fmt.Println("")
	}
	return level
}

func (h *Heightmap) Fill(x, y, basinID int) {
	// Boundary check
	if y >= len(h.cave) || y < 0 || x >= len(h.cave[0]) || x < 0 {
		return
	}

	// check if border or fill goal has been rechaed
	if h.cave[y][x] == 9 || h.cave[y][x] == basinID {
		return
	}

	h.cave[y][x] = basinID
	h.basins[basinID]++

	h.Fill(x+1, y, basinID)
	h.Fill(x-1, y, basinID)
	h.Fill(x, y+1, basinID)
	h.Fill(x, y-1, basinID)
}

func (h Heightmap) BasinProduct2() int {
	for i, point := range h.lowestPoints {
		basinID := (i + 1) * 10
		h.Fill(point.x, point.y, basinID)
	}

	values := make([]int, 0)
	for _, sum := range h.basins {
		values = append(values, sum)
	}

	sort.Ints(values)
	return values[len(values)-1] * values[len(values)-2] * values[len(values)-3]

}

func part1(input []string) int {
	h := NewHeightmap(input)
	return h.RiskLevel()
}

func part2(input []string) int {
	h := NewHeightmap(input)
	_ = h.RiskLevel()
	return h.BasinProduct2()
}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadDefaultString()))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadDefaultString()))
	fmt.Println("bad: 8365427, 40392")
}
