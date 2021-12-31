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

func Join(in []int) string {
	out := ""
	for i := len(in) - 1; i >= 0; i-- {
		out += fmt.Sprintf("%d", in[i])
	}
	return out
}

func startLoop2() {
	/*
	 *           1111
	 * 01234567890123
	 * _99____19__64_
	 * _99____19__64_
	 * _99____19__64_
	 * _99____19__64_
	 * _99____19__64_
	 * 1 + 2 must be 9,
	 * 7 must be 1
	 * 8 must be 9
	 */
	ctx := context.Background()
	sem := semaphore.NewWeighted(int64(runtime.GOMAXPROCS(0)))
	for w1 := 1; w1 <= 9; w1++ {
		for w2 := 1; w2 <= 9; w2++ {
			for w3 := 1; w3 <= 9; w3++ {
				for w5 := 1; w5 <= 9; w5++ {
					sem.Acquire(ctx, 1)
					go func(w1, w2, w3, w5 int, sem *semaphore.Weighted) {
						defer sem.Release(1)
						for w6 := 1; w6 <= 9; w6++ {
							for w8 := 1; w8 <= 9; w8++ {
								for w9 := 1; w9 <= 9; w9++ {
									for w10 := 1; w10 <= 9; w10++ {
										for w11 := 1; w11 <= 9; w11++ {
											for w12 := 1; w12 <= 9; w12++ {
												for w13 := 1; w13 <= 9; w13++ {
													w := []int{w13, w12, w11, w10, w9, w8, 1, w6, w5, 10 - w3, w3, w2, w1, 1}
													if z := start(w, false); z == 0 {
														fmt.Println(">>", Join(w), "<<")
													}
												}
											}
										}
									}
								}
							}
						}
						fmt.Print(".")
					}(w1, w2, w3, w5, sem)
				}
			}
		}
	}
	sem.Acquire(ctx, int64(runtime.GOMAXPROCS(0)))
}
func startLoop1() {
	/*
	 *           1111
	 * 01234567890123
	 * _99____19__64_
	 * _99____19__64_
	 * _99____19__64_
	 * _99____19__64_
	 * _99____19__64_
	 * 1 + 2 must be 9,
	 * 7 must be 1
	 * 8 must be 9
	 */
	ctx := context.Background()
	sem := semaphore.NewWeighted(int64(runtime.GOMAXPROCS(0)))
	for w0 := 1; w0 <= 9; w0++ {
		for w3 := 1; w3 <= 9; w3++ {
			sem.Acquire(ctx, 1)
			go func(w0, w3 int, sem *semaphore.Weighted) {
				defer sem.Release(1)
				for w4 := 1; w4 <= 9; w4++ {
					for w5 := 1; w5 <= 9; w5++ {
						for w6 := 1; w6 <= 9; w6++ {
							for w9 := 1; w9 <= 9; w9++ {
								for w10 := 1; w10 <= 9; w10++ {
									for w13 := 1; w13 <= 9; w13++ {
										w := []int{w13, 4, 6, w10, w9, 9, 1, w6, w5, w4, w3, 9, 9, w0}
										if z := start(w, false); z == 0 {
											fmt.Println(">>", Join(w), "<<")
											os.Exit(0)
										}
									}
								}
							}
						}
					}
				}
				fmt.Print(".")
			}(w0, w3, sem)
		}
	}
	sem.Acquire(ctx, int64(runtime.GOMAXPROCS(0)))
}

func start(digits []int, verbose bool) int {
	index := 0
	z := 0

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

func bruteForce() {
	chunk := 1000000
	ctx := context.Background()
	sem := semaphore.NewWeighted(int64(runtime.GOMAXPROCS(0)))
	begin, end := 79975391969649, 99999969999982
	for i := begin; i < end; i += chunk {
		sem.Acquire(ctx, 1)
		go func(offset, length int, sem *semaphore.Weighted) {
			defer sem.Release(1)
			for j := offset; j < offset+length; j++ {
				digits := ToIntArray(j)
				/*
				 *           1111
				 * 01234567890123
				 * _99____19__64_
				 * _99____19__64_
				 * _99____19__64_
				 * _99____19__64_
				 * _99____19__64_
				 *
				 * Reversed
				 *           1111
				 * 01234567890123
				 * _46__91____99_
				 * 1 + 2 must be 9,
				 * 7 must be 1
				 * 8 must be 9
				 */

				if digits[1] != 4 || digits[2] != 6 || digits[5] != 9 || digits[6] != 1 || digits[11] != 9 || digits[12] != 9 {
					continue
				}
				if z := start(digits, false); z == 0 {
					fmt.Println(" >> ", j, " <<")
				}
			}
			fmt.Println(offset + length)
		}(i, chunk, sem)
	}
	sem.Acquire(ctx, int64(runtime.GOMAXPROCS(0)))
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
				start(ToIntArray(num), true)
			}
			num := util.ToInt(os.Args[1])
			start(ToIntArray(num), true)
			return
		}
	}

	// bruteForce()
	startLoop2()
}
