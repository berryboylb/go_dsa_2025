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

		// fill first position
		triangle[i][0] = 1

		// fill the last position in the array
		if len(triangle[i]) > 1 {
			triangle[i][len(triangle[i])-1] = 1
		}
	}


	// doing the above we solved for 1 and 2 as base case so we start from 3 upwards
	for j := 2; j < numRows; j++ {
		// start from the second element since we prefilled and stop at the second to last element
		for k := 1; k < len(triangle[j])-1; k++ {
			// go 1 row up the previous value and the value on top
			triangle[j][k] = triangle[j-1][k-1] + triangle[j-1][k]
		}
	}

	fmt.Println("[triangle]", triangle)

	return triangle

}

func Main() {
	solution(10)
}
