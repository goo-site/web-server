package utils

import (
	"github.com/goo-site/log"
	"strconv"
)

func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Error("[StringToInt64] parse err")
		return 0
	}
	return i
}
