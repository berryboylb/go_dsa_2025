package dsa_03_11_25

import (
	"fmt"
)

// https://leetcode.com/problems/house-robber/description/



func solution(nums []int) int {
	rob1, rob2, maxRes := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		maxRes = max(rob1+nums[i], rob2)
		rob1 = rob2
		rob2 = maxRes
		fmt.Println("[res]",rob1, rob2, maxRes)
	}

	return maxRes
}

func Main() {
	solution([]int{1, 2, 3, 1})
}
