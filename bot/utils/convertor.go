package utils

import (
	"strconv"
)

func StrToInt(text string) int {
	ret, err := strconv.Atoi(text)
	if err != nil {
		return 0
	}

	return ret
}

func StrToInt64(text string) int64 {
	return int64(StrToInt(text))
}

func IntToStr(integer int) string {
	if integer == 0 {
		return ""
	}

	return strconv.Itoa(integer)
}

func StrToIntSlice(s []string) []int {
	var newIntSlice []int
	for _, val := range s {
		newInt := StrToInt(val)
		newIntSlice = append(newIntSlice, newInt)
	}
	return newIntSlice
}

func StrToInt64Slice(s []string) []int64 {
	var newInt64Slice []int64
	for _, val := range s {
		newInt64 := StrToInt64(val)
		newInt64Slice = append(newInt64Slice, newInt64)
	}
	return newInt64Slice
}
