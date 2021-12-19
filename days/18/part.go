package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/RaphaelPour/aoc2021/util"
)

type SnailfishNumber struct {
	root  *Node
	input string
}

type Node struct {
	left, right, parent *Node
	literal             bool
	number              int
}

func (n *Node) Reset() {
	n.left = nil
	n.right = nil
	n.parent = nil
	n.literal = true
	n.number = 0
}

func (n *Node) LeftLiteral() (*Node, bool) {
	// TODO: Check if final left literal node is
	//       actually one of the caller pair.

	// if literal, go one up since a pair consists of two nodes rather
	// than one with two numbers.
	if n.literal {
		n = n.parent.left
	}

	iterator := n
	// go up until left is different to where we came from.
	// if literal is already very left, this loop will go
	// until root has been reached
	for iterator.parent != nil && iterator.parent.left == iterator {
		iterator = iterator.parent
	}

	// early return if root has been reached (iterator is nil)
	if iterator.parent == nil {
		return nil, false
	}

	iterator = iterator.parent.left

	// the iterator still remains our literal node. In this case
	// the caller is the left most.
	if iterator == n {
		return nil, false
	}

	// otherwise we have reached the (grand*)parent node that has a  non-nil left
	// side. Go to the right most node, which is the left neighbor of the caller
	// node.
	for !iterator.literal {
		iterator = iterator.right
	}

	// GOTCHA!
	return iterator, true
}

func (n *Node) RightLiteral() (*Node, bool) {
	// if literal, go one up since a pair consists of two nodes rather
	// than one with two numbers.
	if n.literal {
		n = n.parent.right
	}

	iterator := n
	// go up until right is different to where we came from.
	// if literal is already very right, this loop will go
	// until root has been reached
	for iterator.parent != nil && iterator.parent.right == iterator {
		iterator = iterator.parent
	}

	// early return if root has been reached (iterator is nil)
	if iterator.parent == nil {
		return nil, false
	}

	iterator = iterator.parent.right

	// the iterator still remains our literal node. In this case
	// the caller is the right most.
	if iterator == n {
		return nil, false
	}

	// otherwise we have reached the (grand*)parent node that has a  non-nil right
	// side. Go to the left most node, which is the right neighbor of the caller
	// node.
	for !iterator.literal {
		iterator = iterator.left
	}

	// GOTCHA!
	return iterator, true
}

func (n *Node) Split() bool {
	if n.literal && n.number >= 10 {
		// create left child, with literal floored half
		left := new(Node)
		left.literal = true
		left.number = n.number / 2

		// create right child, with literal ceiled half
		right := new(Node)
		right.literal = true
		right.number = (n.number / 2) + (n.number % 2)

		// adopt child
		n.literal = false
		n.number = 0
		n.left = left
		n.right = right
		n.parent = n

		return true
	}

	// if node is literal and not grater 10, return
	if n.literal {
		return false
	}

	if !n.left.literal && n.left.Split() {
		return true
	}

	if !n.right.literal && n.right.Split() {
		return true
	}

	return false
}

func (n *Node) Explode(depth int) bool {
	// if node is literal we can't do much since the explode
	// itself is always done by the literal's parent
	if n.literal {
		return false
	}

	if depth >= 4 && n.left.literal && n.right.literal {
		// TODO: find left and right literals via parent
		n.Reset()
		return true
	}

	if !n.left.literal && n.left.Explode(depth+1) {
		return true
	}

	if !n.right.literal && n.right.Explode(depth+1) {
		return true
	}

	return false
}

func (n Node) String() string {
	if n.literal {
		return fmt.Sprintf("%d", n.number)
	}
	return fmt.Sprintf("[%s,%s]", n.left, n.right)
}

func NewSnailfishNumbers(input []string) []SnailfishNumber {
	nodes := make([]SnailfishNumber, 0)
	for _, line := range input {
		s := SnailfishNumber{input: line}
		s.root = new(Node)
		s.parse(s.root)
		nodes = append(nodes, s)
	}

	return nodes
}

func (s SnailfishNumber) String() string {
	return s.root.String()
}

func (s SnailfishNumber) next() string {
	if len(s.input) == 0 {
		return ""
	}

	return string(s.input[0])
}

func (s *SnailfishNumber) consume() string {
	if len(s.input) == 0 {
		return ""
	}

	symbol := string(s.input[0])
	s.input = s.input[1:]
	return symbol
}

func (s *SnailfishNumber) accept(expectation string) bool {
	if s.next() == expectation {
		s.input = s.input[1:]
		return true
	}
	return false
}

func (s *SnailfishNumber) acceptNumber() (int, bool) {
	if len(s.input) == 0 {
		return -1, false
	}

	num, err := strconv.Atoi(string(s.input[0]))
	if err != nil {
		return -1, false
	}

	s.input = s.input[1:]
	return num, true
}

func (s *SnailfishNumber) parse(node *Node) {
	if len(s.input) == 0 {
		panic("parse error, expected something but input is empty")
	}

	// check if leaf, and leave early with a literal node
	if num, ok := s.acceptNumber(); ok {
		node.literal = true
		node.number = num
		return
	}

	// otherwise it must fulfill [<NODE>,<NODE>]
	if !s.accept("[") {
		panic(fmt.Sprintf("syntax error: expected '[' but got '%s'", s.next()))
	}

	/* create left+right nodes such before hand and set neighbors
	 * so on parsing them, they can to the same for their children
	 *
	 *            (NODE)
	 *            /   \
	 *         (LEFT) (RIGHT)
	 *
	 */

	left := new(Node)
	left.parent = node

	right := new(Node)
	right.parent = node

	node.left = left
	node.right = right

	s.parse(left)

	if !s.accept(",") {
		panic(fmt.Sprintf("syntax error: expected ',' but got '%s'", s.next()))
	}

	s.parse(right)

	if !s.accept("]") {
		panic(fmt.Sprintf("syntax error: expected ']' but got '%s'", s.next()))
	}
}

