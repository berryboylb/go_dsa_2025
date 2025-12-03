package dsa_03_12_25

import (
	// "fmt"
	"fmt"
	"sort"
)

//https://leetcode.com/problems/longest-consecutive-sequence/

// 30sec
func femiSolution(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	visited := make(map[int]bool)
	lcs, cs := 1, 1

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if visited[nums[i]] {
			continue
		} else if nums[i]-nums[i-1] == 1 {
			visited[nums[i]] = true
			cs += 1
			lcs = max(cs, lcs)
		} else {
			cs = 1
		}
	}

	return lcs
}

func recur(val int, hashMap map[int]bool, lcs *int, cs *int) {
	if !hashMap[val] {
		*cs = 0
		return
	}

	*cs += 1
	*lcs = max(*cs, *lcs)
	recur(val+1, hashMap, lcs, cs)
}

// limit exceeded
func femiSolution2(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	hashMap := make(map[int]bool)

	for _, val := range nums {
		hashMap[val] = true
	}

	lcs, cs := 0, 0

	for idx, _ := range hashMap {
		recur(idx, hashMap, &lcs, &cs)
	}

	fmt.Println("lsc", lcs)

	return lcs
}

//45 sec leetcode solution
func solution(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	hashMap := make(map[int]bool)

	for _, val := range nums {
		hashMap[val] = true
	}

	lcs := 1

	for idx, _ := range hashMap {
		if hashMap[idx-1] {
			continue
		} else {
			currentNum := idx
			currentSub := 1
			for hashMap[currentNum +1]{
				currentNum++
				currentSub++
			}
			lcs = max(currentSub, lcs)
		}
	}

	fmt.Println("lsc", lcs)

	return lcs
}

func Main() {
	solution([]int{1, 0, 1, 2})
}
