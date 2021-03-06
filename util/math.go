package util

import "math"

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Sign(num int) int {
	if num < 0 {
		return -1
	}
	return 1
}

func Min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func Max(nums ...int) int {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func MinMax(nums ...int) (int, int) {
	min := nums[0]
	max := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}
	return min, max
}

func Pow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	result := base
	for exp-1 > 0 {
		result *= base
		exp--
	}
	return result
}

func InRange(num, min, max int) bool {
	return num >= min && num <= max
}

func Radians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}
