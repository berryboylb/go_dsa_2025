package dsa_15_10_25


func solution(nums []int) int {

	high, low := len(nums)-1, 0

	for  low < high {
		mid := (low + high) / 2

		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

func Main() {
	solution([]int{3, 4, 5, 1, 2})
}
