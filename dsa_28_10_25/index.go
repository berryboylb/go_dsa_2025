package dsa_28_10_25

import (
	"fmt"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays
func merge(a []int, b []int) []int {
	n, m := len(a), len(b)
	i, j, k, c := 0, 0, 0, make([]int, n+m)

	for i < n || j < m {
		if j == m || (i < n && a[i] <= b[j]) {
			c[k] = a[i]
			i++
			k++
		} else {
			c[k] = b[j]
			j++
			k++

		}
	}

	return c
}

func bruteForce(nums1 []int, nums2 []int) float64 {
	arr := merge(nums1, nums2)
	n := len(arr)
	var median float64

	if len(arr)%2 == 0 {
		mid1 := arr[(n/2)-1]
		mid2 := arr[n/2]
		median = float64(mid1+mid2) / 2.0
	} else {
		median = float64(arr[n/2])
	}
	fmt.Println("[median]", median)

	return median
}

func solution(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return solution(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	start, end := 0, x
	for start <= end {
		partitionX := (start + end) / 2
		partitionY := (x+y+1)/2 - partitionX

		var xLeft int
		var xRight int
		var yLeft int
		var yRight int
		if partitionX == 0 {
			//Shift the bit pattern of -1 left 31 times” → gives you the most negative int32
			xLeft = -1 << 31
		} else {
			xLeft = nums1[partitionX-1]
		}

		if partitionX == x {
			//Shift the bit pattern of 1 right 31 times” → gives you the most positive int32
			xRight = 1<<31 - 1
		} else {
			xRight = nums1[partitionX]
		}

		if partitionY == 0 {
			yLeft = -1 << 31
		} else {
			yLeft = nums2[partitionY-1]
		}

		if partitionY == y {
			yRight = 1<<31 - 1
		} else {
			yRight = nums2[partitionY]
		}

		if xLeft <= yRight && yLeft <= xRight {
			if (x+y)%2 == 0 {
				return float64(max(xLeft, yLeft)+min(xRight, yRight)) / 2.0
			} else {
				return float64(max(xLeft, yLeft))
			}
		}else if xLeft > yRight {
			end = partitionX - 1
		}else {
			start = partitionX + 1
		}
	}

	return 0.0
}

func Main() {
	solution([]int{1, 2}, []int{3, 4})
}
