package dsa_22_10_25


func solution(words []string, order string) bool {
	orderMap := make(map[rune]int)
	for idx, val := range order {
		orderMap[val] = idx
	}

	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := min(len(w1), len(w2))
		isDifferent := false

		for j := 0; j < minLen; j++ {
			if w1[j] != w2[j] {
				currLetter := orderMap[rune(w1[j])]
				nextLetter := orderMap[rune(w2[j])]
				if nextLetter < currLetter {
					return false
				}
				isDifferent = true
				break
			}
		}

		// If all previous letters are same but w1 is longer (e.g., "batman" > "bat")
		if !isDifferent && len(w1) > len(w2) {
			return false
		}
	}

	return true
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }


func Main() {

}
