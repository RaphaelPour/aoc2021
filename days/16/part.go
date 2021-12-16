package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

const (
	INFINITY = int(^uint(0) >> 1)

	SUM_PACKET     = 0
	PRODUCT_PACKET = 1
	MINIMUM_PACKET = 2
	MAXIMUM_PACKET = 3
	LITERAL_PACKET = 4
	GREATER_PACKET = 5
	LESSER_PACKET  = 6
	EQUAL_PACKET   = 7
)

type Decoder struct {
	input      string
	index      int
	versionSum int
}

func NewDecoder(input string) *Decoder {
	converted := ""
	for _, hexDigit := range input {
		num, err := strconv.ParseInt(string(hexDigit), 16, 64)
		if err != nil {
			panic(fmt.Sprintf("error converting %s from hex to dec", string(hexDigit)))
		}

		converted += fmt.Sprintf("%04b", num)
	}
	return &Decoder{input: converted}
}

func (d *Decoder) Read(bits int) int {
	if bits == 0 {
		return -1
	}
	if len(d.input) < d.index+bits {
		panic(fmt.Sprintf(
			"eof reached after %d bytes, wanted: %d",
			len(d.input),
			d.index+bits,
		))
	}
	s := d.input[d.index : d.index+bits]
	d.index += bits

	num, err := strconv.ParseInt(string(s), 2, 64)
	if err != nil {
		panic(fmt.Sprintf("error converting %s from bin to dec", string(s)))
	}
	return int(num)
}

func (d *Decoder) EOF() bool {
	return d.index >= len(d.input)
}

func (d *Decoder) ParseOperator() []int {
	// handle operator packet
	lengthTypeID := d.Read(1) == 1
	fmt.Printf("%02d I %#v +1\n", d.index-1, lengthTypeID)

	values := make([]int, 0)
	if lengthTypeID {
		subPacketCount := d.Read(11)
		fmt.Printf("%02d L %04b (%d) +11\n", d.index-11, subPacketCount, subPacketCount)
		for i := 0; i < subPacketCount; i++ {
			values = append(values, d.Parse())
		}
		return values
	}

	length := d.Read(15)
	fmt.Printf("%02d L %04b (%d) +15\n", d.index-15, length, length)
	start := d.index
	for d.index < start+length {
		values = append(values, d.Parse())
	}
	return values
}

func (d *Decoder) ParseHeader() (int, int) {
	packetVersion := d.Read(3)
	fmt.Printf("%02d V %03b (%d) +3\n", d.index-3, packetVersion, packetVersion)

	packetType := d.Read(3)
	fmt.Printf("%02d T %03b (%d) +3\n", d.index-3, packetType, packetType)

	d.versionSum += packetVersion
	return packetVersion, packetType
}

func (d *Decoder) ParseLiteral() int {
	// read until last group has been reached indicated by the
	// most significant bit
	r := 'A'
	result := 0
	for {
		last := d.Read(1) == 0
		fmt.Printf("%02d %c %#v +1\n", d.index-1, r, last)

		value := d.Read(4)
		result = (result << 4) | value
		fmt.Printf("%02d %c %05b (%d) +4\n", d.index-4, r, value, value)
		if last {
			break
		}
		r++
	}

	fmt.Printf("result: %b (%d)\n", result, result)
	return result
}

func (d *Decoder) StartParse() int {
	return d.Parse()
}

func (d *Decoder) Parse() int {
	_, packetType := d.ParseHeader()
	if packetType == LITERAL_PACKET {
		return d.ParseLiteral()
	}

	values := d.ParseOperator()
	result := 0
	switch packetType {
	case SUM_PACKET:
		for _, v := range values {
			result += v
		}
		fmt.Printf("sum(%#v) = %d\n", values, result)
	case PRODUCT_PACKET:
		result = 1
		for _, v := range values {
			result *= v
		}
		fmt.Printf("prod(%#v) = %d\n", values, result)
	case MINIMUM_PACKET:
		result = INFINITY
		for _, v := range values {
			if v < result {
				result = v
			}
		}
		fmt.Printf("min(%#v) = %d\n", values, result)
	case MAXIMUM_PACKET:
		for _, v := range values {
			if v > result {
				result = v
			}
		}
		fmt.Printf("max(%#v) = %d\n", values, result)
	case GREATER_PACKET:
		a, b := values[0], values[1]
		if a > b {
			result = 1
		}
		fmt.Printf("%d > %d = %d\n", a, b, result)
	case LESSER_PACKET:
		a, b := values[0], values[1]
		if a < b {
			result = 1
		}
		fmt.Printf("%d < %d = %d\n", a, b, result)
	case EQUAL_PACKET:
		a, b := values[0], values[1]
		if a == b {
			result = 1
		}
		fmt.Printf("%d == %d = %d\n", a, b, result)
	}
	return result
}

func part1(input string) int {
	d := NewDecoder(input)
	d.Parse()
	return d.versionSum
}

func part2(input string) int {
	d := NewDecoder(input)
	return d.Parse()
}

func main() {
	input := "input"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)[0]))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString(input)[0]))
	fmt.Println("too low: 58056359978")
}
