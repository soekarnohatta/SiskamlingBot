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
	}
	return ret
}

func IntToStr(integer int) string {
	if integer == 0{
		return ""
	}

	return strconv.Itoa(integer)
}
