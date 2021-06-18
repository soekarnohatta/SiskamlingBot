package util

import (
	"log"
	"strconv"
)

func StrToInt(text string) int {
	if text == "" {
		return 0
	}

	ret, err := strconv.Atoi(text)
	if err != nil {
		log.Println(err.Error())
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
