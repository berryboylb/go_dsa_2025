package dsa_09_10_25

import (
	"fmt"
	"math"
	"sort"
)

func solution(nums []int, k int) bool {

	seen := map[int]int{}

	for index, value := range nums {

		if prevIndex, ok := seen[value]; ok {
			if int(math.Abs(float64(index-prevIndex))) <= k {
				return true
			}

		}else {
seen[value] = index
		}

		
	}

	return false
}

func bruteForce(nums []int, k int) bool {
	for i, iValue := range nums {

		for j, jValue := range nums {
			if i == j {
				continue
			}

			if iValue == jValue {
				isEqual := int(math.Abs(float64(i - j)))
				if isEqual <= k {
					return true
				}

			}

		}
	}

	return false
}

func bruteForce2(nums []int, k int) bool {
	for i := range nums {
		fmt.Println("[i]", nums[i])

		for j := i + 1; j <= i+k && j < len(nums); j++ {

			fmt.Println("[jjjj]", nums[j])

			if nums[i] == nums[j] {
				isEqual := int(math.Abs(float64(i - j)))
				if isEqual <= k {
					return true
				}
			}

		}
	}

	return false
}

func sorting(nums []int, k int) bool {
	sort.Ints(nums)
	i := 1

	for i < len(nums) {

		fmt.Println("[nu]", nums[i-1], nums[i])
		// if nums[i-1] == nums[i] &&int (math.Abs(float64(i - (i -1))))<= k{
		// 	return true
		// }

		i++

	}

	return false
}

func Main() {
	sorting([]int{1,2,3,1,2,3}, 2)
}
