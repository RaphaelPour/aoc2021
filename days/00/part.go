package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2021/util"
)

func count(input []string) int {
	result := 0
	for _, line := range input {
		for _, chr := range line {
			if "<" == string(chr) {
				result++
			}
		}
	}

	return result
}

func part1() {
	// input := util.LoadString("input")
	// input := util.LoadDefaultInt()
	// input := util.LoadInt("input")
	input := util.LoadDefaultString()
	fmt.Println(count(input))
}

func part2() {

}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	part1()

	fmt.Println("== [ PART 2 ] ==")
	part2()
}
