package dsa_10_10_25

import (
	"fmt"
)

func bruteForce(nums []int) []int {
	arr := make([]int, len(nums))

	for i := 0; i < len(arr); i++ {
		fmt.Println("[val]", nums[i])

		sum := 1

		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			sum *= nums[j]

		}

		arr[i] = sum
	}
	fmt.Println("[result]", arr)

	return arr
}

func solution(nums []int) []int {
	arr := make([]int, len(nums))

	prefix := 1

	for i := 0; i < len(nums); i++ {
		arr[i] = prefix
		prefix = prefix * nums[i]

	}

	suffix := 1

	for j := len(nums) - 1; j >= 0; j-- {
		arr[j] = arr[j] * suffix
		suffix = nums[j] * suffix

	}

	fmt.Println("[result]", arr)

	return arr
}

func Main() {
	solution([]int{1, 2, 3, 4})
}
