package dsa_01_05_25


/*
   169. Majority Element
   Given an array nums of size n, return the majority element.

   The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.
*/


/*
   Pseudocode

   1. initialize a map, max and max occurrence

   2. go through loop and increment current index occurrence in the map

   3. if current index occurrence is higher than max occurrence
    replace max and max occurrence

   4. return max
*/

func solution(nums []int) int {
	intMap := map[int]int{}
	max := 0
	occ := 0

	for index := 0; index < len(nums); index++ {
		intMap[nums[index]]++

		if intMap[nums[index]] > occ {
			max = nums[index]
			occ = intMap[nums[index]]
		}
	}

	return max
}

func Main() {
	solution([]int{2, 2, 1, 1, 1, 2, 2})
}
