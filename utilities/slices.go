package utilities

import (
	"strconv"
)

func CreateSlice[T comparable](size int, initial T) []T {
	s := make([]T, 0)
	for i := 0; i < size; i++ {
		s = append(s, initial)
	}
	return s
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
