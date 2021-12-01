package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2021/util"
)

func part1(input []int) int {
	last := input[0]
	result := 0
	for _, current := range input[1:len(input)] {
		if current > last {
			result++
		}
		last = current
	}

	// good: 1342
	return result
}

func part2(input []int) int {
	result := 0
	lastSum := 0
	windowSize := 3
	for i := range input {
		sum := 0
		for j := i; j < i+windowSize && j < len(input); j++ {
			sum += input[j]
		}

		/* skip if it's the first sum */
		if sum > lastSum && lastSum > 0 {
			result++
		}
		if i+2 >= len(input) {
			break
		}
		lastSum = sum
	}

	// good: 1378
	return result
}

func main() {
	input := util.LoadDefaultInt()

	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(input))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(input))
}
