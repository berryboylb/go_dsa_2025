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
	for j := 1; j < numRows; j++ {
		row := triangle[j]
		row[0] = 1
		row[len(row)-1] = 1
	}

	// doing the above we solved for 1 and 2 as base case so we start from 3 upwards
	for k := 2; k < numRows; k++ {
		// start from the second element since we prefilled and stop at the second to last element
		for l := 1; l < len(triangle[k])-1; l++ {
			// go 1 row up the previous value and the value on top
			triangle[k][l] = triangle[k-1][l-1] + triangle[k-1][l]
		}
	}

	fmt.Println("[triangle]", triangle)

	return triangle

}

func Main() {
	solution(5)
}
