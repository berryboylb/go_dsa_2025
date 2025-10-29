package dsa_29_10_25

// https://leetcode.com/problems/climbing-stairs/

import (
	"fmt"
)

func solution(n int) int {
	ways := map[int]int{
		0: 1,
		1: 1,
	}

	for i := 2; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]

	}

	fmt.Println("[ways]", ways[n])

	return ways[n]
}

func solution2(n int) int {
	prev1 := 1
	prev2 := 1

	for i := 2; i <= n; i++ {
		curr := prev1 + prev2
		prev2 = prev1
		prev1 = curr
	}

	fmt.Println("[prev1]", prev1)

	return prev1
}

var memo = map[int]int{
	0: 1,
	1: 1,
}

func recursion(n int) int {
	if n <= 1 {
		return 1
	}
	memo[n] = recursion(n-1) + recursion(n-2)
	return memo[n]
}

func Main() {
	solution2(5)
}
