package dsa_11_10_25

func solution(nums []int) int {
	curr, max := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if curr < 0 {
			curr = 0
		}
		curr += nums[i]
		if curr > max {
			max = curr
		}
	}

	return max
}

func bruteForce(nums []int) int {
    maxSum := nums[0]

    for i := 0; i < len(nums); i++ {
        currentSum := 0
        for j := i; j < len(nums); j++ {
            currentSum += nums[j]
            if currentSum > maxSum {
                maxSum = currentSum
            }
        }
    }

    return maxSum
}

func Main() {
	solution([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}
