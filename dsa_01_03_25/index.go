package dsa_01_02_25

import (
	"fmt"
)

/*

PROBLEM
Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k.

*/

/*
   Psuedocode
   1. set a counter and map for unique values
   2. loop through nums and check if it's not in the map
   3. add to the map, increase the array and move to the next item
*/

func removeDuplicates(nums []int) int {
	count := 0
	items := make(map[int]bool)

	for index := 0; index < len(nums); {
		if items[nums[index]] {
			fmt.Println(nums[index+1:])
			nums = append(nums[:index], nums[index+1:]...)
		} else {
			items[nums[index]] = true
			count++
			index++
		}
	}

	fmt.Println(count, nums)
	return count
}

/*
	nums = append(nums[:index], nums[index+1:]...)

	nums[:index] creates a subslice of the array which stops at last item before the index

	nums[index+1:] creates a subslice of items after the index

	... is used to flatten the new array

	the aim of this operation is to remove the element at that index
*/

func Main() {
	removeDuplicates([]int{1, 1, 2})
}
