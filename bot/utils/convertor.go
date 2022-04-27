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

func IntToStr(integer int) string {
	if integer == 0 {
		return ""
	}

	return strconv.Itoa(integer)
}

func StrToIntSlice(s []string) []int {
	var newIntSlice []int
	for _, val := range s {
		newInt, _ := strconv.Atoi(val)
		newIntSlice = append(newIntSlice, newInt)
	}
	return newIntSlice
}
