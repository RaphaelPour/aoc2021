package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2021/util"
)

/* TODO
 *
 * - [ ] parse input -> grid or map?
 * - [ ] A* search
 */

func part1(input []string) int {

}

func part2(input []string) int {
}

func main() {
	input := "input_example"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	// fmt.Println(part1(util.LoadString(input)))
	fmt.Println("too high: 42756")
	fmt.Println("bad: 42678, 42588")
	fmt.Println("min: 40284")
}
