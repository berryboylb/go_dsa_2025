package dsa_27_10_25

// https://leetcode.com/problems/trapping-rain-water/description/

func solution(height []int) int {
	total, left, right := 0, 0, len(height)-1
	leftMax := height[left]
	rightMax := height[right]

	for left <= right {
		if height[left] < height[right] {
			leftMax = max(leftMax, height[left])
			canStore := leftMax - height[left]
			if canStore > 0 {
				total += canStore
			}
			left++
		} else {
			rightMax = max(rightMax, height[right])
			canStore := rightMax - height[right]
			if canStore > 0 {
				total += canStore
			}
			right--
		}
	}

	return total
}

func Main() {

}
