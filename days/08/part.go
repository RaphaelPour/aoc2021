package main

import (
	"fmt"
	"strings"

	"github.com/RaphaelPour/aoc2021/util"
)

const (
	SEGMENT_INDEX = 0
	OUTPUT_INDEX  = 1
)

type Display struct {
	Segments    []string
	Output      []string
	Connections map[string]int
	Numbers     map[int]string
}

func NewDisplay(input string) (*Display, error) {
	d := new(Display)
	parts := strings.Split(input, "|")

	if len(parts) != 2 {
		return nil, fmt.Errorf("expecting two parts, got %d", len(parts))
	}

	d.Segments = strings.Split(strings.TrimSpace(parts[SEGMENT_INDEX]), " ")
	d.Output = strings.Split(strings.TrimSpace(parts[OUTPUT_INDEX]), " ")
	d.Connections = make(map[string]int)
	d.Numbers = make(map[int]string)
	return d, nil
}

func (d Display) Rewire() {
	for _, out := range d.Segments {
		length := len(out)
		if length == 2 {
			d.Connections[out] = 1
			d.Numbers[1] = out
		} else if length == 4 {
			d.Connections[out] = 4
			d.Numbers[4] = out
		} else if length == 3 {
			d.Connections[out] = 7
			d.Numbers[7] = out
		} else if length == 7 {
			d.Connections[out] = 8
			d.Numbers[8] = out
		}
	}

	for _, out := range d.Segments {
		length := len(out)
		if _, ok := d.Numbers[9]; !ok && length == 6 && sub(d.Numbers[4], out) == 2 {
			d.Connections[out] = 9
			d.Numbers[9] = out
			break
		}
	}

	for _, out := range d.Segments {
		length := len(out)
		if _, ok := d.Numbers[6]; !ok && length == 6 && d.Numbers[9] != out && sub(out, d.Numbers[1]) == 6 {
			d.Connections[out] = 6
			d.Numbers[6] = out
		} else if _, ok := d.Numbers[3]; !ok && length == 5 && within(d.Numbers[1], out) {
			d.Connections[out] = 3
			d.Numbers[3] = out
		}
	}

	for _, out := range d.Segments {
		length := len(out)
		if _, ok := d.Numbers[0]; !ok && length == 6 && d.Numbers[9] != out && d.Numbers[6] != out {
			d.Connections[out] = 0
			d.Numbers[0] = out
		} else if _, ok := d.Numbers[2]; !ok && length == 5 && d.Numbers[3] != out && !within(out, d.Numbers[9]) {
			d.Connections[out] = 2
			d.Numbers[2] = out
		} else if _, ok := d.Numbers[5]; !ok && length == 5 && d.Numbers[3] != out && within(out, d.Numbers[6]) {
			d.Connections[out] = 5
			d.Numbers[5] = out
		}
	}

}

func (d Display) Result() (int, error) {
	result := 0
	for i, out := range d.Output {
		num := -1
		for con, val := range d.Connections {
			if sub(con, out) == 0 {
				num = val
				break
			}
		}
		if num == -1 {
			return -1, fmt.Errorf("%s missing in %#v", out, d.Connections)
		}

		result += num * util.Pow(10, len(d.Output)-i-1)
	}
	return result, nil
}

func (d Display) Dump() {
	fmt.Printf(" Output: %#v\n", d.Output)
	fmt.Printf("Numbers: %#v\n", d.Numbers)
	fmt.Printf("Connections: %#v\n", d.Connections)
}

func part1(input []string) int {
	sum := 0
	for _, line := range input {
		display, err := NewDisplay(line)
		if err != nil {
			panic(err)
		}
		for _, out := range display.Output {
			length := len(out)
			if length == 2 || length == 4 || length == 3 || length == 7 {
				sum++
			}
		}
	}
	return sum
}

func within(needle, haystack string) bool {
	for j := 0; j < len(needle); j++ {
		found := false
		for i := 0; i < len(haystack); i++ {
			if needle[j] == haystack[i] {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

func sub(a, b string) int {
	hist := make(map[rune]int)
	for _, c := range a {
		hist[c] = hist[c] + 1
	}

	for _, c := range b {
		hist[c] = hist[c] + 1
	}

	diff := 0
	for _, sum := range hist {
		if sum != 2 {
			diff++
		}
	}

	return diff
}

func part2(input []string) int {
	sum := 0
	for _, line := range input {
		display, err := NewDisplay(line)
		if err != nil {
			panic(err)
		}

		display.Rewire()
		result, err := display.Result()
		if err != nil {
			panic(err)
		}
		sum += result
	}

	return sum
}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadDefaultString()))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadDefaultString()))
}
