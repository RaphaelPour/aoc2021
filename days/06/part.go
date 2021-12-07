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
	cache := make(map[int]int)

	for _, fish := range rawFishes {
		num, err := strconv.Atoi(fish)
		if err != nil {
			panic(err)
		}

		value := cache[num]
		value++
		cache[num] = value
	}

	for day := 0; day < days; day++ {
		newCache := make(map[int]int)
		for age, sum := range cache {
			if age == 0 {
				newCache[8] = sum + newCache[8]
				newCache[6] = sum + newCache[6]
				continue
			}

			newCache[age-1] = sum + newCache[age-1]
		}
		cache = newCache
	}

	total := 0
	for _, sum := range cache {
		total += sum
	}
	return total
}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString("input")[0], 80))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString("input")[0], 256))
}