func (s *SnailfishNumber) Add(other SnailfishNumber) {
	newRoot := new(Node)
	newRoot.left = s.root
	newRoot.right = other.root

	s.root = newRoot

	fmt.Printf("after addition: %s\n", s.root)
	// reduce until numbers stay the same
	for {
		// explode everything before continuing
		for {
			if !s.root.Explode(0) {
				break
			}
			fmt.Printf("after explode:  %s\n", s.root)
		}
		// split
		if !s.root.Split() {
			break
		}
		fmt.Printf("after split:    %s\n", s.root)
	}
}

func part1(input []string) int {
	s := NewSnailfishNumbers(input)
	sum := s[0]
	fmt.Printf("befire:         %s\n", sum)
	for _, num := range s[1:] {
		sum.Add(num)
		fmt.Println(sum)
	}
	return 0
}

func part2() {

}

func evalInput(input string) (SnailfishNumber, bool) {
	ok := true
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
			ok = false
		}
	}()

	s := SnailfishNumber{input: input}
	s.root = new(Node)
	s.parse(s.root)
	fmt.Println(s.root)
	return s, ok
}

func addNumbers(a, b *SnailfishNumber) {
	ok := true
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
			ok = false
		}
	}()
	fmt.Printf(" %s\n", a)
	fmt.Printf("+%s\n", b)
	a.Add(*b)
	if ok {
		fmt.Printf("=%s\n", a)
	}
}

func explodeNumber(a SnailfishNumber) {
	ok := true
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
			ok = false
		}
	}()
	a.root.Explode(1)
	if ok {
		fmt.Println(a.root)
	}
}

func shell() {
	reader := bufio.NewReader(os.Stdin)
	numbers := make([]SnailfishNumber, 0)
	for {
		fmt.Print("@y>")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}
		// trim new line at the end
		input = strings.TrimSpace(input)
		split := strings.Fields(input)
		if len(split) == 0 {
			continue
		}
		command := split[0]

		switch command {
		case "exit", "q":
			return
		case "clear", "c":
			numbers = numbers[0:0]
		case "add", "a":
			if len(split) < 3 {
				fmt.Println("argument a,b missing")
				continue
			}

			a, err := strconv.Atoi(split[1])
			if err != nil {
				fmt.Printf("error converting '%s'\n", split[1])
				continue
			}
			b, err := strconv.Atoi(split[2])
			if err != nil {
				fmt.Printf("error converting '%s'\n", split[2])
				continue
			}

			if a >= len(numbers) {
				fmt.Printf("index %d out of range %d\n", a, len(numbers))
				continue
			}

			if b >= len(numbers) {
				fmt.Printf("index %d out of range %d\n", b, len(numbers))
				continue
			}

			addNumbers(&numbers[a], &numbers[b])
		case "explode", "e":
			if len(split) < 2 {
				fmt.Println("argument a missing")
				continue
			}

			a, err := strconv.Atoi(split[1])
			if err != nil {
				fmt.Printf("error converting '%s'\n", split[1])
				continue
			}
			if a >= len(numbers) {
				fmt.Printf("index %d out of range %d\n", a, len(numbers))
				continue
			}
			explodeNumber(numbers[a])
		case "load", "l":
			if len(split) < 2 {
				fmt.Println("argument file missing")
				continue
			}

			if _, err := os.Stat(split[1]); errors.Is(err, os.ErrNotExist) {
				fmt.Printf("file '%s' does't exist\n", split[1])
				continue
			}
			for _, line := range util.LoadString(split[1]) {
				if num, ok := evalInput(line); ok {
					numbers = append(numbers, num)
				}
			}
		case "save", "s":
			if len(split) < 2 {
				fmt.Println("argument file missing")
				continue
			}

			if _, err := os.Stat(split[1]); err == nil {
				fmt.Printf("file '%s' already exists\n", split[1])
				continue
			}

			file, err := os.Create(split[1])
			if err != nil {
				fmt.Printf("error opening file '%s': %s\n", split[1], err)
				continue
			}

			for _, num := range numbers {
				file.WriteString(num.root.String() + "\n")
			}

			file.Close()

		case "print", "p":
			// pourman's log_10, -1 since index 9 shoudn't be padded like 10
			l := len(fmt.Sprintf("%d", len(numbers)-1))
			for i, num := range numbers {
				fmt.Printf("%*d: %s\n", l, i, num.root)
			}
		case "help", "h", "?":
			fmt.Println("  print           print all numbers")
			fmt.Println("    add  <a> <b>  add two numbers where a,b are the")
			fmt.Println("explode  <a>      explode number")
			fmt.Println("                  indices based on the loaded set")
			fmt.Println("  clear           clear all numbers")
			fmt.Println("   load  <file>   load input file")
			fmt.Println("   save  <file>   save input file")
			fmt.Println("   exit           exit shell")
			fmt.Println("   help           show this help")
			fmt.Println(" <else>           parse input as snail number and add")
			fmt.Println("                 it to the other numbers on success")
		default:
			if num, ok := evalInput(input); ok {
				numbers = append(numbers, num)
			}
		}
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "shell" {
		shell()
		return
	}

	input := "input_example3"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)))

	fmt.Println("== [ PART 2 ] ==")
	part2()
}
