package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2021/util"
)

const (
	WALL = iota
	FREE
	AMBER
	BRONZE
	COPPER
	DESERT
	VOID
)

var (
	constMap = map[string]int{
		"A": AMBER,
		"B": BRONZE,
		"C": COPPER,
		"D": DESERT,
		" ": VOID,
		"#": WALL,
		".": FREE,
	}
	stringMap = []string{"#", ".", "A", "B", "C", "D", " "}
	energyMap = map[int]int{
		AMBER:  1,
		BRONZE: 10,
		COPPER: 100,
		DESERT: 1000,
	}
	goalLowMap = map[int]Point{
		AMBER:  Point{3, 2},
		BRONZE: Point{5, 2},
		COPPER: Point{7, 2},
		DESERT: Point{9, 2},
	}
	goalHighMap = map[int]Point{
		AMBER:  Point{3, 3},
		BRONZE: Point{5, 3},
		COPPER: Point{7, 3},
		DESERT: Point{9, 3},
	}
)

func isAmphipod(kind int) bool {
	return kind >= AMBER && kind <= DESERT
}

type Point struct {
	x, y int
}

type Amphipod struct {
	kind              int
	pos               Point
	goalLow, goalHigh Point
}

func (a Amphipod) GoalReached() bool {
	return a.pos == a.goalLow || a.pos == a.goalHigh
}

func (a Amphipod) Move(x, y int) (Amphipod, int) {
	a.pos.x += x
	a.pos.y += y
	return a, (x + y) * energyMap[a.kind]
}

func (a Amphipod) String() string {
	return fmt.Sprintf(
		"%s at %s with goals [%s,%s]",
		stringMap[a.kind],
		a.pos,
		a.goalLow,
		a.goalHigh,
	)
}

type Amphipods []Amphipod

func (a Amphipods) GoalReached() bool {
	return a[0].GoalReached() && a[1].GoalReached()
}

func (a Amphipods) String() string {
	return fmt.Sprintf("[%s, %s]", a[0], a[1])
}

func (p Point) String() string {
	return fmt.Sprintf("%d|%d", p.x, p.y)
}

type Situation struct {
	fields    [][]int
	amphipods map[int]Amphipods
	cache     map[CacheConfig]bool
}

type CacheConfig struct {
	a1, a2, b1, b2, c1, c2, d1, d2 Amphipod
}

func (s *Situation) StartSearch() int {
	ok := false
	cost := 0
	for limit := 1; !ok; limit++ {
		s.cache = make(map[CacheConfig]bool)

		cost, ok = s.Search(
			s.amphipods[AMBER][0], s.amphipods[AMBER][1],
			s.amphipods[BRONZE][0], s.amphipods[BRONZE][1],
			s.amphipods[COPPER][0], s.amphipods[COPPER][1],
			s.amphipods[DESERT][0], s.amphipods[DESERT][1],
			0, limit,
		)
		fmt.Println(limit, ":", cost)
	}
	return cost
}

