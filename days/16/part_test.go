package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	d := NewDecoder("D2FE28")
	require.Equal(t, 0, d.index)
	require.Equal(t, "110100101111111000101000", d.input)
}

func TestParseLiteral1(t *testing.T) {
	d := NewDecoder("D2FE28")

	d.ParseHeader()
	require.Equal(t, 6, d.versionSum)
}

func TestEval1(t *testing.T) {
	require.Equal(t, 3, NewDecoder("C200B40A82").StartParse())
}

func TestEval2(t *testing.T) {
	require.Equal(t, 54, NewDecoder("04005AC33890").StartParse())
}

func TestEval3(t *testing.T) {
	require.Equal(t, 7, NewDecoder("880086C3E88112").StartParse())
}

func TestEval4(t *testing.T) {
	require.Equal(t, 9, NewDecoder("CE00C43D881120").StartParse())
}

func TestEval5(t *testing.T) {
	require.Equal(t, 1, NewDecoder("D8005AC2A8F0").StartParse())
}
