package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/RaphaelPour/aoc2021/util"
)

type RuleSet struct {
	start         string
	rules         map[string]string
	realCacheHits int
}

func NewRuleSet(input []string) *RuleSet {
	r := new(RuleSet)
	r.start = input[0]
	r.rules = make(map[string]string)
	for _, rawRule := range input[2:] {
		parts := strings.Split(rawRule, " -> ")
		if len(parts) != 2 {
			panic(fmt.Sprintf(
				"error parsing %s, expected length 2 but got %d",
				rawRule, len(parts),
			))
		}
		r.rules[parts[0]] = parts[1]
	}
	return r
}

func (rs *RuleSet) SubstituteOnce() {
	newStart := string(rs.start[0])
	for i := 0; i < len(rs.start)-1; i++ {
		premise := string(rs.start[i]) + string(rs.start[i+1])
		conclusion, ok := rs.rules[premise]
		if !ok {
			panic(fmt.Sprintf("no rule for premise '%s'", premise))
		}
		newStart += conclusion + string(rs.start[i+1])
	}
	fmt.Printf("%s -> %s\n", rs.start, newStart)
	rs.start = newStart
}

func (rs *RuleSet) SubstituteFast() {
	newStart := string(rs.start[0])
	oldStart := rs.start
	for len(oldStart) >= 2 {

		// try to find great element strings in cache to boost performance
		// on cache miss, decrease length down to two, the minimal element
		// substitution size
		for i := len(oldStart); i >= 2; i-- {
			// try to find current premise in cache
			conclusion, ok := rs.rules[string(oldStart[:i])]
			if !ok {
				// cache miss...
				continue
			}
			// CACHE HIT!
			if len(string(oldStart[:i])) > 2 {
				rs.realCacheHits++

				fmt.Printf("cache hit: %s -> %s\n", string(oldStart[:i]), conclusion)
				/*
					fmt.Printf(
						"           expand output: %s -> %s\n",
						newStart, newStart+conclusion+string(oldStart[1]),
					)
					fmt.Printf(
						"            reduce input: %s -> %s\n",
						oldStart,
						string(oldStart[i-1:]),
					)*/
			}
			newStart += conclusion + string(oldStart[1])
			oldStart = string(oldStart[i-1:])
			break
		}
	}

	// fmt.Printf("%s -> %s\n", rs.start, newStart)

	// cache result for later

	/*
		for window := 4; window < len(rs.start); window += 2 {
			for i := 0; i < len(rs.start); i += window {
				newPremise := rs.start[i : i+window]
				if _, ok := rs.rules[newPremise]; ok {
					// fmt.Printf("rule for %s already existing\n", rs.start[:i])
					continue
				}
				newConclusion := newStart[window/2 : (i+window)-window/2]
				rs.rules[newPremise] = newConclusion
				fmt.Printf("New rule %s -> %s\n", newPremise, newConclusion)
			}
		}*/
	// fmt.Printf("|rules| = %d\n", len(rs.rules))

	rs.start = newStart
}

func (rs RuleSet) ElementStats() (int, int) {
	hist := make(map[string]int)
	for _, element := range rs.start {
		hist[string(element)] = hist[string(element)] + 1
	}

	min, max := 10000, 0
	for _, sum := range hist {
		if sum < min {
			min = sum
		}

		if sum > max {
			max = sum
		}
	}
	return min, max
}

func (rs RuleSet) Result() int {
	min, max := rs.ElementStats()
	return max - min
}

func part1(input []string) int {
	ruleSet := NewRuleSet(input)

	for round := 0; round < 7; round++ {
		start := time.Now()
		ruleSet.SubstituteOnce()

		fmt.Printf(
			"round %d in %s with result %d\n",
			round,
			time.Since(start),
			ruleSet.Result(),
		)
	}

	min, max := ruleSet.ElementStats()
	fmt.Printf("min= %d, max=%d\n", min, max)
	return max - min
}

func part2(input []string) int {
	ruleSet := NewRuleSet(input)
	ruleSetTest := NewRuleSet(input)

	for round := 0; round < 7; round++ {
		start := time.Now()
		ruleSet.SubstituteFast()
		ruleSetTest.SubstituteOnce()

		// fmt.Printf("round %d in %s with output %s\n", round, time.Since(start), ruleSet.start)
		fmt.Printf(
			"round %d in %s with result %d and %d real cache hits\n",
			round,
			time.Since(start),
			ruleSet.Result(),
			ruleSet.realCacheHits,
		)
		if ruleSet.start != ruleSetTest.start {
			fmt.Printf("start differs:\n  valid: %s\ninvalid: %s\n", ruleSetTest.start, ruleSet.start)
			break
		}
	}

	min, max := ruleSet.ElementStats()
	fmt.Printf("min= %d, max=%d\n", min, max)
	return max - min
}

func main() {
	input := "input_example"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString(input)))
}
