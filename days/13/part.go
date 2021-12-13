package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/RaphaelPour/aoc2021/util"
)

const (
	NO_AXIS = -1
)

type Fold struct {
	axis   string
	offset int
}

func (f Fold) String() string {
	return fmt.Sprintf("axis=%s offset=%d", f.axis, f.offset)
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

func (p Paper) Width() int {
	max := 0
	for y := range p.fields {
		for x := range p.fields[y] {
			if x > max {
				max = x
			}
		}
	}
	return max + 1 + (max % 2)
}

func (p Paper) Height() int {
	max := 0
	for y := range p.fields {
		if y > max {
			max = y
		}
	}

	return max + 1 + (max % 2)
}

func (p Paper) Output() string {
	height := p.Height()
	width := p.Width()
	out := ""

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if p.fields[y][x] {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	return out
}

func (p Paper) Dump(axisX, axisY int) {
	height := p.Height()
	width := p.Width()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if axisX == x {
				fmt.Printf("\033[33m|\033[0m")
			} else if axisY == y {
				fmt.Printf("\033[33m-\033[0m")
			} else if p.fields[y][x] {
				fmt.Printf("\033[32m#\033[0m")
			} else {
				fmt.Printf("\033[31m.\033[0m")
			}
		}
		fmt.Println("")
	}
}

func (p *Paper) Fold() {
	// apply each fold in the order of the fold list
	for _, fold := range p.folds {
		// for each index from 0 to the axis
		for i := 0; i < fold.offset; i++ {
			// e.g. width=10 and i=0, the index needs to be 9
			mirrorI := (2 * fold.offset) - i

			// differ between x axis (x coord is constant) and y axis (y const)
			if fold.axis == "x" {

				// red flag that fold offset and mirrorI are calculated
				// correctly
				if mirrorI == fold.offset {
					break
				}
				// fold x axis
				for y := range p.fields {
					if _, ok := p.fields[y][mirrorI]; ok {
						p.fields[y][i] = true
						delete(p.fields[y], mirrorI)
					}
				}
			} else {
				if _, ok := p.fields[mirrorI]; !ok {
					continue
				}

				// fold y axis, go through all fields on the bottom and top
				// half separately
				for x := range p.fields[mirrorI] {
					// there could be mirrored y lines at the bottom half
					// that have no point at the top half yet.
					if _, ok := p.fields[i]; !ok {
						p.fields[i] = make(map[int]bool)
					}

					p.fields[i][x] = true
				}

				// delete whole line at the bottom half
				delete(p.fields, mirrorI)
			}
		}

		// remove fold axis
		if fold.axis == "x" {
			for y := range p.fields {
				delete(p.fields[y], fold.offset)
			}
		} else {
			delete(p.fields, fold.offset)
		}
	}
}

func (p *Paper) DotCount() int {
	sum := 0
	for _, row := range p.fields {
		sum += len(row)
	}
	return sum
}

func part1(input []string) int {
	p := NewPaper(input, true)
	p.Fold()
	return p.DotCount()
}

func part2(input []string) string {
	p := NewPaper(input, false)
	p.Fold()
	return p.Output()
}

func main() {
	input := "input"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString(input)))
}
