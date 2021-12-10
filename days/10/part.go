package main

import (
	"fmt"
	"sort"

	"github.com/RaphaelPour/aoc2021/util"
)

var (
	parScore = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	parScore2 = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	counterpart = map[string]string{
		"[": "]",
		"(": ")",
		"<": ">",
		"{": "}",
	}
)

type Line struct {
	input     string
	score     int
	tail      string
	corrupted bool
}

func NewLine(input string) *Line {
	return &Line{input: input}
}

func (l *Line) Match(par string) bool {
	if l.Empty() {
		return false
	}

	return string(l.input[0]) == par
}

func (l *Line) Consume() string {
	if l.Empty() {
		return ""
	}
	par := string(l.input[0])
	l.input = l.input[1:]
	return par
}

func (l *Line) Accept(par string) bool {
	if !l.Match(par) {
		return false
	}

	l.Consume()
	return true
}

func (l *Line) Empty() bool {
	return len(l.input) == 0
}

func (l *Line) Reduce() bool {
	if l.Empty() {
		return true
	}

	for _, par := range []string{"(", "[", "{", "<"} {
		if l.Accept(par) {
			if l.Reduce() && l.Accept(counterpart[par]) {
				return l.Reduce()
			}

			if l.Empty() {
				return true
			}
			if l.score == 0 {
				l.score = parScore[string(l.input[0])]
			}
			return false
		}
	}
	return true
}

func (l *Line) Reduce2() bool {
	if l.Empty() {
		return true
	}

	for _, par := range []string{"(", "[", "{", "<"} {
		if l.Accept(par) {
			cpar := counterpart[par]
			if l.Reduce2() && l.Accept(cpar) {
				return l.Reduce2()
			}

			if l.Empty() {
				l.tail += cpar
				l.score = l.score*5 + parScore2[cpar]
				return true
			}

			// corrupted
			return false
		}
	}
	// nothing to reduce, like a closed par
	return true
}

func part1(input []string) int {
	sum := 0
	for _, l := range input {
		line := NewLine(l)
		line.Reduce()
		sum += line.score
	}
	return sum
}

func part2(input []string) int {
	scores := make([]int, 0)
	for _, l := range input {
		line := NewLine(l)
		if line.Reduce2() {
			scores = append(scores, line.score)
		}
	}

	if len(scores) == 0 {
		fmt.Println("Nothing found")
		return -1
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadDefaultString()))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString("input")))
}
