package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/RaphaelPour/aoc2021/util"
	"golang.org/x/sync/semaphore"
)

var (
	params = [][]int{
		{1, 14, 8},
		{1, 13, 8},
		{1, 13, 3},
		{1, 12, 10},
		{26, -12, 8},
		{1, 12, 8},
		{26, -2, 8},
		{26, -11, 5},
		{1, 13, 9},
		{1, 14, 3},
		{26, 0, 4},
		{26, -12, 9},
		{26, -13, 2},
		{26, -6, 7},
	}
)

func ToIntArray(input int) []int {
	result := make([]int, 0)
	for input > 0 {
		result = append(result, input%10)
		input /= 10
	}
	return result
}

func start(input int, verbose bool) int {
	index := 0
	z := 0
	digits := ToIntArray(input)
	for i := len(digits) - 1; i >= 0; i-- {
		w := digits[i]
		if w == 0 {
			// skip if any digit is zero
			return -1
		}
		z = reduce(
			w,
			z,
			params[index][0],
			params[index][1],
			params[index][2],
		)
		if verbose {
			fmt.Printf("%2d -> %d\n", w, z)
		}
		index++
	}
	if verbose {
		fmt.Println("")
	}
	return z
}

func reduce(w, z, p1, p2, p3 int) int {
	x := (z % 26) + p2
	z /= p1
	if x != w {
		x = 1
	} else {
		x = 0
	}

	y := x*25 + 1
	z *= y
	y = (w + p3) * x
	z += y

	return z
}

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "shell" {
			reader := bufio.NewReader(os.Stdin)
			for {
				fmt.Print(">")
				input, err := reader.ReadString('\n')
				if err != nil {
					fmt.Printf("error: %s\n", err)
					return
				}
				// trim new line at the end
				input = strings.TrimSpace(input)
				if len(input) == 0 {
					continue
				}

				if input == "q" {
					return
				}
				num, err := strconv.Atoi(input)
				if err != nil {
					fmt.Printf("unknown input '%s'\n", input)
					continue
				}
				start(num, true)
			}
			num := util.ToInt(os.Args[1])
			start(num, true)
			return
		}
	}
	chunk := 10000000
	ctx := context.Background()
	sem := semaphore.NewWeighted(int64(runtime.GOMAXPROCS(0)))
	for i := 19975399999999; i >= 19975311111111; i -= chunk {
		sem.Acquire(ctx, 1)
		go func(offset, length int, sem *semaphore.Weighted) {
			defer sem.Release(1)
			for j := offset; j > offset-length; j-- {
				if z := start(j, false); z == 0 {
					fmt.Println(" >> ", j, " <<")
				}
			}
			fmt.Println(offset - length + 1)
		}(i, chunk, sem)
	}
	sem.Acquire(ctx, int64(runtime.GOMAXPROCS(0)))
}
