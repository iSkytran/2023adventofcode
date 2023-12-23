package utilities

import (
	"math"
	"strconv"
)

func CreateSlice[T comparable](size int, initial T) []T {
	s := make([]T, 0)
	for i := 0; i < size; i++ {
		s = append(s, initial)
	}
	return s
}

func ElementHammingDistance[T comparable](s1 []T, s2 []T) int {
	if len(s1) != len(s2) {
		return int(math.Abs(float64(len(s1) - len(s2))))
	}

	distance := 0
	for idx := range s1 {
		if s1[idx] != s2[idx] {
			distance++
		}
	}

	return distance
}

func StringsToInts(stringSlice []string) []int {
	// Convert list of strings to a list of integers.
	intSlice := make([]int, 0)
	for _, str := range stringSlice {
		val, _ := strconv.Atoi(str)
		intSlice = append(intSlice, val)
	}
	return intSlice
}

func IntsToStrings(intSlice []int) []string {
	// Convert list of integers to a list of integers.
	strSlice := make([]string, 0)
	for _, val := range intSlice {
		str := strconv.Itoa(val)
		strSlice = append(strSlice, str)
	}
	return strSlice
}
