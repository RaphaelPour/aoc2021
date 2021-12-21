package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

type Game struct {
	dice                       int
	player1Score, player2Score int
	player1Pos, player2Pos     int
	round                      int
	cache                      []int
}

func NewGame(input []string) *Game {
	pattern := regexp.MustCompile(`^Player \d starting position: (\d+)$`)
	match := pattern.FindStringSubmatch(input[0])
	if len(match) != 2 {
		panic(fmt.Sprintf("error parsing %s", input[0]))
	}

	player1, err := strconv.Atoi(match[1])
	if err != nil {
		panic(fmt.Sprintf("error converting %s", match[1]))
	}

	match = pattern.FindStringSubmatch(input[1])
	if len(match) != 2 {
		panic(fmt.Sprintf("error parsing %s", input[1]))
	}

	player2, err := strconv.Atoi(match[1])
	if err != nil {
		panic(fmt.Sprintf("error converting %s", match[1]))
	}

	cache := make([]int, 7)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				index := (i + j + k) - 3
				cache[index]++
			}
		}
	}

	return &Game{
		player1Pos: player1,
		player2Pos: player2,
		dice:       1,
		cache:      cache,
	}
}

func (g *Game) turn() (int, bool) {
	g.round++
	n := g.dice*3 + 3
	g.dice += 3
	if g.round%2 == 1 {
		g.player1Pos = ((g.player1Pos + n - 1) % 10) + 1
		g.player1Score += g.player1Pos
		if g.player1Score >= 1000 {
			return g.player2Score, true
		}
	} else {
		g.player2Pos = ((g.player2Pos + n - 1) % 10) + 1
		g.player2Score += g.player2Pos
		if g.player2Score >= 1000 {
			return g.player1Score, true
		}
	}

	return 0, false
}

func (g *Game) quantumTurn(currentScore, otherScore, currentPos, otherPos int) (int, int) {
	currentWins, otherWins := 0, 0
	for i := 3; i <= 9; i++ {
		pos := ((currentPos + i - 1) % 10) + 1
		if currentScore+pos >= 21 {
			currentWins += g.cache[i-3]
		} else {
			other, current := g.quantumTurn2(otherScore, currentScore+pos, otherPos, pos)
			currentWins += current * g.cache[i-3]
			otherWins += other * g.cache[i-3]
		}
	}
	return currentWins, otherWins
}

func part1(input []string) int {
	game := NewGame(input)
	for {
		if score, won := game.turn(); won {
			return score * game.round * 3
		}
	}
}

func part2(input []string) int {
	// count of sums that add up to 21
	game := NewGame(input)
	return util.Max(game.quantumTurn2(0, 0, game.player1Pos, game.player2Pos))
}

func main() {
	input := "input"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString(input)))
}
