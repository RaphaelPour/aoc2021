package main

import (
	"testing"

	"github.com/RaphaelPour/aoc2021/util"
	"github.com/stretchr/testify/require"
)

func TestSubstituteOnce(t *testing.T) {
	rs := NewRuleSet(util.LoadString("input_example"))
	require.Equal(t, "NNCB", rs.start)
	rs.SubstituteOnce()
	require.Equal(t, "NCNBCHB", rs.start)
	require.Equal(t, 1, rs.Result())
	rs.SubstituteOnce()
	require.Equal(t, "NBCCNBBBCBHCB", rs.start)
	require.Equal(t, 5, rs.Result())
	rs.SubstituteOnce()
	require.Equal(t, "NBBBCNCCNBBNBNBBCHBHHBCHB", rs.start)
	require.Equal(t, 7, rs.Result())
	rs.SubstituteOnce()
	require.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", rs.start)
	require.Equal(t, 18, rs.Result())
	rs.SubstituteOnce()
	require.Equal(t, 33, rs.Result())
	rs.SubstituteOnce()
	require.Equal(t, 82, rs.Result())
	rs.SubstituteOnce()
	require.Equal(t, 160, rs.Result())
	rs.SubstituteOnce()
	require.Equal(t, 366, rs.Result())
	rs.SubstituteOnce()
	require.Equal(t, 727, rs.Result())
	rs.SubstituteOnce()
	require.Equal(t, 1588, rs.Result())
	rs.SubstituteOnce()
}

func TestSubstituteFast(t *testing.T) {
	rs := NewRuleSet(util.LoadString("input_example"))
	require.Equal(t, "NNCB", rs.start)
	rs.SubstituteFast()
	require.Equal(t, "NCNBCHB", rs.start)
	require.Equal(t, 1, rs.Result())
	rs.SubstituteFast()
	require.Equal(t, "NBCCNBBBCBHCB", rs.start)
	require.Equal(t, 5, rs.Result())
	rs.SubstituteFast()
	require.Equal(t, "NBBBCNCCNBBNBNBBCHBHHBCHB", rs.start)
	require.Equal(t, 7, rs.Result())
	rs.SubstituteFast()
	require.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", rs.start)
	require.Equal(t, 18, rs.Result())
	rs.SubstituteFast()
	require.Equal(t, 33, rs.Result())
	rs.SubstituteFast()
	require.Equal(t, 82, rs.Result())
	rs.SubstituteFast()
	require.Equal(t, 160, rs.Result())
	rs.SubstituteFast()
	require.Equal(t, 366, rs.Result())
	rs.SubstituteFast()
	require.Equal(t, 727, rs.Result())
	rs.SubstituteFast()
	require.Equal(t, 1588, rs.Result())
	rs.SubstituteFast()
}
