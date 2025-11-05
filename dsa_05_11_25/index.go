package dsa_05_11_25

import (
	"strconv"
)

func solution(s string) int {
	if s[0] == '0' {
		return 0
	}
	one, two := 1, 1

	for i := 1; i < len(s); i++ {
		var curr int

		if s[i] != '0' {
			curr = one
		}

		val, _ := strconv.ParseInt(s[i-1:i+1], 10, 64)

		if val >= 10 && val <= 26{
			curr += two
		}

		two = one
		one = curr

	}

	return  one;
}

func Main() {

}
