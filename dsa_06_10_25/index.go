package dsa_06_10_25



func solution(nums []int, target int) []int {

	items := map[int]int{}

	for i, v := range nums {
		remainder := target - v
		if val, ok := items[remainder]; ok {
			return []int{val, i}
		}

		items[v] = i
	}

	return []int{}

}

func bruteForce(nums []int, target int) []int {

	for i, v := range nums {

		remainder := target - v

		for index, innerValue := range nums {
			if innerValue == remainder && i != index {
				return []int{i, index}
			}

		}

	}
	return []int{}
}


func Main() {
	solution([]int{2, 7, 11, 15}, 9)
}
