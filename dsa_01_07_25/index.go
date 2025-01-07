package dsa_01_06_25

import "fmt"

/*
  PROBLEM

  You are given an array prices where prices[i] is the price of a given stock on the ith day.

  You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

   Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
*/

/*
  Pseudocode

  1. if no price return zero

  2. initialize a min and profit

  3. loop through the prices, if the  current price is less tha the min, reassign the price to minimum

  4. if the price - min is greater than profit reassign
*/

func solution(prices []int) int {

	if len(prices) == 0 {
		return 0
	}
	profit := 0
	min := prices[0]

	for _, price := range prices {
		if price < min {
			min = price
		} else if price-min > profit {
			profit = price - min
		}
	}

	fmt.Println(profit, min)

	return profit
}

func Main() {
	solution([]int{7, 1, 5, 3, 6, 4})
}
