package util

import (
	"fmt"
	"strconv"
)

func Reverse(in string) string {
	result := ""
	for _, char := range in {
		result = string(char) + result
	}
	return result
}

func ToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(fmt.Sprintf("error converting %s to int: %s", in, err))
	}
	return out
}
