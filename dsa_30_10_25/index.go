package dsa_30_10_25

// https://leetcode.com/problems/coin-change/description/

func solution(coins []int, amount int) int {
	arr := make([]int, amount+1)

	for i := 1; i < len(arr); i++ {
		arr[i] = amount + 1
	}

	for j := 1; j <= amount; j++ {

		for k := 0; k < len(coins); k++ {
			if j >= coins[k] {
				arr[j] = min(arr[j], 1+arr[j-coins[k]])
			}
		}
	}

	if arr[amount] < amount+1 {
		return arr[amount]
	}

	return -1
}

func Main() {
	solution([]int{1, 2, 5}, 11)
}
