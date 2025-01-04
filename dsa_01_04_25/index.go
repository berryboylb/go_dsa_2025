package dsa_01_04_25

import "fmt"

/*
   80. Remove Duplicates from Sorted Array II

   Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same.

   Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums. More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result. It does not matter what you leave beyond the first k elements.

   Return k after placing the final result in the first k slots of nums.

   Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.
*/

/*
   Psuedocode
   1. initialize a counter and a map

   2. loop through the array of numbers

   3. if the item does not exist in the map or count of num is less than 2,
    increment the count in the map, the counter and move to the next array item

   4. if the item exceeds above conditions remove from the array
*/

func solution(nums []int) int {
	count := 0
	intMap := map[int]int{}

	for index := 0; index < len(nums); {
		val, ok := intMap[nums[index]]
		if !ok || val < 2 {
			intMap[nums[index]]++
			count++
			index++
		} else {
			nums = append(nums[:index], nums[index+1:]...)
		}
	}

	fmt.Println(count, nums)
	return count
}

func Main() {
	solution([]int{1, 1, 1, 2, 2, 3})
}
