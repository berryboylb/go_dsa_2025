package dsa_08_11_25

import (
	"strings"
)

// https://leetcode.com/problems/generate-parentheses/

func backtrack(ans *[]string, cur *strings.Builder, open, close, maxNum int) {
    if cur.Len() == maxNum*2 {
        *ans = append(*ans, cur.String())
        return
    }

    if open < maxNum {
        cur.WriteByte('(')
        backtrack(ans, cur, open+1, close, maxNum)
        // Backtrack: remove last character
        curStr := cur.String()
        cur.Reset()
        cur.WriteString(curStr[:len(curStr)-1])
    }

    if close < open {
        cur.WriteByte(')')
        backtrack(ans, cur, open, close+1, maxNum)
        curStr := cur.String()
        cur.Reset()
        cur.WriteString(curStr[:len(curStr)-1])
    }
}


func solution(n int) []string {
	var ans []string
	var builder strings.Builder
	backtrack(&ans, &builder, 0, 0, n)
	return ans
}

func Main() {

}
