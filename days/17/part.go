package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d|%d", p.x, p.y)
}

type Rectangle struct {
	p1, p2 Point
}

func (r Rectangle) String() string {
	return fmt.Sprintf("%s -> %s", r.p1, r.p2)
}

type Parabola struct {
	a0, a1, a2 float64
}

func (p Parabola) String() string {
	return fmt.Sprintf("%.4fx^2 + %.4fx + %.4f", p.a2, p.a1, p.a0)
}

func (p Parabola) Eval(x float64) float64 {
	return p.a2*x*x + p.a1*x + p.a0
}

func (p Parabola) Vertex() Point {
	x := -p.a1 / (2 * p.a0)
	y := p.Eval(x)
	return Point{int(x), int(y)}
}

func (p Parabola) IntersectsWith(area Rectangle) bool {
	y := area.p1.y

	for x := area.p1.x; x <= area.p2.x; x++ {
		if int(p.Eval(float64(x))) >= y {
			return true
		}
	}
	return false
}

func part1(input string) int {
	pattern := regexp.MustCompile(`^target area: x=(-?\d+)..(-?\d+), y=(-?\d+)..(-?\d+)$`)
	match := pattern.FindStringSubmatch(input)
	if len(match) != 5 {
		panic(fmt.Sprintf("error parsing %s: %#v", input, match))
	}

	x1, err := strconv.Atoi(match[1])
	if err != nil {
		panic(fmt.Sprintf("error converting %s", match))
	}
	x2, err := strconv.Atoi(match[2])
	if err != nil {
		panic(fmt.Sprintf("error converting %s", match))
	}
	y1, err := strconv.Atoi(match[3])
	if err != nil {
		panic(fmt.Sprintf("error converting %s", match))
	}
	y2, err := strconv.Atoi(match[4])
	if err != nil {
		panic(fmt.Sprintf("error converting %s", match))
	}

	targetArea := Rectangle{
		p1: Point{x1, y1},
		p2: Point{x2, y2},
	}

	start := Point{x: 0, y: 0}

	fmt.Println(targetArea)
	fmt.Println(start)

	p := Parabola{a2: -1, a1: float64(x1), a0: float64(y1)}

	for p.IntersectsWith(targetArea) || p.a2 < 0 {
		p.a2 += 0.1
		fmt.Println(p)
		fmt.Println(p.Vertex())
	}
	fmt.Println(p.Vertex())

	return 0
}

func part2() {

}

func main() {
	input := "input_example"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)[0]))

	fmt.Println("== [ PART 2 ] ==")
	part2()
}
