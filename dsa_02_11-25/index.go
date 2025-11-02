package dsa_02_11_25

// https://leetcode.com/problems/word-break

func soln(s string, wordDict []string) bool {

	arr := make([]bool, len(s)+1)
	arr[0] = true
	set := map[string]struct{}{}

	for _, w := range wordDict {
		set[w] = struct{}{}
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {

			if _, ok := set[s[j:i]]; ok && arr[j] {
				arr[i] = true
				break
			}

		}
	}

	return arr[len(s)]
}

func Main() {

}
