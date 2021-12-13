package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/RaphaelPour/aoc2021/util"
)

type Fold struct {
	axis   string
	offset int
}

type Paper struct {
	fields map[int]map[int]bool
	folds  []Fold
}

func NewPaper(input []string, firstFoldOnly bool) *Paper {
	p := new(Paper)
	p.fields = make(map[int]map[int]bool)

	// parse dots
	i := 0
	for _, row := range input {
		// break on empty line
		if row == "" {
			break
		}

		parts := strings.Split(row, ",")
		if len(parts) != 2 {
			panic(fmt.Sprintf("expected two parts, got %d with %s", len(parts), row))
		}
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(fmt.Sprintf("error parsing %s", parts[0]))
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(fmt.Sprintf("error parsing %s", parts[1]))
		}

		if _, ok := p.fields[y]; !ok {
			p.fields[y] = make(map[int]bool)
		}

		p.fields[y][x] = true
		i++
	}

	// parse folds
	foldPattern := regexp.MustCompile(`^fold along (x|y)=(\d+)$`)
	p.folds = make([]Fold, 0)
	for _, row := range input[i+1:] {
		match := foldPattern.FindStringSubmatch(row)
		if len(match) != 3 {
			panic(fmt.Sprintf("error parsing %s, expected three matches, got %s", row, match))
		}

		num, err := strconv.Atoi(match[2])
		if err != nil {
			panic(fmt.Sprintf("error parsing %s", match[2]))
		}

		if match[1] != "x" && match[1] != "y" {
			panic(fmt.Sprintf("unknown fold axis %s", match[1]))
		}
		p.folds = append(p.folds, Fold{axis: match[1], offset: num})

		if firstFoldOnly {
			break
		}
	}

	return p
}

func (p *Paper) Fold() {

}

func (p *Paper) DotCount() int {
	sum := 0
	for _, row := range p.fields {
		for _, cell := range row {
			if cell {
				sum++
			}
		}
	}
	return sum
}

func part1(input []string) int {
	p := NewPaper(input, true)
	p.Fold()
	return p.DotCount()
}

func part2() {

}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadDefaultString()))

	fmt.Println("== [ PART 2 ] ==")
	part2()
}
