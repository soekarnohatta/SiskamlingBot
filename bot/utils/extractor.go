package utils

import (
	"strconv"
	"strings"
	"time"
)

func ExtractBool(text string) bool {
	return strings.ToLower(text) == "true"
}

func ExtractTime(timeVal string) int64 {
	lastLetter := timeVal[len(timeVal)-1:]
	lastLetter = strings.ToLower(lastLetter)
	var ret int64 = 0

	if strings.ContainsAny(lastLetter, "m & d & h") {
		t := timeVal[:len(timeVal)-1]
		timeNum, err := strconv.Atoi(t)

		if err != nil {
			return -1
		}

		if lastLetter == "m" {
			ret = time.Now().Unix() + int64(timeNum*60)
		} else if lastLetter == "h" {
			ret = time.Now().Unix() + int64(timeNum*60*60)
		} else if lastLetter == "d" {
			ret = time.Now().Unix() + int64(timeNum*24*60*60)
		} else {
			return -1
		}

		return ret
	} else {
		return -1
	}
}
