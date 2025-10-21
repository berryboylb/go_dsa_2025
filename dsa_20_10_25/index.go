package dsa_20_10_25

import (
	"fmt"
	"sort"
	"strconv"
)

func twoSum(nums []int, i int, result [][]int) [][]int {
	left, right := i+1, len(nums)-1

	for left < right {
		sum := nums[i] + nums[left] + nums[right]
		if sum < 0 {
			left++
		} else if sum > 0 {
			right--
		} else {
			result = append(result, []int{nums[i], nums[left], nums[right]})
			left++
			right--

			// skip duplicates
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		}
	}

	return result
}

func solution(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			res = twoSum(nums, i, res)
		}

		// left, right, rem := i+1, len(nums)-1, 0-nums[i]
		// for left <= right {
		// 	if left+right < rem {
		// 		left += 1
		// 	} else if left+right > rem {
		// 		right -= 1
		// 	} else {
		// 		res = append(res, []int{nums[i], nums[left], nums[right]})
		// 	}
		// }
	}

	fmt.Println("[data]", res)

	return res
}

func bruteForce(nums []int) [][]int {
	res := make([][]int, 0)
	seen := map[string]bool{}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					arr := []int{nums[i], nums[j], nums[k]}
					sort.Ints(arr)
					key := strconv.Itoa(arr[0]) + strconv.Itoa(arr[1]) + strconv.Itoa(arr[2])
					if _, ok := seen[key]; ok {
						continue
					}
					res = append(res, arr)
					seen[key] = true

				}
			}
		}
	}

	fmt.Println("[data]", res)

	return res
}

func Main() {
	solution([]int{0, 0, 0})
}
