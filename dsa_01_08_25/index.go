package dsa_01_08_25

import "fmt"

/*

122. Best Time to Buy and Sell Stock II

You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

Find and return the maximum profit you can achieve.
*/

/*
  Pseudocode

  1. loop through the array from the second index

  2. if day i  - day i-1 profit i greater than zero add it
*/



func solution(prices []int) int {
	maxProfit := 0
    
    for i := 1; i < len(prices); i++ {
        profit := prices[i] - prices[i-1]
        if profit > 0 {
            maxProfit += profit
        }
    }

	fmt.Println(maxProfit)
    
    return maxProfit
}

func Main() {
	solution([]int{7, 1, 5, 3, 6, 4})
}
