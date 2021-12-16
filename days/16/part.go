package main

import (
	"fmt"
	"strconv"

	"github.com/RaphaelPour/aoc2021/util"
)

const (
	LITERAL_PACKET  = 4
	OPERATOR_PACKET = 6
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
		return 0
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
	fmt.Println("read", num, "(", bits, ") now at", d.index)
	return int(num)
}

func (d *Decoder) EOF() bool {
	return d.index >= len(d.input)
}

func (d *Decoder) ParseOperator() {
	fmt.Println("parse operator")
	// handle operator packet
	fmt.Println("read 1 byte: length type id")
	lengthTypeID := d.Read(1) == 1

	if lengthTypeID {
		fmt.Println("read 11 bytes: sub packet count")
		subPacketCount := d.Read(11)
		for i := 0; i < subPacketCount; i++ {
			d.Parse()
		}
		return
	}

	fmt.Println("read 15 bytes: sub packet length")
	length := d.Read(15)
	start := d.index
	fmt.Println("sub package: expect length", length, "until", start+length, "is reached")
	for start+length <= d.index {
		d.Parse()
	}
}

func (d *Decoder) ParseHeader() (int, int) {
	fmt.Println("read version")
	packetVersion := d.Read(3)
	fmt.Println("read type")
	packetType := d.Read(3)
	d.versionSum += packetVersion
	fmt.Println("header", packetVersion, packetType)
	return packetVersion, packetType
}

func (d *Decoder) ParseLiteral() {
	fmt.Println("parse literal")

	// read until last group has been reached indicated by the
	// most significant bit
	readBits := 0
	for {
		fmt.Println("read 5 bytes: literal value")
		value := d.Read(5)
		if value&0x10 == 0 {
			fmt.Println("msb set: leaving literal value")
			break
		}
		readBits += 5
	}

	// align input to 4 bits by reading the remaining padding
	fmt.Println("read", 4-(d.index%4), " bytes: padding of literal value")
	d.Read(4 - (d.index % 4))
}

func (d *Decoder) StartParse() {
	for !d.EOF() {
		d.Parse()
	}
}

func (d *Decoder) Parse() {
	_, packetType := d.ParseHeader()
	if packetType == LITERAL_PACKET {
		d.ParseLiteral()
	} else {
		d.ParseOperator()
	}
}

func part1(input string) int {
	d := NewDecoder(input)
	d.Parse()
	return d.versionSum
}

func part2(input string) int {
	return 0
}

func main() {
	input := "input_example"
	fmt.Println("== [ PART 1 ] ==")
	fmt.Println(part1(util.LoadString(input)[0]))

	fmt.Println("== [ PART 2 ] ==")
	fmt.Println(part2(util.LoadString(input)[0]))
}
