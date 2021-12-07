package main

import (
	"fmt"

	"github.com/RaphaelPour/aoc2021/util"
)

type CostFunction func(int, int) int

func linearDistance(target, start int) int {
	return util.Abs(target - start)
}
func exponentialDistance(target, start int) int {
	distance := util.Abs(target - start)
	return (distance*distance + distance) / 2
}

func part(input []int, costFunction CostFunction) int {
	costs := make([]int, 0)
	min, max := util.MinMax(input...)

	for i := min; i < max; i++ {
		cost := 0
		for _, pos := range input {
			cost += costFunction(i, pos)
		}
		costs = append(costs, cost)
	}

	return util.Min(costs...)
}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part(
		util.LoadIntList("input"),
		linearDistance,
	))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part(
		util.LoadIntList("input"),
		exponentialDistance,
	))
}
