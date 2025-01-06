package dsa_01_06_25

import "fmt"

/*
   Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
*/

/*
   Pseudocode

   1. check for edge cases where k is zero or array len is 1

   2. if k is greater than array length  subtract the length from k

   3. get truncation index by subtracting k for array length

   4. divide the array into two, the part to shifted and part adjusted

   5. for a new array

   6. reassign the values to the first array

   
*/

func solution(nums []int, k int) {
	if k == 0 || len(nums) ==1 {
        return
    }
	for k > len(nums){
		k -= len(nums)
	}
	index := len(nums) -k
	newArr := append(nums[index:], nums[:index+1]...)
	for i := 0; i < len(nums); i++ {
		// nums = append(nums[len(nums)-1:], nums[:len(nums)-1]...)
		nums[i] = newArr[i]
	}
	fmt.Println(nums)
}

func Main() {
	solution([]int{-1, -100, 3, 99}, 2)
}
