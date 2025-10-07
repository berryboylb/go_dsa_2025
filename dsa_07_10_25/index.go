package dsa_07_10_25

import "fmt"

func solution(prices []int) int {
	min, profit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		currentProfit := prices[i] - min

		fmt.Println(">>> [log] <<<", currentProfit, prices[i], min)
		if currentProfit > profit {
			profit = currentProfit
		}

		if prices[i] < min {
			min = prices[i]
		}
	}

	fmt.Println(">>> [profit] <<<", profit)

	return profit
}

func bruteForce(prices []int) int {

	profit := 0

	for i := 0; i < len(prices); i++ {

		for j := i + 1; j < len(prices); j++ {
			curr := prices[j] - prices[i]

			if curr > profit {
				profit = curr
			}
		}

	}
	fmt.Println(">>> [profit] <<<", profit)
	return profit
}

func Main() {
	bruteForce([]int{7,1,5,3,6,4})
}