func (s *Situation) Search(a1, a2, b1, b2, c1, c2, d1, d2 Amphipod, energy, limit int) (int, bool) {
	// abort if energy limit has been reached
	if energy > limit {
		fmt.Println("LIMIT: ", energy, ">", limit)
		return 0, false
	}

	if a1.GoalReached() && a2.GoalReached() &&
		b1.GoalReached() && b2.GoalReached() &&
		c1.GoalReached() && c2.GoalReached() &&
		d1.GoalReached() && d2.GoalReached() {
		fmt.Println("GOAL", energy)
		return energy, true
	}

	min := 1000000000000000
	anyOk := false
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}

			n, cost := a1.Move(x, y)
			if _, ok := s.cache[CacheConfig{n, a2, b1, b2, c1, c2, d1, d2}]; ok {
				return 0, false
			}
			s.cache[CacheConfig{n, a2, b1, b2, c1, c2, d1, d2}] = true
			if e, ok := s.Search(n, a2, b1, b2, c1, c2, d1, d2, energy+cost, limit); ok && e < min {
				min = e
				anyOk = true
			}

			n, cost = a2.Move(x, y)
			if _, ok := s.cache[CacheConfig{a1, n, b1, b2, c1, c2, d1, d2}]; ok {
				return 0, false
			}
			s.cache[CacheConfig{a1, n, b1, b2, c1, c2, d1, d2}] = true
			if e, ok := s.Search(a1, n, b1, b2, c1, c2, d1, d2, energy+cost, limit); ok && e < min {
				min = e
				anyOk = true
			}

			n, cost = b1.Move(x, y)
			if _, ok := s.cache[CacheConfig{a1, a2, n, b2, c1, c2, d1, d2}]; ok {
				return 0, false
			}
			s.cache[CacheConfig{a1, a2, n, b2, c1, c2, d1, d2}] = true
			if e, ok := s.Search(a1, a2, n, b2, c1, c2, d1, d2, energy+cost, limit); ok && e < min {
				min = e
				anyOk = true
			}
			n, cost = b2.Move(x, y)
			if _, ok := s.cache[CacheConfig{a1, a2, b1, n, c1, c2, d1, d2}]; ok {
				return 0, false
			}
			s.cache[CacheConfig{a1, a2, b1, n, c1, c2, d1, d2}] = true
			if e, ok := s.Search(a1, a2, b1, n, c1, c2, d1, d2, energy+cost, limit); ok && e < min {
				min = e
				anyOk = true
			}
			n, cost = c1.Move(x, y)
			if _, ok := s.cache[CacheConfig{a1, a2, b1, b2, n, c2, d1, d2}]; ok {
				return 0, false
			}
			s.cache[CacheConfig{a1, a2, b1, b2, n, c2, d1, d2}] = true
			if e, ok := s.Search(a1, a2, b1, b2, n, c2, d1, d2, energy+cost, limit); ok && e < min {
				min = e
				anyOk = true
			}
			n, cost = c2.Move(x, y)
			if _, ok := s.cache[CacheConfig{a1, a2, b1, b2, c1, n, d1, d2}]; ok {
				return 0, false
			}
			s.cache[CacheConfig{a1, a2, b1, b2, c1, n, d1, d2}] = true
			if e, ok := s.Search(a1, a2, b1, b2, c1, n, d1, d2, energy+cost, limit); ok && e < min {
				min = e
				anyOk = true
			}
			n, cost = d1.Move(x, y)
			if _, ok := s.cache[CacheConfig{a1, a2, b1, b2, c1, c2, n, d2}]; ok {
				return 0, false
			}
			s.cache[CacheConfig{a1, a2, b1, b2, c1, c2, n, d2}] = true
			if e, ok := s.Search(a1, a2, b1, b2, c1, c2, n, d2, energy+cost, limit); ok && e < min {
				min = e
				anyOk = true
			}
			n, cost = d2.Move(x, y)
			s.cache[CacheConfig{a1, a2, b1, b2, c1, c2, d1, n}] = true
			if _, ok := s.cache[CacheConfig{a1, a2, b1, b2, c1, c2, d1, n}]; ok {
				return 0, false
			}
			if e, ok := s.Search(a1, a2, b1, b2, c1, c2, d1, n, energy+cost, limit); ok && e < min {
				min = e
				anyOk = true
			}
		}
	}

	return min, anyOk
}

func NewSituation(input []string) *Situation {
	s := new(Situation)
	s.fields = make([][]int, len(input))
	s.amphipods = make(map[int]Amphipods)
	s.cache = make(map[CacheConfig]bool)
	for i := range input {
		s.fields[i] = make([]int, len(input[0]))
		for j := range input[i] {
			field, ok := constMap[string(input[i][j])]
			if !ok {
				panic(fmt.Sprintf("unknown char '%s'\n", input[i][j]))
			}

			// convert void into wall to simplify map
			if field == VOID {
				field = WALL
			}

			if isAmphipod(field) {
				if _, ok := s.amphipods[field]; !ok {
					s.amphipods[field] = make(Amphipods, 0)
				}
				s.amphipods[field] = append(s.amphipods[field], Amphipod{field, Point{j, i}, goalLowMap[field], goalHighMap[field]})
			}
			// now that we stored the amphipods, let's turn their fields to free ones
			if field != WALL {
				field = FREE
			}
			s.fields[i][j] = field
		}
	}
	return s
}

func (s Situation) Dump(withAmphipods bool) {
	amphMap := make(map[Point]string)
	for kind, pods := range s.amphipods {
		sKind := stringMap[kind]
		amphMap[pods[0].pos] = sKind
		amphMap[pods[1].pos] = sKind

		fmt.Println(pods)
	}

	for y := range s.fields {
		for x := range s.fields[y] {
			if withAmphipods {
				if kind, ok := amphMap[Point{x, y}]; ok {
					fmt.Print(kind)
					continue
				}
			}
			fmt.Print(stringMap[s.fields[y][x]])
		}
		fmt.Println("")
	}

}

func part1(input []string) int {
	s := NewSituation(input)
	s.Dump(true)
	return s.StartSearch()
}

func part2() {

}

func main() {
	input := "input_example"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	part2()
}
