package dsa_21_10_25

import (
	"fmt"
)

func solution(height []int) int {
	max, left, right := 0, 0, len(height)-1

	for left <= right {
		currHeight := 0
		moveRight := false
		if height[left] > height[right] {
			currHeight = height[right]
		} else {
			moveRight = true
			currHeight = height[left]
		}
		width := right - left
		area := currHeight * width

		if area > max {
			max = area
		}

		if moveRight {
			left++
		} else {
			right--
		}
	}

	fmt.Println("[max]", max)

	return max
}

func bruteForce(height []int) int {
	max := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			currHeight := 0
			if height[i] > height[j] {
				currHeight = height[j]
			} else {
				currHeight = height[i]
			}
			width := j - i
			area := currHeight * width

			if area > max {
				max = area
			}
		}
	}

	fmt.Println("[max]", max)

	return max
}

func Main() {
	solution([]int{
		1, 8, 6, 2, 5, 4, 8, 3, 7,
	})
}
