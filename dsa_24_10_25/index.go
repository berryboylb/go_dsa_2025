package dsa_24_10_25

import (
	"fmt"
)

func solution(nums []int) int {
	seen := map[int]int{}

	for i := 0; i < len(nums)-1; {
		if _, ok := seen[nums[i]]; ok {
			// this mutates the array to be everything  before and after i
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			seen[nums[i]] = nums[i]
			i++
		}
	}

	fmt.Println("[k]", len(seen))
	fmt.Println("[nums]", nums)

	return len(seen)
}

func Main() {
	solution([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
}
