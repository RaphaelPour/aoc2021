package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/aoc2021/util"
)

/*
 * idea from https://github.com/goggle/AdventOfCode2021.jl/blob/master/src/day22.jl
 *
 * 1) parse each line as cubeoid having from,to vector3's and a on/off state
 * 2) intersect each parsed cuboid with already parsed ones:
 * 2.1) if intersect: add intersection cuboid with off-state to the cuboid list
 * 3) if cuboid has on-state: add it to cuboid list
 */

type Range struct {
	from, to int
}

func (r Range) StartBehind(other Range) bool {
	return other.from > r.to
}

/* intersection of two sets https://scicomp.stackexchange.com/a/26260 */
func (r Range) Intersect(other Range) (*Range, bool) {
	if r.StartBehind(other) || other.StartBehind(r) {
		return nil, false
	}

	return &Range{
		from: util.Max(r.from, other.from),
		to:   util.Min(r.to, other.to),
	}, true
}

func (r Range) Length() int {
	/* add one since the end of the interval is also contained in the
	 * count of cuboids
	 */
	return util.Abs(r.to-r.from) + 1
}

func (r Range) String() string {
	return fmt.Sprintf("%d..%d", r.from, r.to)
}

type Cube struct {
	x, y, z Range
	on      bool
}

func (c Cube) Intersection(other Cube) (*Cube, bool) {
	newX, ok := c.x.Intersect(other.x)
	if !ok {
		return nil, false
	}

	newY, ok := c.y.Intersect(other.y)
	if !ok {
		return nil, false
	}

	newZ, ok := c.z.Intersect(other.z)
	if !ok {
		return nil, false
	}

	return &Cube{
		x: *newX,
		y: *newY,
		z: *newZ,
	}, true
}

func (c Cube) String() string {
	return fmt.Sprintf("%s|%s|%s", c.x, c.y, c.z)
}

func (c Cube) Volume() int {
	coeff := 1
	if !c.on {
		coeff = -1
	}
	return coeff * c.x.Length() * c.y.Length() * c.z.Length()
}

type Cubes []Cube

func (c Cubes) Volume() int {
	sum := 0
	for _, cube := range c {
		sum += cube.Volume()
	}
	return sum
}

func NewCubes(input []string, crop bool) Cubes {
	pattern := regexp.MustCompile(`(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)
	cubes := make(Cubes, 0)
	for _, line := range input {
		match := pattern.FindStringSubmatch(line)
		if len(match) != 8 {
			panic(fmt.Sprintf("error parsing %s", line))
		}

		cube := Cube{
			on: match[1] == "on",
			x: Range{
				util.ToInt(match[2]),
				util.ToInt(match[3]),
			},
			y: Range{
				util.ToInt(match[4]),
				util.ToInt(match[5]),
			},
			z: Range{
				util.ToInt(match[6]),
				util.ToInt(match[7]),
			},
		}

		if crop {
			if util.Abs(cube.x.from) >= 50 || util.Abs(cube.x.to) > 50 {
				continue
			}
			if util.Abs(cube.y.from) >= 50 || util.Abs(cube.y.to) > 50 {
				continue
			}
			if util.Abs(cube.z.from) >= 50 || util.Abs(cube.z.to) > 50 {
				continue
			}
		}

		oldLength := len(cubes)
		for i := 0; i < oldLength; i++ {
			intersection, ok := cube.Intersection(cubes[i])
			if ok {
				/* use inverted state from cube from the list:
				 * cube: we want to subtract the intersection since two cubes
				 *       would count the intersection area twice
				 * hole: add intersection again
				 */
				intersection.on = !cubes[i].on
				cubes = append(cubes, *intersection)
			}
		}

		if cube.on {
			cubes = append(cubes, cube)
		}
	}

	return cubes
}

func part1(input []string) int {
	c := NewCubes(input, true)
	return c.Volume()
}

func part2(input []string) int {
	c := NewCubes(input, false)
	return c.Volume()
}

func main() {
	input := "input"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString(input)))
}
