package dsa_21_11_25

import (
	"fmt"
)

// pass the pointer of isPossible so we can overide, the rest are gonna pass by reference except node
func dfs(node int, isPossible *bool, color map[int]int, adjList map[int][]int, GRAY int, WHITE int, BLACK int, topologicalOrder *[]int) {

	if !*isPossible {
		return
	}

	// begin racism
	color[node] = GRAY

	// segregate the people, whites only
	for _, neighbour := range adjList[node] {
		// if that mf white pay him
		if color[neighbour] == WHITE {
			dfs(neighbour, isPossible, color, adjList, GRAY, WHITE, BLACK, topologicalOrder)
		} else if color[neighbour] == GRAY {
			*isPossible = false // if not slightly deny wages
		}
	}

	// send to the gulag
	color[node] = BLACK
	*topologicalOrder = append(*topologicalOrder, node)

}

func solution(numCourses int, prerequisites [][]int) []int {
	WHITE := 1
	GRAY := 2
	BLACK := 3

	isPossible := true
	color := make(map[int]int)
	adjList := make(map[int][]int)
	topologicalOrder := []int{}

	// build the graph and start racism
	for num := 0; num < numCourses; num++ {
		color[num] = WHITE
	}

	// capture everyone
	for _, value := range prerequisites {
		dest, src := value[0], value[1]
		adjList[src] = append(adjList[src], dest)
	}

	// pay the whites
	for num := 0; num < numCourses; num++ {
		if color[num] == WHITE {
			dfs(num, &isPossible, color, adjList, GRAY, WHITE, BLACK, &topologicalOrder)
		}
	}

	order := []int{}

	if isPossible {
		order = make([]int, numCourses)
		for i := 0; i < numCourses; i++ {
			order[i] = topologicalOrder[numCourses-i-1]
		}
	}

	return order

}

func femiSolution(numCourses int, prerequisites [][]int) []int {
	prereq := make(map[int]int)
	res := []int{}

	for _, v := range prerequisites {
		prereq[v[0]] = v[1]
	}

	for course := 0; course < numCourses; course++ {
		// no prerequisite → start of chain
		if _, ok := prereq[course]; !ok {
			res = append(res, course)
		}
	}

	for course := 0; course < numCourses; course++ {
		// has prerequisite → later in chain
		if _, ok := prereq[course]; ok {
			res = append(res, course)
		}
	}

	fmt.Println("[res]", res)

	return res

}

func Main() {
	solution(2, [][]int{{1, 0}})
}
