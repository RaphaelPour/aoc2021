package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

type TrenchMap struct {
	enhancement   []bool
	image         [][]bool
	width, height int
}

func NewTrenchMap(input []string) *TrenchMap {
	t := new(TrenchMap)
	t.enhancement = make([]bool, len(input[0]))

	// parse enhancement algorithm line
	for i, symbol := range input[0] {
		t.enhancement[i] = (symbol == rune('#'))
	}

	// skip enhancement + new line
	input = input[2:]

	// parse input image
	t.width = len(input)
	t.height = len(input[0])
	t.image = make([][]bool, t.height)
	for i, row := range input {
		if len(row) != t.width {
			panic(fmt.Sprintf(
				"row width (%d) differs from overall width %d",
				len(row),
				t.width,
			))
		}
		t.image[i] = make([]bool, t.width)
		for j, pixel := range row {
			t.image[i][j] = (pixel == rune('#'))
		}
	}
	return t
}

func (t TrenchMap) Dump() {
	for i, truth := range t.enhancement {
		if i%16 == 0 {
			fmt.Printf("\n0x%03x: ", i)
		}
		if truth {
			fmt.Print("\033[32m#\033[0m")
		} else {
			fmt.Print("\033[31m.\033[0m")
		}
	}
	fmt.Println("")

	for _, row := range t.image {
		for _, pixel := range row {
			if pixel {
				fmt.Print("\033[32m#\033[0m")
			} else {
				fmt.Print("\033[31m.\033[0m")
			}
		}
		fmt.Println("")
	}
}

func (t TrenchMap) LitCount() int {
	count := 0
	for y := 0; y < t.height; y++ {
		for x := 0; x < t.width; x++ {
			if t.image[y][x] {
				count++
			}
		}
	}
	return count
}

func (t TrenchMap) GetPixel(x, y int) bool {
	// out-of-bound = pixel is dark
	if x < 0 || x >= t.width || y < 0 || y >= t.height {
		return false
	}
	return t.image[y][x]
}

func (t TrenchMap) CalcIndex(x, y int) int {
	i := ""
	// go through each cell of a 3x3 grid with x,y as center
	for yGrid := -1; yGrid <= 1; yGrid++ {
		for xGrid := -1; xGrid <= 1; xGrid++ {
			if t.GetPixel(x+xGrid, y+yGrid) {
				i += "1"
			} else {
				i += "0"
			}
		}
	}

	num, err := strconv.ParseInt(i, 2, 64)
	if err != nil {
		panic(fmt.Sprintf("error converting %s", i))
	}

	return int(num)
}

func (t *TrenchMap) Enhance() {
	// the new image gets an extra pixels per side
	newWidth := t.width + 2
	newHeight := t.height + 2
	newImage := make([][]bool, newHeight)

	// go through new image and calculate each pixel
	for y := 0; y < newHeight; y++ {
		newImage[y] = make([]bool, newWidth)
		for x := 0; x < newWidth; x++ {

			// offset old image by -1/-1
			index := t.CalcIndex(x-1, y-1)
			if index >= len(t.enhancement) {
				panic(fmt.Sprintf(
					"enhancement index %d is out of range %d",
					index,
					len(t.enhancement),
				))
			}
			newImage[y][x] = t.enhancement[index]
		}
	}

	t.image = newImage
	t.height = newHeight
	t.width = newWidth
}

func part1(input []string) int {
	t := NewTrenchMap(input)
	fmt.Println(" -- original --")
	t.Dump()
	t.Enhance()
	fmt.Println(" -- enhanced (1) --")
	t.Dump()
	t.Enhance()
	fmt.Println(" -- enhanced (2) --")
	t.Dump()
	return t.LitCount()
}

func part2() {

}

func main() {
	input := "input_example"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))
	fmt.Println("too low: 5288")
	fmt.Println("too high: 5353")

	fmt.Println("== [ PART 2 ] ==")
	part2()
}
