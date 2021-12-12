package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/RaphaelPour/aoc2021/util"
)

type Caves struct {
	paths map[string][]string
}

func NewCaves(input []string) *Caves {
	c := new(Caves)
	c.paths = make(map[string][]string)

	// go through all paths and parse them separately
	for _, path := range input {

		// a path is formatted like <from>-<to>
		nodes := strings.Split(path, "-")
		if len(nodes) != 2 {
			panic(fmt.Sprintf("unexpected path '%s'", path))
		}

		from, to := nodes[0], nodes[1]

		if _, ok := c.paths[from]; !ok {
			c.paths[from] = make([]string, 0)
		}
		c.paths[from] = append(c.paths[from], to)

		// store reverse path since cavern map is unidirectional
		if _, ok := c.paths[to]; !ok {
			c.paths[to] = make([]string, 0)
		}
		c.paths[to] = append(c.paths[to], from)
	}

	return c
}

func (c Caves) pathExamine(start string, visited string, smallCave string, path string) int {
	path += "-" + start
	if unicode.IsLower(rune(start[0])) {
		visited += start + ","
	}
	if start == "end" {
		return 1
	}

	paths := 0
	for _, neighbor := range c.paths[start] {
		if strings.Contains(visited, neighbor) {
			if smallCave != "" || neighbor == "start" {
				continue
			}
			paths += c.pathExamine(neighbor, visited, neighbor, path)
			continue
		}

		paths += c.pathExamine(neighbor, visited, smallCave, path)
	}
	return paths
}

func part1(input []string) int {
	c := NewCaves(input)
	return c.pathExamine("start", "", "NOPE", "")
}

func part2(input []string) int {
	c := NewCaves(input)
	return c.pathExamine("start", "", "", "")
}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString("input")))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString("input")))
	fmt.Println("bad: 20806")
}
