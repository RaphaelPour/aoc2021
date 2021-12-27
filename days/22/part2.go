package main

import (
	"fmt"
	"regexp"

	"github.com/RaphaelPour/aoc2021/util"
)

type Vector3 struct {
	x, y, z int
}

func (v Vector3) String() string {
	return fmt.Sprintf("%d|%d|%d", v.x, v.y, v.z)
}

type Cube struct {
	from, to Vector3
	on       bool
}

func (c Cube) String() string {
	return fmt.Sprintf("(%s)..(%s)", c.from, c.to)
}

func (c Cube) Volume() int {
	return util.Abs(c.from.x+c.to.x) *
		util.Abs(c.from.y+c.to.y) *
		util.Abs(c.from.z+c.to.z)
}

type Cubes []Cube

func NewCubes(input []string) Cubes {
	pattern := regexp.MustCompile(`(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)`)
	cubes := make(Cubes, 0)
	for _, line := range input {
		match := pattern.FindStringSubmatch(line)
		if len(match) != 8 {
			panic(fmt.Sprintf("error parsing %s", line))
		}

		cube := Cube{
			on: match[1] == "on",
			from: Vector3{
				util.ToInt(match[2]),
				util.ToInt(match[4]),
				util.ToInt(match[6]),
			},
			to: Vector3{
				util.ToInt(match[3]),
				util.ToInt(match[5]),
				util.ToInt(match[7]),
			},
		}
		cubes = append(cubes, cube)
	}

	return cubes
}

func (c Cubes) Volume() int {
	sum := 0
	for _, cube := range c {
		sum += cube.Volume()
	}
	return sum
}

func part2(input []string) int {
	c := NewCubes(input)
	return c.Volume()
}

func main() {
	input := "input"
	fmt.Println(part2(util.LoadString(input)))
}
