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
	pairMap       map[string]int
}

func NewRuleSet(input []string) *RuleSet {
	r := new(RuleSet)
	r.start = input[0]
	r.rules = make(map[string]string)
	r.pairMap = make(map[string]int)
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

	for i := 0; i < len(r.start)-1; i++ {
		pair := r.start[i : i+2]
		r.pairMap[pair] = 1
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
	// fmt.Printf("%s -> %s\n", rs.start, newStart)
	rs.start = newStart
}

func (rs *RuleSet) SubstituteFast2() {
	newPairs := make(map[string]int)
	for premise, conclusion := range rs.rules {
		if _, ok := rs.pairMap[premise]; !ok {
			// rule doesn't match
			continue
		}

		pair1 := string(premise[0]) + conclusion
		pair2 := conclusion + string(premise[1])

		newPairs[pair1] = newPairs[pair1] + rs.pairMap[premise]
		if pair1 != pair2 {
			newPairs[pair2] = newPairs[pair2] + rs.pairMap[premise]
		}
	}

	rs.pairMap = newPairs
}

func (rs *RuleSet) SubstituteFast(expandCache bool) {
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
			}
			newStart += conclusion + string(oldStart[len(conclusion)/2+1])
			oldStart = string(oldStart[i-1:])
			break
		}
	}

	// fmt.Printf("%s -> %s\n", rs.start, newStart)

	// cache result for later
	if expandCache {
		rs.EnrichCache(rs.start, newStart)
	}
	rs.start = newStart
}

func (rs *RuleSet) EnrichCache(original, substitute string) {
	// for each window size
	for windowSize := 4; windowSize <= len(original)/2; windowSize += 2 {
		// for each window in input
		for windowOffset := 0; windowOffset < len(original)-windowSize+1; windowOffset++ {
			premise := original[windowOffset : windowOffset+windowSize]
			if _, ok := rs.rules[premise]; ok {
				continue
			}

			conclusionIndex := windowOffset * 2
			conclusionSize := (2*windowSize - 1)
			conclusion := substitute[conclusionIndex+1 : conclusionIndex+conclusionSize-1]
			rs.rules[premise] = conclusion
		}
	}
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

	for round := 0; round < 10; round++ {
		start := time.Now()
		ruleSet.SubstituteOnce()

		fmt.Printf(
			"round %d in %s with result %d\n",
			round+1,
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

	for round := 0; round < 40; round++ {
		start := time.Now()
		ruleSet.SubstituteFast2()

		// fmt.Printf("round %d in %s with output %s\n", round, time.Since(start), ruleSet.start)
		fmt.Printf(
			"round %d in %s\n",
			round+1,
			time.Since(start),
		)
	}

	hist := make(map[string]int)
	hist[string(ruleSet.start[0])] = hist[string(ruleSet.start[0])] + 1
	for pair, sum := range ruleSet.pairMap {
		// hist[string(pair[0])] = hist[string(pair[0])] + sum
		hist[string(pair[1])] = hist[string(pair[1])] + sum
	}

	min, max := hist[string(ruleSet.start[0])], hist[string(ruleSet.start[0])]
	for _, sum := range hist {
		if sum > max {
			max = sum
		}
		if sum < min {
			min = sum
		}
	}

	fmt.Printf("min=%d, max=%d\n", min, max)
	return max - min
}

func main() {
	input := "input"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString(input)))
	fmt.Println("too high: 4288726369828")
	fmt.Println("bad: 3277772741531")
}
