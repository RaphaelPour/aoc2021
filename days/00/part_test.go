package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	require.Equal(t, count([]string{""}), 0)
}
