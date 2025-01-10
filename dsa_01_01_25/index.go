package dsa_01_01_25

import (
	"fmt"
)

/*
  Problem

  Merge Sorted Array:

  You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

  Merge nums1 and nums2 into a single array sorted in non-decreasing order.

  The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/

/*************  ✨ Codeium Command ⭐  *************/
/*
  Merge nums1 and nums2 into a single array sorted in non-decreasing order.


/******  f52b953c-608c-4d08-a2c5-2a49af96412b  *******/func merge(nums1 []int, m int, nums2 []int, n int) {

	for index, val := range nums2 {
		zerothIndex := index + m // get the current avaliable zeroth position
		nums1[zerothIndex] = val // assign to zeroth index

		// swap if previous integer in nums1 is greater than newly inserted integer from nums2
		for index > 0 && nums1[zerothIndex] < nums1[zerothIndex-1] {
			nums1[zerothIndex], nums1[zerothIndex-1] = nums1[zerothIndex-1], nums1[zerothIndex]
			zerothIndex--
		}
	}

	fmt.Println(nums1)
}

func Main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
