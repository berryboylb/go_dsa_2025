package dsa_01_02_25

import (
	"fmt"
	"sort"
)

/*
  Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
Return k.

*/

/* 

	1. loop through the code and count the occurence that are not val
	2. set every occurence of val to -1
	3. sort the array
	4. return the count

*/

func removeElement(nums []int, val int) int {
	count := 0
	for index, num := range nums {
		if num != val {
			count++
		} else {
			nums[index] = -1
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	 })

	fmt.Println(count, nums)
	return count
}

func Main() {
	removeElement([]int{0,1,2,2,3,0,4,2}, 2)
}
