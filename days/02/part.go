package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RaphaelPour/aoc2021/util"
)

func part1(input []string) int {

	h, d := 0, 0

	for _, line := range input {
		tmp := strings.Split(line, " ")
		if len(tmp) != 2 {
			fmt.Printf("error parsing '%s'\n", line)
			return -1
		}

		value, err := strconv.Atoi(tmp[1])
		if err != nil {
			fmt.Printf("error parsing value '%s': %s\n", tmp[1], err)
			return -1
		}

		switch tmp[0] {
		case "forward":
			h += value
		case "down":
			d += value
		case "up":
			d -= value
		}
	}

	return h * d
}

func part2(input []string) int {
	h, d := 0, 0
	aim := 0

	for _, line := range input {
		tmp := strings.Split(line, " ")
		if len(tmp) != 2 {
			fmt.Printf("error parsing '%s'\n", line)
			return -1
		}

		value, err := strconv.Atoi(tmp[1])
		if err != nil {
			fmt.Printf("error parsing value '%s': %s\n", tmp[1], err)
			return -1
		}

		switch tmp[0] {
		case "forward":
			h += value
			d += value * aim
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	return h * d
}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadDefaultString()))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadDefaultString()))
}
