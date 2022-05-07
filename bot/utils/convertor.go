package utils

import (
	"math"
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

func Int64ToStr(integer int64) string {
	if integer == 0 {
		return ""
	}

	return strconv.Itoa(int(integer))
}

func StrToIntSlice(s []string) []int {
	var newIntSlice []int
	for _, val := range s {
		var newInt = StrToInt(val)
		newIntSlice = append(newIntSlice, newInt)
	}
	return newIntSlice
}

func StrToInt64Slice(s []string) []int64 {
	var newInt64Slice []int64
	for _, val := range s {
		var newInt64 = StrToInt64(val)
		newInt64Slice = append(newInt64Slice, newInt64)
	}
	return newInt64Slice
}

func ConvertSeconds(input uint64) (result string) {
	if input != 0 {
		var years = math.Floor(float64(input) / 60 / 60 / 24 / 7 / 30 / 12)
		var seconds = input % (60 * 60 * 24 * 7 * 30 * 12)
		var months = math.Floor(float64(seconds) / 60 / 60 / 24 / 7 / 30)
		seconds = input % (60 * 60 * 24 * 7 * 30)
		var weeks = math.Floor(float64(seconds) / 60 / 60 / 24 / 7)
		seconds = input % (60 * 60 * 24 * 7)
		var days = math.Floor(float64(seconds) / 60 / 60 / 24)
		seconds = input % (60 * 60 * 24)
		var hours = math.Floor(float64(seconds) / 60 / 60)
		seconds = input % (60 * 60)
		var minutes = math.Floor(float64(seconds) / 60)
		seconds = input % 60

		if years > 0 {
			result = plural(int(years), "year") + plural(int(months), "month") + plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
		} else if months > 0 {
			result = plural(int(months), "month") + plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
		} else if weeks > 0 {
			result = plural(int(weeks), "week") + plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
		} else if days > 0 {
			result = plural(int(days), "day") + plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
		} else if hours > 0 {
			result = plural(int(hours), "hour") + plural(int(minutes), "minute") + plural(int(seconds), "second")
		} else if minutes > 0 {
			result = plural(int(minutes), "minute") + plural(int(seconds), "second")
		} else {
			result = plural(int(seconds), "second")
		}

		return
	}

	return
}

func plural(count int, singular string) (result string) {
	if (count == 1) || (count == 0) {
		result = strconv.Itoa(count) + " " + singular + " "
	} else {
		result = strconv.Itoa(count) + " " + singular + "s "
	}

	return
}
