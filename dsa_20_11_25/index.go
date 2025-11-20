package dsa_20_11_25

func courseSchedule(currentCourse int, visited map[int]bool, courseGraph map[int][]int) bool {
	if visited[currentCourse] {
		return false
	}

	if val, ok := courseGraph[currentCourse]; ok && len(val) == 0 {
		return true
	}

	visited[currentCourse] = true

	for _, pre := range courseGraph[currentCourse] {
		if !courseSchedule(pre, visited, courseGraph) {
			return false
		}
	}

	visited[currentCourse] = false
	courseGraph[currentCourse] = []int{}

	return true
}

func solution(numCourses int, prerequisites [][]int) bool {

	courseGraph := make(map[int][]int)

	// building the graph
	for _, val := range prerequisites {
		courseGraph[val[1]] = append(courseGraph[val[1]], val[0])
	}

	// create the map
	visited := make(map[int]bool)

	
	for val := range numCourses {
		if !courseSchedule(val, visited, courseGraph) {
			return false
		}
	}

	return true
}

func Main() {

}
