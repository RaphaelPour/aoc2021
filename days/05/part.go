package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

var (
	pattern = regexp.MustCompile(`^(\d+),(\d+)\s->\s(\d+),(\d+)$`)
)

type SeaMap map[int]map[int]int

func NewSeaMap(input []string, skipDiagonal bool) (SeaMap, error) {
	seaMap := make(SeaMap)

	for _, line := range input {
		match := pattern.FindStringSubmatch(line)
		if len(match) != 5 {
			panic(fmt.Sprintf("regex mismatch with '%s'", line))
		}

		fromX, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, fmt.Errorf(
				"error converting number '%s': %w",
				match[1],
				err,
			)
		}
		fromY, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, fmt.Errorf(
				"error converting number '%s': %w",
				match[2],
				err,
			)
		}

		toX, err := strconv.Atoi(match[3])
		if err != nil {
			return nil, fmt.Errorf(
				"error converting number '%s': %w",
				match[3],
				err,
			)
		}
		toY, err := strconv.Atoi(match[4])
		if err != nil {
			return nil, fmt.Errorf(
				"error converting number '%s': %w",
				match[4],
				err,
			)
		}

		seaMap.MarkLine(fromX, fromY, toX, toY, skipDiagonal)
	}

	return seaMap, nil
}

func (s SeaMap) MarkPoint(x, y int) {
	if _, ok := s[y]; !ok {
		s[y] = make(map[int]int)
	}
	s[y][x]++
}

func (s SeaMap) MarkLine(x1, y1, x2, y2 int, skipDiagonal bool) {
	if x1 == x2 {
		start, end := y1, y2
		if start > end {
			start, end = end, start
		}
		for y := start; y <= end; y++ {
			s.MarkPoint(x1, y)
		}
	} else if y1 == y2 {
		start, end := x1, x2
		if start > end {
			start, end = end, start
		}
		for x := start; x <= end; x++ {
			s.MarkPoint(x, y1)
		}
	} else {
		if skipDiagonal {
			return
		}

		for y := 0; y <= util.Abs(y1-y2); y++ {
			for x := 0; x <= util.Abs(x1-x2); x++ {
				if x == y {
					s.MarkPoint(x2+x*util.Sign(x1-x2), y2+y*util.Sign(y1-y2))
				}
			}
		}
	}
}

func (s SeaMap) Get(x, y int) int {
	if _, ok := s[y]; !ok {
		return 0
	}

	return s[y][x]
}

func (s SeaMap) Overlapping() int {
	var overlap int
	for y := range s {
		for x := range s[y] {
			if s[y][x] > 1 {
				overlap++
			}
		}
	}
	return overlap
}

func (s SeaMap) Dump() {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if _, ok := s[y]; ok && s[y][x] > 0 {
				fmt.Printf("%d", s[y][x])
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
}

func part1(input []string) int {
	seaMap, err := NewSeaMap(input, true)
	if err != nil {
		panic(err.Error())
	}

	return seaMap.Overlapping()
}

func part2(input []string) int {
	seaMap, err := NewSeaMap(input, false)
	if err != nil {
		panic(err.Error())
	}

	return seaMap.Overlapping()
}

func main() {
	input := util.LoadDefaultString()

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(input))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(input))
}
