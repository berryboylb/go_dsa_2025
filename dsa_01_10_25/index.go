package dsa_01_10_25



/*
28. Find the Index of the First Occurrence in a String
Solved
Easy
Topics
Companies
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.


*/

func solution(nums []int) int {
	if len(nums) <= 1 {
		return 0 // No jumps needed if the array has 1 or 0 elements
	}

	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < len(nums)-1; i++ {
		// Update the farthest point reachable
		farthest = max(farthest, i+nums[i])

		// If we reach the end of the current range
		if i == currentEnd {
			jumps++
			currentEnd = farthest

			// If the farthest point is beyond or at the last index, stop
			if currentEnd >= len(nums)-1 {
				break
			}
		}
	}
	return jumps
}

// Helper function to find the maximum of two numbers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func Main() {
	solution([]int{2, 3, 1, 1, 4})
}
