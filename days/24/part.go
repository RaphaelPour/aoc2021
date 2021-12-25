package main

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/RaphaelPour/aoc2021/util"
	"golang.org/x/sync/semaphore"
)

const (
	INPUT = iota
	ADD
	MULTIPLY
	DIVIDE
	MOD
	EQUAL
)

var (
	instructionMap = map[string]int{
		"inp": INPUT,
		"add": ADD,
		"mul": MULTIPLY,
		"div": DIVIDE,
		"mod": MOD,
		"eql": EQUAL,
	}

	enumMap = []string{"inp", "add", "mul", "div", "mod", "eql"}
)

type Instruction struct {
	command  int
	op1, op2 string
	value    int
}

func (i Instruction) String() string {
	com := enumMap[i.command]
	if i.command == INPUT {
		return fmt.Sprintf("com=%s op1=%s", com, i.op1)
	}

	if i.op2 != "" {
		return fmt.Sprintf("com=%s op1=%s op2=%s", com, i.op1, i.op2)
	}
	return fmt.Sprintf("com=%s op1=%s val=%d", com, i.op1, i.value)
}

type ALU struct {
	program   []Instruction
	variables map[string]int
}

func NewALU(input []string) *ALU {
	a := new(ALU)
	a.program = make([]Instruction, len(input))
	for i, line := range input {
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			panic(fmt.Sprintf("expected at least 2 parts, got %d: %s", len(parts), line))
		}

		ins := Instruction{}

		// parse instruction
		var ok bool
		ins.command, ok = instructionMap[parts[0]]
		if !ok {
			panic(fmt.Sprintf("unknown instruction %s", parts[0]))
		}

		// parse first operand, always a variable
		ins.op1 = parts[1]

		if len(parts) == 2 {
			a.program[i] = ins
			continue
		}

		// parse second operand
		if val, err := strconv.Atoi(parts[2]); err == nil {
			ins.value = val
		} else {
			ins.op2 = parts[2]
		}

		a.program[i] = ins
	}
	a.ResetRAM()
	return a
}

func (a *ALU) ResetRAM() {
	a.variables = map[string]int{
		"x": 0,
		"y": 0,
		"z": 0,
		"w": 0,
	}
}

func (a ALU) getValue(ins Instruction) int {
	if ins.op2 == "" {
		return ins.value
	}
	if val, ok := a.variables[ins.op2]; ok {
		return val
	}
	panic(fmt.Sprintf("unknown variable %s", ins.op2))
}

func (a *ALU) Run(rawInput int) (int, bool) {
	a.ResetRAM()
	input := strconv.Itoa(rawInput)
	for _, ins := range a.program {
		switch ins.command {
		case INPUT:
			if len(input) == 0 {
				return 0, false
			}
			num, err := strconv.Atoi(string(input[0]))
			if err != nil {
				panic(fmt.Sprintf("error converting %s", string(input[0])))
			}
			a.variables[ins.op1] = num
			if num == 0 {
				// skip zero digits
				return 0, false
			}
			if len(input) > 0 {
				input = input[1:len(input)]
			}
		case ADD:
			a.variables[ins.op1] += a.getValue(ins)
		case MULTIPLY:
			a.variables[ins.op1] *= a.getValue(ins)
		case DIVIDE:
			op2 := a.getValue(ins)
			if op2 == 0 {
				panic("division by zero")
			}
			a.variables[ins.op1] = int(float64(a.variables[ins.op1]) / float64(op2))
		case MOD:
			op2 := a.getValue(ins)
			if op2 == 0 {
				panic("division by zero")
			}
			a.variables[ins.op1] = a.variables[ins.op1] % op2
		case EQUAL:
			if a.variables[ins.op1] == a.getValue(ins) {
				a.variables[ins.op1] = 1
			} else {
				a.variables[ins.op1] = 0
			}
		default:
			panic(fmt.Sprintf("Unknown instruction %d\n", ins.command))
		}
	}
	return a.variables["z"], true
}

func compute(start, length int, input []string, sem *semaphore.Weighted) {
	defer sem.Release(1)
	a := NewALU(input)
	for i := start; i > start-length; i-- {
		if _, ok := a.Run(i); ok {
			fmt.Println(i)
			break
		}
		if i%1000000 == 0 {
			fmt.Print(".")
		}
	}
}

func parallel(input []string) {
	start := 99999999999999
	chunkSize := 10000000

	ctx := context.Background()
	sem := semaphore.NewWeighted(int64(runtime.GOMAXPROCS(0)))
	for offset := start; offset > 0; offset -= chunkSize {
		if err := sem.Acquire(ctx, 1); err != nil {
			panic("error acquiring semaphore")
		}
		go compute(offset, chunkSize, input, sem)
	}

	sem.Acquire(ctx, int64(runtime.GOMAXPROCS(0)))
}

func (a ALU) analyze(digit, base int) (int, int) {
	min, max := 100000000000000000, 0
	minDigit, maxDigit := 0, 0
	for i := 1; i <= 9; i++ {
		a.Run(i*util.Pow(10, digit) + base)
		z := a.variables["z"]

		if z < min {
			min = z
			minDigit = i
		}
		if z > max {
			max = z
			maxDigit = i
		}
		fmt.Println(i, a.variables)
	}

	return minDigit, maxDigit
}

func part1(input []string) int {
	a := NewALU(input)
	number := 0
	for digit := 0; digit < 14; digit++ {
		_, max := a.analyze(digit, number)
		number += max * util.Pow(10, digit)
		fmt.Println(number)
	}
	return 0
}

func part2() {

}

func main() {
	input := "input"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))
	fmt.Println("too high: 99999999901273")

	fmt.Println("== [ PART 2 ] ==")
	part2()
}
