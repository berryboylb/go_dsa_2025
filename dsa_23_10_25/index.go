package dsa_23_10_25

// next permutation
// https://leetcode.com/problems/next-permutation/description/

func swap(nums []int, i int, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverse(nums []int, i int) {
	j := len(nums) - 1

	for i < j {
		swap(nums, i, j)
		//use this in the next iteration of the loop
		i++
		j--
	}

}

func solution(nums []int) {
	i := len(nums) - 2

	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1

		for nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}

	reverse(nums, i+1)
}
