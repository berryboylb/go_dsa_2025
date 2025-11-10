package dsa_10_11_25

func solution(nums []int) int {

	if nums[0] == 0 {
		return 0
	}

	jump, currMax, currEnd := 0, 0, 0
	

	for i := 0; i < len(nums)-1; i++ {
		currMax = max(currMax, i+nums[i])
		if i == currEnd {
			jump++
			currEnd = currMax
		}
	}

	return jump

}

func Main() {

}
