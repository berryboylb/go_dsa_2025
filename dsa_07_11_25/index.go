package dsa_07_11_25

//https://leetcode.com/problems/pascals-triangle/submissions/1823550661/

import (
	"fmt"
)

func solution(numRows int) [][]int {

	if numRows == 0 {
		return [][]int{}
	}

	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
	}

	// fill first position
	triangle[0][0] = 1

	// fill the first and last elements in other rows with 1
	for i := 1; i < numRows; i++ {
		row := triangle[i]
		row[0] = 1
		row[len(row)-1] = 1
	}

	// doing this we solved for 1 and 2 as base case so we start from 3 upwards
	for i := 2; i < numRows; i++ {
		// start from the second element since we prefilled
		for j := 1; j < len(triangle[i])-1; j++ {
			// go 1 row up the previous value and the value on top
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	fmt.Println("[triangle]", triangle)

	return triangle

}

func Main() {
	solution(5)
}
