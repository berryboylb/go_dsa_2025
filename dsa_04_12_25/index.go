package dsa_04_12_25

// https://leetcode.com/problems/number-of-provinces/description/



func bfs(isConnected [][]int, visited []bool, start int) {
	queue := []int{start}
	visited[start] = true

	for len(queue) > 0 {
		city := queue[0]
		queue = queue[1:]

		for next := 0; next < len(isConnected); next++ {
			if isConnected[city][next] == 1 && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
}


func solution(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)
	provinces := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			bfs(isConnected, visited, i)
			provinces++
		}
	}

	return provinces
}

func Main() {
}
