package dsa_09_11_25

func solution(nums []int) bool {

	final := len(nums) - 1

	for i := len(nums) - 2; i >= 0; i-- {

		if i+nums[i] >= final {
			final = i
		}

	}

	return final == 0
}

func Main() {

}
