package dsa_12_11_25


func racecar(target int) int {
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = -1
	}
	return helper(target, dp)
}

func helper(target int, dp []int) int {
	if dp[target] >= 0 {
		return dp[target]
	}

	n := 1
	for (1<<n)-1 < target {
		n++
	}

	// Case 1: exact match — reach target in n steps
	if (1<<n)-1 == target {
		dp[target] = n
		return n
	}

	// Case 2: overshoot and reverse
	dp[target] = n + 1 + helper((1<<n)-1-target, dp)

	// Case 3: undershoot, reverse earlier
	for m := 0; m < n-1; m++ {
		distance := target - ((1 << (n-1)) - 1) + ((1 << m) - 1)
		dp[target] = min(dp[target], (n-1)+1+m+1+helper(distance, dp))
	}

	return dp[target]
}

func Main() {

}