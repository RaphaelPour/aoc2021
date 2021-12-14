package main

import (
	"fmt"
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
	t.Skip()
	rs := NewRuleSet(util.LoadString("input_example"))
	require.Equal(t, "NNCB", rs.start)
	rs.SubstituteFast(true)
	require.Equal(t, "NCNBCHB", rs.start)
	require.Equal(t, 1, rs.Result())
	rs.SubstituteFast(true)
	require.Equal(t, "NBCCNBBBCBHCB", rs.start)
	require.Equal(t, 5, rs.Result())
	rs.SubstituteFast(true)
	require.Equal(t, "NBBBCNCCNBBNBNBBCHBHHBCHB", rs.start)
	require.Equal(t, 7, rs.Result())
	rs.SubstituteFast(true)
	require.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", rs.start)
	require.Equal(t, 18, rs.Result())
	rs.SubstituteFast(true)
	require.Equal(t, 33, rs.Result())
	rs.SubstituteFast(true)
	require.Equal(t, 82, rs.Result())
	rs.SubstituteFast(true)
	require.Equal(t, 160, rs.Result())
	rs.SubstituteFast(true)
	require.Equal(t, 366, rs.Result())
	rs.SubstituteFast(true)
	require.Equal(t, 727, rs.Result())
	rs.SubstituteFast(true)
	require.Equal(t, 1588, rs.Result())
	rs.SubstituteFast(true)
}

func TestCache1(t *testing.T) {
	r := new(RuleSet)
	r.rules = make(map[string]string)

	r.EnrichCache("ABCD", "AaBbCcD")
	fmt.Println(r.rules)
	require.Equal(t, 1, len(r.rules))
	require.Equal(t, "aBbCc", r.rules["ABCD"])
}

func TestCache2(t *testing.T) {
	r := new(RuleSet)
	r.rules = make(map[string]string)

	r.EnrichCache("ABCDE", "AaBbCcDdE")
	fmt.Println(r.rules)
	require.Equal(t, 2, len(r.rules))
	require.Equal(t, "aBbCc", r.rules["ABCD"])
	require.Equal(t, "bCcDd", r.rules["BCDE"])
}

func TestCache3(t *testing.T) {
	r := new(RuleSet)
	r.rules = make(map[string]string)

	r.EnrichCache("ABCDEF", "AaBbCcDdEeF")
	fmt.Println(r.rules)
	require.Equal(t, 4, len(r.rules))
	require.Equal(t, "aBbCc", r.rules["ABCD"])
	require.Equal(t, "bCcDd", r.rules["BCDE"])
	require.Equal(t, "cDdEe", r.rules["CDEF"])
	require.Equal(t, "aBbCcDdEe", r.rules["ABCDEF"])

	r.start = "ABCDEF"
	r.SubstituteFast(false)
	require.Equal(t, "AaBbCcDdEeF", r.start)
}
