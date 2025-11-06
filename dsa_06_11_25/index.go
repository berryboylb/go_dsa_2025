package dsa_06_11_25


func solution(m int, n int) int {

	// Create a 2D slice (matrix) with (len1+1) rows
	matrix := make([][]int, m)

	// Initialize each row with (len2+1) columns, all set to 0
	// This creates a grid of size (len1+1) x (len2+1)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Fill top row with 1s
	for j := 0; j < n; j++ {
		matrix[0][j] = 1
	}

	// Fill left column with 1s
	for i := 1; i < m; i++ {
		matrix[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// to go up a row reduce index of  outer loop
			top := matrix[i-1][j]
			// to go back a column reduce index of inner loop
			left := matrix[i][j-1]
			matrix[i][j] = left + top
		}
	}


	// return the last element in the 2d array by reducing both values m and n by one since array are zero indexed
	return matrix[m-1][n-1]
}

func Main() {
	solution(3, 3)
}
