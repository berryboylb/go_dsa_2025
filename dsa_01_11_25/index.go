package dsa_01_11_25

// https://leetcode.com/problems/longest-common-subsequence

func solution(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i-1][j]
				if dp[i][j-1] > dp[i][j] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
		return dp[len1][len2]
	}

	return 0
}

func sol(text1 string, text2 string) int {
	// Get the lengths of both strings
	len1, len2 := len(text1), len(text2)

	// Create a 2D slice (matrix) with (len1+1) rows
	matrix := make([][]int, len1+1)

	// Initialize each row with (len2+1) columns, all set to 0
	// This creates a grid of size (len1+1) x (len2+1)
	for i := range matrix {
		matrix[i] = make([]int, len2+1)
	}

	// Loop backward through both strings (bottom-up approach)
	// This ensures subproblems are already solved when needed
	for j := len2 - 1; j >= 0; j-- {
		for i := len1 - 1; i >= 0; i-- {

			// If characters match, take 1 plus the value diagonally down-right
			// (which represents the LCS length of the remaining substrings)
			if text1[i] == text2[j] {
				matrix[i][j] = 1 + matrix[i+1][j+1]
			} else {
				// Otherwise, take the maximum of moving down or right
				// (skip one character from either string)
				matrix[i][j] = max(matrix[i+1][j], matrix[i][j+1])
			}
		}
	}

	// The top-left cell contains the final LCS length
	return matrix[0][0]
}


func Main() {
}
