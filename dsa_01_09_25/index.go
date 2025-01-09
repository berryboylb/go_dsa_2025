package dsa_01_09_25

import (
	"Math"
)

/*
55. Jump Game
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.
*/

/*

  Pseudocode

*/

func solution(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	maxReach := float64(0)

	for i := 0; i < len(nums); i++ {
		if i > int(maxReach) {
			return false
		}
		maxReach = math.Max(float64(maxReach), float64(i+nums[i]))
	}
	return maxReach > 0
}

func Main() {
	solution([]int{2, 3, 1, 1, 4})
}
