package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

type Cavern struct {
	Field   [][]int
	Flashes int
	Visited map[string]bool
}

func NewCavern(input []string) *Cavern {
	c := new(Cavern)
	c.Field = make([][]int, len(input))

	var err error
	for i, line := range input {
		c.Field[i] = make([]int, len(line))
		for j, cell := range line {
			c.Field[i][j], err = strconv.Atoi(string(cell))
			if err != nil {
				panic(fmt.Sprintf("%s is not a number\n", string(cell)))
			}
		}
	}

	return c
}

func (c Cavern) Width() int {
	return len(c.Field[0])
}

func (c Cavern) Height() int {
	return len(c.Field)
}

func (c Cavern) AlreadyVisited(x, y int) bool {
	_, ok := c.Visited[strPos(x, y)]
	return ok
}

func (c Cavern) Dump() {
	for _, row := range c.Field {
		for _, cell := range row {
			if cell == 0 {
				fmt.Printf("\033[32m%d\033[0m", cell)
			} else {
				fmt.Printf("\033[31m%d\033[0m", cell)
			}
		}
		fmt.Println("")
	}
}

func (c *Cavern) NextGen() bool {
	// incrase all energy levels
	for i, row := range c.Field {
		for j := range row {
			c.Field[i][j]++
		}
	}

	// track which octopus has already been flashes as they sould each
	// flash only once per generation
	c.Visited = make(map[string]bool, 0)

	// loop until no octopus is left has needs flash
	anyFlashed := true
	for anyFlashed {
		anyFlashed = false
		for i, row := range c.Field {
			for j := range row {
				if !c.AlreadyVisited(j, i) && c.Field[i][j] > 9 {
					c.Visited[strPos(j, i)] = true
					anyFlashed = true
					c.Flash(j, i)
				}
			}
		}
	}

	synchronized := true
	for i, row := range c.Field {
		for j := range row {
			if c.Field[i][j] > 9 {
				c.Flashes++
				c.Field[i][j] = 0
			}

			// check if all octopuses are synchronized
			if c.Field[i][j] != 0 {
				synchronized = false
			}
		}
	}

	return synchronized
}

// provide easy coordinate to map key conversion
// thanks to @markus.freitag
func strPos(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func (c *Cavern) Flash(x, y int) {
	// tell friends about the flash
	for _, offY := range []int{-1, 0, 1} {
		for _, offX := range []int{-1, 0, 1} {
			// skip own position
			if offX == 0 && offY == 0 {
				continue
			}

			nX, nY := x+offX, y+offY
			// check boundaries
			if nX < 0 || nX >= c.Width() || nY < 0 || nY >= c.Height() {
				continue
			}

			// increase value
			c.Field[nY][nX]++
		}
	}
}

func part1(input []string) int {
	c := NewCavern(input)
	for i := 0; i < 100; i++ {
		c.NextGen()
	}
	return c.Flashes
}

func part2(input []string) int {
	c := NewCavern(input)
	steps := 1
	for ; !c.NextGen(); steps++ {
	}
	return steps
}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString("input")))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString("input")))
}
