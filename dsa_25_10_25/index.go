package dsa251025



func solution(nums []int, target int) []int {

	left, right := 0, len(nums)-1

	for left <= right {

		if nums[left] < target {
			left++
		} else if nums[right] > target {
			right--
		} else if nums[left] == target && nums[right] == target {
			return []int{left, right}
		} else {
			break
		}
	}

	return []int{-1, -1}
}

func Main() {
	solution([]int{5, 7, 7, 8, 8, 10}, 8)
}
