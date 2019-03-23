package golang

import (
	"strconv"
)

func addBinary(a string, b string) string {
	var result string
	var carry, sum int
	ia, ib := len(a)-1, len(b)-1
	for ia >= 0 || ib >= 0 {
		ca, cb := 0, 0
		if ia >= 0 {
			ca, _ = strconv.Atoi(string(a[ia]))
			ia--
		}
		if ib >= 0 {
			cb, _ = strconv.Atoi(string(b[ib]))
			ib--
		}
		sum = carry + ca + cb
		result = strconv.Itoa(sum%2) + result
		carry = sum / 2
	}
	if carry == 1 {
		result = "1" + result
	}
	return result
}
