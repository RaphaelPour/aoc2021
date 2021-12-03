package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

type histogram struct {
	input  []string
	ones   []int
	zeroes []int
}

func (h histogram) numbers() (int, int) {
	number := ""
	antagonNumber := ""
	for i := range h.ones {
		if h.ones[i] > h.zeroes[i] {
			number += "1"
			antagonNumber += "0"
		} else {
			number += "0"
			antagonNumber += "1"
		}
	}

	i, err := strconv.ParseInt(number, 2, 32)
	if err != nil {
		panic(fmt.Sprintf("error converting '%s': %s", number, err))
	}
	i2, err := strconv.ParseInt(antagonNumber, 2, 32)
	if err != nil {
		panic(fmt.Sprintf("error converting '%s': %s", number, err))
	}

	return int(i), int(i2)
}

func (h histogram) powerConsumption() int {
	x, y := h.numbers()
	return x * y
}

func (h histogram) oxygen() int {
	numbersLeft := h.input
	for i := range h.ones {
		numbers := make([]string, 0)
		for _, num := range numbersLeft {
			criteria := "0"

			ones, zeroes := qhist(numbersLeft)
			if ones[i] >= zeroes[i] {
				criteria = "1"
			}
			if criteria == string(num[i]) {
				numbers = append(numbers, num)
			}
		}
		numbersLeft = numbers
		if len(numbersLeft) == 1 {
			break
		}
	}

	if len(numbersLeft) != 1 {
		fmt.Printf(
			"expected numbers to have only one left, instead %d\n",
			len(numbersLeft),
		)
		return -1
	}

	i, err := strconv.ParseInt(numbersLeft[0], 2, 32)
	if err != nil {
		panic(fmt.Sprintf("error converting '%s': %s", numbersLeft[0], err))
	}
	return int(i)
}

func (h histogram) carbon() int {
	numbersLeft := h.input
	for i := range h.ones {
		numbers := make([]string, 0)
		for _, num := range numbersLeft {
			criteria := "0"

			ones, zeroes := qhist(numbersLeft)
			fmt.Printf("  ones: %d\n", ones[i])
			fmt.Printf("zeroes: %d\n", zeroes[i])
			if ones[i] < zeroes[i] {
				criteria = "1"
			}
			if criteria == string(num[i]) {
				numbers = append(numbers, num)
			}

			fmt.Println(numbersLeft)
		}
		numbersLeft = numbers
		if len(numbersLeft) == 1 {
			break
		}
	}

	if len(numbersLeft) != 1 {
		fmt.Printf(
			"expected numbers to have only one left, instead %d\n",
			len(numbersLeft),
		)
		return -1
	}

	i, err := strconv.ParseInt(numbersLeft[0], 2, 32)
	if err != nil {
		panic(fmt.Sprintf("error converting '%s': %s", numbersLeft[0], err))
	}
	return int(i)
}

func qhist(input []string) ([]int, []int) {
	ones := make([]int, len(input[0]))
	zeroes := make([]int, len(input[0]))

	for _, number := range input {
		for i, digit := range number {
			if string(digit) == "1" {
				ones[i]++
			} else {
				zeroes[i]++
			}
		}
	}
	return ones, zeroes
}

func NewHistogram(input []string) *histogram {
	hist := &histogram{input: input}
	hist.ones = make([]int, len(input[0]))
	hist.zeroes = make([]int, len(input[0]))

	for _, number := range input {
		for i, digit := range number {
			if string(digit) == "1" {
				hist.ones[i]++
			} else {
				hist.zeroes[i]++
			}
		}
	}
	return hist
}

func part1() {
	hist := NewHistogram(util.LoadString("input"))
	fmt.Println(hist.powerConsumption())
}

func part2() {
	hist := NewHistogram(util.LoadString("input"))
	fmt.Println(hist.oxygen() * hist.carbon())
}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	part1()

	fmt.Println("== [ PART 2 ] ==")
	part2()
}
