package dsa_08_10_25

import (
	"sort"
)

func solution(nums []int) bool {

	seen := map[int]bool{}

	for _, val := range nums {

		if _, ok := seen[val]; ok {
			return true
		}

		seen[val] = true
	}

	return false
}

// brute force works but not optimal
func bruteForce(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i != j {
				return true
			}
		}

	}
	return false
}

// sorting
func sorting(nums []int) bool {
	sort.Ints(nums)
	i := 1

	for i < len(nums) {
		if nums[i-1] == nums[i] {
			return true
		}

		i++
	}

	return false
}

func Main() {
	sorting([]int{1,2,3,1})
}
