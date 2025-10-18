package dsa_18_10_25

import (
	"fmt"
)

// passed  using a lot of space O(n) but still fast
func solution(numbers []int, target int) []int {

	seen := map[int]int{}

	for i := 0; i < len(numbers); i++ {
		rem := target - numbers[i]
		if value, ok := seen[rem]; ok {
			fmt.Println("[value]", []int{value + 1, i + 1})
			return []int{value + 1, i + 1}
		}
		seen[numbers[i]] = i
	}

	return []int{}
}

// passed using brute force O(n2)
func bruteForce(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}

			if (numbers[i] + numbers[j]) == target {
				fmt.Println("[value]", []int{i + 1, j + 1})
				return []int{i + 1, j + 1}
			}
		}
	}

	return []int{}
}

// too slow
func binarySearch(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		rem := target - numbers[i]
		subSlice := numbers[i+1:]
		low, high := 0, len(subSlice)-1
		for low <= high {
			mid := low + (high-low)/2
			if subSlice[mid] == rem {
				// real index in numbers is i+1+mid
				return []int{i + 1, i + 2 + mid}
			} else if subSlice[mid] < rem {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
		fmt.Println("[value]", rem, subSlice)
	}
	return []int{}
}

// two pointer
func pointer(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left <= right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right -= 1
		} else if sum < target {
			left += 1
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return []int{}
}

func Main() {
	pointer([]int{2, 7, 11, 15}, 9)
}
