package dsa_01_11_25

import "sort"

/*
274. H-Index

Hint
Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return the researcher's h-index.

According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.
*/

/*

  Pseudocode

*/

func solution(citations []int) int {

	n := len(citations)

	sort.Slice(citations, func(i, j int) bool {
		return citations[j] > citations[i]
	})

	for index, _ := range citations {
		if citations[index] >= n-index {
			return n - index
		}
	}

	return 0
}

func Main() {
	solution([]int{3, 0, 6, 1, 5})
}
