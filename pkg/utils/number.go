package utils

import "strconv"

func StringToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}
