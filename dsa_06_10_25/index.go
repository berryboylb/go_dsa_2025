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

func Main() {
	solution([]int{2, 7, 11, 15}, 9)
}
