package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2021/util"
)

var (
	parScore = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	counterpart = map[string]string{
		"[": "]",
		"(": ")",
		"<": ">",
		"{": "}",
	}
)

type Line struct {
	input string
	score int
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
	// fmt.Printf("reducing %s\n", l.input)
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
				l.score = -1
				fmt.Printf("%s: %s\n", par, l.input)
				if !l.Empty() {
					s := parScore[string(l.input[0])]
					fmt.Printf(
						"expected %s, got %s instead, adding %d to score\n",
						par,
						l.input,
						s,
					)
					l.score = s
				}
			}
			return false
		}
	}
	return true
}

func part1(input []string) int {
	sum := 0
	for _, l := range input {
		line := NewLine(l)
		// fmt.Println(line)
		ok := line.Reduce()
		fmt.Println(ok, line.score)
		sum += line.score
	}
	return sum
}

func part2() {

}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadDefaultString()))
	fmt.Println("bad answer: 268854, 288240")
	fmt.Println("too high  : 487878")

	fmt.Println("== [ PART 2 ] ==")
	part2()
}
