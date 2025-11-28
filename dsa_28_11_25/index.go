package dsa_28_11_25

import "fmt"

// graph
// https://leetcode.com/problems/pacific-atlantic-water-flow/

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func dfs(r int, c int, reachable [][]bool, heights [][]int, label string) {
	// Mark cell as reachable
	reachable[r][c] = true
	fmt.Printf("[%s] Visiting (%d, %d) height=%d\n", label, r, c, heights[r][c])

	for _, d := range directions {
		nr := r + d[0]
		nc := c + d[1]

		// Bounds check
		if nr < 0 || nr >= len(heights) || nc < 0 || nc >= len(heights[0]) {
			fmt.Printf("[%s]   Skip (%d,%d) -> out of bounds\n", label, nr, nc)
			continue
		}

		// Already visited
		if reachable[nr][nc] {
			fmt.Printf("[%s]   Skip (%d,%d) -> already reachable\n", label, nr, nc)
			continue
		}

		// Height condition
		if heights[nr][nc] < heights[r][c] {
			fmt.Printf("[%s]   Skip (%d,%d) -> height %d < %d\n",
				label, nr, nc, heights[nr][nc], heights[r][c])
			continue
		}

		fmt.Printf("[%s]   Flowing from (%d,%d) -> (%d,%d)\n",
			label, r, c, nr, nc)

		dfs(nr, nc, reachable, heights, label)
	}
}

func solution(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	row, col := len(heights), len(heights[0])
	pacific := make([][]bool, row)
	atlantic := make([][]bool, row)

	for i := 0; i < row; i++ {
		pacific[i] = make([]bool, col)
		atlantic[i] = make([]bool, col)
	}

	// Left + Right edges
	for i := 0; i < row; i++ {
		fmt.Printf("\n=== DFS Pacific from Left (%d,0) ===\n", i)
		dfs(i, 0, pacific, heights, "PACIFIC")

		fmt.Printf("\n=== DFS Atlantic from Right (%d,%d) ===\n", i, col-1)
		dfs(i, col-1, atlantic, heights, "ATLANTIC")
	}

	// Top + Bottom edges
	for i := 0; i < col; i++ {
		fmt.Printf("\n=== DFS Pacific from Top (0,%d) ===\n", i)
		dfs(0, i, pacific, heights, "PACIFIC")

		fmt.Printf("\n=== DFS Atlantic from Bottom (%d,%d) ===\n", row-1, i)
		dfs(row-1, i, atlantic, heights, "ATLANTIC")
	}

	result := [][]int{}

	fmt.Println("\n=== INTERSECTION RESULTS ===")
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if pacific[r][c] && atlantic[r][c] {
				fmt.Printf("Cell (%d,%d) reachable by BOTH\n", r, c)
				result = append(result, []int{r, c})
			}
		}
	}

	fmt.Println("\n=== FINAL RESULT ===")
	fmt.Println(result)

	return result
}




func Main() {
	// Example usage
	heights := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	solution(heights)
}
