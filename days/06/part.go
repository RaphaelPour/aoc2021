package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RaphaelPour/aoc2021/util"
)

func part1(input string, days int) int {

	rawFishes := strings.Split(input, ",")
	fishes := make([]int, len(rawFishes))
	for i, fish := range rawFishes {
		num, err := strconv.Atoi(fish)
		if err != nil {
			panic(err)
		}

		fishes[i] = num
	}

	for day := 0; day < days; day++ {
		for i := range fishes {
			if fishes[i] > 0 {
				fishes[i]--
			} else {
				fishes = append(fishes, 8)
				fishes[i] = 6
			}
		}
	}

	return len(fishes)
}

func calc(fish, days int) []int {
	fishes := make([]int, 1)
	fishes[0] = fish
	for day := 0; day < days; day++ {
		for i := range fishes {
			if fishes[i] > 0 {
				fishes[i]--
			} else {
				fishes = append(fishes, 8)
				fishes[i] = 6
			}
		}
	}

	return fishes
}

func part2(input string, days int) int {
	rawFishes := strings.Split(input, ",")
	fishes := make([]int, len(rawFishes))
	cache := make(map[int][]int, 0)
	for i, fish := range rawFishes {
		num, err := strconv.Atoi(fish)
		if err != nil {
			panic(err)
		}
		fishes[i] = num

		if _, ok := cache[num]; !ok {
			cache[num] = calc(num, 128)
		}
	}

	fmt.Println("Done caching")

	for i := 0; i < 2; i++ {
		fmt.Printf("round %d with size=%d\n", i, len(fishes))
		for j, num := range fishes {
			if _, ok := cache[num]; !ok {
				fmt.Printf("caching %d\n", num)
				cache[num] = calc(num, 64)
			}
			fishes = append(fishes[:j], fishes[j+1:]...)
			fishes = append(fishes, cache[num]...)
		}
	}

	return len(fishes)
}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString("input")[0], 80))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString("input")[0], 256))
}
