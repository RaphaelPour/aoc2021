package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

func main() {
	input := util.LoadDefaultString()
	fmt.Println(input)
	for _, value := range input {
		num, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("error converting number: %s\n", err)
			return
		}

		fmt.Printf("%d\n", num*num)
	}
}
