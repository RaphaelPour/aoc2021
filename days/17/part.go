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

type Launcher struct {
	target Rectangle
}

func NewLauncher(input string) *Launcher {
	p := new(Launcher)

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

	p.target = Rectangle{
		v1: Vector{x1, y1},
		v2: Vector{x2, y2},
	}

	return p
}

func (p *Launcher) probe() (int, int) {
	maxPeak := 0
	velocities := 0
	for y := p.target.v1.y; y <= 1000; y++ {
		for x := 1; x <= p.target.v2.x; x++ {
			if peak, ok := p.shoot(Vector{x, y}); ok {
				velocities++
				if peak > maxPeak {
					maxPeak = peak
				}
			}
		}
	}
	return maxPeak, velocities
}

func (p *Launcher) shoot(velocity Vector) (int, bool) {
	peak := 0
	start := Vector{0, 0}

	// check if target is not reachable anymore
	for p.target.Reachable(start) {
		// check if at target
		if p.target.Contains(start) {
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
	maxPeak, _ := NewLauncher(input).probe()
	return maxPeak
}

func part2(input string) int {
	_, velocities := NewLauncher(input).probe()
	return velocities
}

func main() {
	input := "input"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)[0]))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString(input)[0]))
}
