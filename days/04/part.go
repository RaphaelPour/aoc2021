package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RaphaelPour/aoc2021/util"
)

type Field struct {
	number int
	marked bool
}

func (f Field) Dump() {
	if f.marked {
		fmt.Printf("\033[32m%2d \033[0m", f.number)
	} else {
		fmt.Printf("\033[31m%2d \033[0m", f.number)
	}
}

type Board struct {
	gameField [][]Field
}

func (b Board) String() string {
	out := ""
	for _, row := range b.gameField {
		out += fmt.Sprintf("%2d", row[0].number)
		for _, field := range row[1:len(row)] {
			out += fmt.Sprintf(" %2d", field.number)
		}
		out += "\n"
	}

	return out
}

func (b Board) Dump() {
	for _, row := range b.gameField {
		for _, field := range row {
			field.Dump()
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func (b *Board) Mark(num int) {
	for i := range b.gameField {
		for j := range b.gameField[i] {
			if b.gameField[i][j].number == num {
				b.gameField[i][j].marked = true
			}
		}
	}
}

func (b Board) IsBingo() bool {
	for i := range b.gameField {
		rowBingo := true
		colBingo := true
		for j := range b.gameField[i] {
			if !b.gameField[i][j].marked {
				rowBingo = false
			}

			if !b.gameField[j][i].marked {
				colBingo = false
			}
		}
		if rowBingo || colBingo {
			return true
		}
	}
	return false
}

func (b Board) UnmarkedSum() int {
	var sum int
	for i := range b.gameField {
		for j := range b.gameField[i] {
			if !b.gameField[i][j].marked {
				sum += b.gameField[i][j].number
			}
		}
	}

	return sum
}

type Boards []*Board

func (b Boards) String() string {
	out := ""
	for i, board := range b {
		out += fmt.Sprintf("== Board %02d ==\n%s\n", i+1, board)
	}
	return out
}

func (b Boards) Dump() {
	for i, board := range b {
		fmt.Printf("== Board %02d ==\n", i+1)
		board.Dump()
		fmt.Println("")
	}
}

func (b Boards) Mark(num int) {
	for _, board := range b {
		board.Mark(num)
	}
}

func (b Boards) CheckBingo() *Board {
	for _, board := range b {
		if board.IsBingo() {
			return board
		}
	}

	return nil
}

func LoadBingoInput(file string) (Boards, []int) {
	input := util.LoadString(file)

	// parse turns
	turns := make([]int, 0)
	for _, turn := range strings.Split(input[0], ",") {
		num, err := strconv.Atoi(turn)
		if err != nil {
			panic(fmt.Sprintf("error converting '%s': %s", turn, err))
		}
		turns = append(turns, num)
	}

	// parse boards

	boards := make(Boards, 0)
	var currentBoard *Board
	for _, rawRow := range input[1:len(input)] {
		if rawRow == "" {
			if currentBoard != nil {
				boards = append(boards, currentBoard)
			}
			currentBoard = new(Board)
			currentBoard.gameField = make([][]Field, 0)
			continue
		}

		row := make([]Field, 0)
		for _, field := range strings.Split(rawRow, " ") {
			if field == "" {
				continue
			}

			num, err := strconv.Atoi(field)
			if err != nil {
				panic(fmt.Sprintf("error converting field '%s': %s", field, err))
			}
			row = append(row, Field{number: num})
		}
		currentBoard.gameField = append(currentBoard.gameField, row)
	}

	// add last board
	boards = append(boards, currentBoard)
	return boards, turns
}

func part1(file string) int {
	boards, turns := LoadBingoInput("input_example")

	for _, turn := range turns {
		boards.Mark(turn)
		if winner := boards.CheckBingo(); winner != nil {
			return winner.UnmarkedSum() * turn
		}
	}

	return -1
}

func part2(file string) int {
	boards, turns := LoadBingoInput(file)

	bingoed := make([]bool, len(boards))
	for _, turn := range turns {
		fmt.Printf("turn: %2d\n", turn)
		boards.Mark(turn)

		for i, board := range boards {
			// ignore boards that have already won
			if bingoed[i] {
				continue
			}

			if board.IsBingo() {
				fmt.Println(bingoed)
				bingoed[i] = true

				bingoCount := 0
				for _, bingo := range bingoed {
					if bingo {
						bingoCount++
					}
				}

				board.Dump()
				sum := board.UnmarkedSum()
				fmt.Printf("%d*%d = %d\n", sum, turn, sum*turn)
				// is there only one board left?
				if bingoCount == len(bingoed) {
					return board.UnmarkedSum() * turn
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1("input_example"))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2("input_example"))
	fmt.Println("Too low: 2884")
	fmt.Println("Too high: 3710")
}
