package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2021/util"
)

const (
	EMPTY = iota
	EAST
	SOUTH
)

var (
	runeMap = map[rune]int{
		'.': EMPTY,
		'>': EAST,
		'v': SOUTH,
	}
	enumMap = []string{".", ">", "v"}
)

type SeaFloor struct {
	fields [][]int
}

func NewSeaFloor(input []string) *SeaFloor {
	s := new(SeaFloor)
	s.fields = make([][]int, len(input))

	for i, line := range input {
		s.fields[i] = make([]int, len(line))
		for j, field := range line {
			field, ok := runeMap[rune(field)]
			if !ok {
				panic(fmt.Sprintf("unknown field %s", string(field)))
			}
			s.fields[i][j] = field
		}
	}
	return s
}

func (s *SeaFloor) MoveLoop() int {
	var round int
	anyMove := true
	for anyMove {
		newFields := make([][]int, len(s.fields))
		anyMove = false

		// process east
		for y := range s.fields {
			newFields[y] = make([]int, len(s.fields[y]))
			for x := range s.fields[y] {
				if s.fields[y][x] != EAST {
					continue
				}
				nextX := x + 1
				if nextX >= len(s.fields[y]) {
					nextX = 0
				}

				// check if neighbor
				if s.fields[y][nextX] == EMPTY {
					newFields[y][nextX] = EAST
					anyMove = true
				} else {
					newFields[y][x] = EAST
				}
			}
		}

		// process south
		for y := range s.fields {
			for x := range s.fields[y] {
				if s.fields[y][x] != SOUTH {
					continue
				}
				nextY := y + 1
				if nextY >= len(s.fields) {
					nextY = 0
				}

				// check if neighbor
				if s.fields[nextY][x] != SOUTH && newFields[nextY][x] != EAST {
					newFields[nextY][x] = SOUTH
					anyMove = true
				} else {
					newFields[y][x] = SOUTH
				}
			}
		}
		s.fields = newFields
		round++
	}
	return round
}

func (s SeaFloor) Dump() {
	for y := range s.fields {
		for x := range s.fields[y] {
			if s.fields[y][x] == EMPTY {
				fmt.Printf("\033[32m.\033[0m")
			} else {
				fmt.Printf("\033[31m%s\033[0m", enumMap[s.fields[y][x]])
			}
		}
		fmt.Println("")
	}
}

func part1(input []string) int {
	s := NewSeaFloor(input)
	return s.MoveLoop()
}

func main() {
	input := "input"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))
}
