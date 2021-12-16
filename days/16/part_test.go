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

	d.ParseLiteral()
	require.Equal(t, len(d.input), d.index)
	require.True(t, d.EOF())
}

func TestParseLiteral2(t *testing.T) {
	d := NewDecoder("D2FE28")

	d.Parse()
	require.Equal(t, len(d.input), d.index)
	require.True(t, d.EOF())
}

func TestParseLiteral3(t *testing.T) {
	d := NewDecoder("D2FE28")

	d.StartParse()
	require.Equal(t, len(d.input), d.index)
	require.True(t, d.EOF())
}

func TestParseOperator(t *testing.T) {
	d := NewDecoder("38006F45291200")

	d.StartParse()
	require.Equal(t, len(d.input), d.index)
	require.True(t, d.EOF())
}
