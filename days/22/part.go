package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/aoc2021/util"
)

type Point struct {
	x, y, z int
}

func (p Point) String() string {
	return fmt.Sprintf("%d|%d|%d", p.x, p.y, p.z)
}

type Range struct {
	from, to int
}

func (r Range) Intersection(other Range) (*Range, bool) {
	if other.from > r.to {
		return nil, false
	}

	return &Range{
		from: util.Max(other.from, r.from),
		to:   util.Min(other.to, r.to),
	}, true
}

func (r Range) String() string {
	return fmt.Sprintf("%d..%d", r.from, r.to)
}

type RebootStep struct {
	x, y, z Range
	on      bool
}

func (r RebootStep) String() string {
	status := "off"
	if r.on {
		status = "on"
	}
	return fmt.Sprintf("%s x=%s,y=%s,z=%s", status, r.x, r.y, r.z)
}

type Cubes struct {
	cubeMap     map[Point]bool
	rebootSteps []RebootStep
}

func NewCubes(input []string) *Cubes {
	c := new(Cubes)
	c.cubeMap = make(map[Point]bool)
	c.rebootSteps = make([]RebootStep, len(input))
	pattern := regexp.MustCompile(`(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)
	for i, line := range input {
		match := pattern.FindStringSubmatch(line)
		if len(match) != 8 {
			panic(fmt.Sprintf("error parsing %s", line))
		}

		var on bool
		if match[1] == "on" {
			on = true
		}

		xFrom := util.ToInt(match[2])
		xTo := util.ToInt(match[3])
		yFrom := util.ToInt(match[4])
		yTo := util.ToInt(match[5])
		zFrom := util.ToInt(match[6])
		zTo := util.ToInt(match[7])

		c.rebootSteps[i] = RebootStep{
			Range{xFrom, xTo},
			Range{yFrom, yTo},
			Range{zFrom, zTo},
			on,
		}

		if util.Abs(xFrom) >= 50 || util.Abs(xTo) > 50 {
			continue
		}

		if util.Abs(yFrom) >= 50 || util.Abs(yTo) > 50 {
			continue
		}

		if util.Abs(zFrom) >= 50 || util.Abs(zTo) > 50 {
			continue
		}

		for z := zFrom; z <= zTo; z++ {
			for y := yFrom; y <= yTo; y++ {
				for x := xFrom; x <= xTo; x++ {
					if on {
						c.cubeMap[Point{x, y, z}] = false
					} else {
						delete(c.cubeMap, Point{x, y, z})
					}
				}
			}
		}
	}

	return c
}

func (c *Cubes) Reboot() {
	for _, step := range c.rebootSteps {
		fmt.Printf("Processing step %s\n", step)
		for z := step.z.from; z <= step.z.to; z++ {
			for y := step.y.from; y <= step.y.to; y++ {
				for x := step.x.from; x <= step.x.to; x++ {
					c.cubeMap[Point{x, y, z}] = step.on
				}
			}
		}

	}
}

func (c Cubes) CubesOnCount() int {
	count := 0
	for _, ok := range c.cubeMap {
		if ok {
			count++
		}
	}
	return count
}

func part1(input []string) int {
	return len(NewCubes(input).cubeMap)
}

func part2(input []string) int {
	c := NewCubes(input)
	c.Reboot()
	return c.CubesOnCount()
}

func main() {
	input := "input"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString(input)))
}
