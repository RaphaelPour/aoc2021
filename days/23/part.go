package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2021/util"
)

/* TODO
 *
 * - [ ] parse input -> grid or map?
 * - [ ] A* search
 */

const (
	AMBER = iota
	BRONZE
	COPPER
	DESERT
)

var (
	cost     = []int{1, 10, 100, 1000}
	podReMap = []rune{'A', 'B', 'C', 'D'}
	podMap   = map[rune]int{
		'A': AMBER,
		'B': BRONZE,
		'C': COPPER,
		'D': DESERT,
	}
	queues = []int{3, 5, 7, 9}
)

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type State struct {
	pods map[int][]Point
}

func (s State) String() string {
	out := ""
	for i, kind := range podReMap {
		out += fmt.Sprintf("%c%s ", kind, s.pods[i])
	}
	return out
}

func NewState(input []string) *State {
	s := new(State)
	s.pods = make(map[int][]Point)

	for y, line := range input {
		for x := range line {
			if pod, ok := podMap[rune(input[y][x])]; ok {
				if _, ok := s.pods[pod]; !ok {
					s.pods[pod] = make([]Point, 0)
				}
				s.pods[pod] = append(s.pods[pod], Point{x, y})
			}
		}
	}

	fmt.Println(s)

	return s
}

func (s State) Copy() State {
	newState = State{}
	newState.pods = make(map[int][]Point)

	for key, value := range s.pods {
		newState.pods[key] = make([]Point, len(value))
		copy(newState.pods[key], value)
	}
}

func (s State) GoalReached() bool {
	for kind := range podReMap {
		for _, pod := range s.pods[kind] {
			if pod.y < 3 {
				/* pod is still in hallway */
				return false
			}

			if pod.x != queues[kind] {
				/* pod is not inside the queue */
				return false
			}

		}
	}

	return true
}

func (s State) Next() <-chan State {
	ch := make(chan State)
	go func() {
		for kind := range podReMap {
			for _, pod := range s.pods[kind] {
				newState = State{}
				newState.pods = make(map[int][]Point)

			}
		}
	}()
	return ch
}

func part1(input []string) int {
	s := NewState(input)
	return len(s.pods)
}

func part2(input []string) int {
	return len(input)
}

func main() {
	input := "input_example"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part1(util.LoadString(input)))
	fmt.Println("too high: 42756")
	fmt.Println("bad: 42088, 42588 ,42678, 42786")
	fmt.Println("min: 40284")
}
