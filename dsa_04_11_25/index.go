package dsa_04_11_25



// https://leetcode.com/problems/house-robber-ii/description/

func solution(nums []int) int {
	rob1, rob2, max1, max2 := 0, 0, nums[0], 0

	for i := 0; i < len(nums)-1; i++ {
		max1 = max(rob1+nums[i], rob2)
		rob1 = rob2
		rob2 = max1
	}

	//reset
	rob1 = 0
	rob2 = 0

	for i := 1; i < len(nums); i++ {
		max2 = max(rob1+nums[i], rob2)
		rob1 = rob2
		rob2 = max2
	}

	return max(max1, max2)
}

func Main() {
	solution([]int{1, 2, 3, 1})
}
