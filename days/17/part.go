package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

type Vector struct {
	x, y int
}

func (v *Vector) Add(other Vector) {
	v.x += other.x
	v.y += other.y
}

func (v Vector) String() string {
	return fmt.Sprintf("(%d %d)", v.x, v.y)
}

type Rectangle struct {
	v1, v2 Vector
}

func (r Rectangle) String() string {
	return fmt.Sprintf("%s -> %s", r.v1, r.v2)
}

func (r Rectangle) Contains(v Vector) bool {
	return v.x >= r.v1.x && v.x <= r.v2.x &&
		v.y >= r.v1.y && v.y <= r.v2.y
}

func (r Rectangle) Reachable(v Vector) bool {
	return v.y >= r.v1.y
}

func trickshot(velocity, start Vector, target Rectangle) (int, bool) {
	peak := 0

	// check if target is not reachable anymore
	for target.Reachable(start) {
		// fmt.Println(start, velocity)
		// check if at target
		if target.Contains(start) {
			return peak, true
		}

		start.Add(velocity)
		peak = util.Max(peak, start.y)

		// apply drag
		if velocity.x > 0 {
			velocity.x--
		} else if velocity.x < 0 {
			velocity.x++
		}

		// apply gravity
		velocity.y--
	}

	return peak, false
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

	target := Rectangle{
		v1: Vector{x1, y1},
		v2: Vector{x2, y2},
	}

	maxPeak := 0
	velocities := 0
	for y := -1000; y <= 1000; y++ {
		for x := -1000; x <= 1000; x++ {
			if peak, ok := trickshot(Vector{x, y}, Vector{0, 0}, target); ok {
				velocities++
				if peak > maxPeak {
					maxPeak = peak
				}
			}
		}
	}

	fmt.Println("velocities:", velocities)
	return maxPeak
}

func part2() {

}

func main() {
	input := "input"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)[0]))
	fmt.Println("too low: 5050")

	fmt.Println("== [ PART 2 ] ==")
	part2()
}
