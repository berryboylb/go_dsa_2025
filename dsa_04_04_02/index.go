package dsa_04_04_02

import (
	"fmt"
	// "math"
)

// func solution(nums []int) []int {
// 	res := make([]int, len(nums))
// 	product := nums[0]
// 	hasZero := false

// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] != 0 {
// 			product = int(math.Abs(float64(product * nums[i])))
// 		} else {
// 			hasZero = true
// 		}

// 	}

// 	for i := 0; i < len(nums); i++ {
// 		fmt.Println(nums[i])
// 		if nums[i] == 0 {
// 			res[i] = product
// 		} else {
// 			if hasZero {
// 				res[i] = 0
// 			} else {
// 				curr := product / nums[i]
// 				fmt.Println(curr)
// 				res[i] = curr
// 			}
// 		}

// 	}

// 	fmt.Println(product, res)

// 	return res
// }

// too slow
// func solution(nums []int) []int {
// 	res := make([]int, len(nums))

// 	for i := 0; i < len(nums); i++ {
// 		left, right := i-1, i+1
// 		leftProduct, rightProduct := 1, 1

// 		for left >= 0 {
// 			leftProduct = leftProduct * nums[left]
// 			left--
// 		}

// 		for right < len(nums) {
// 			// fmt.Println("keys",nums[i], right)
// 			rightProduct = rightProduct * nums[right]
// 			right++
// 		}
// 		product := 0
// 		if leftProduct != 0 && rightProduct != 0 {
// 			product = leftProduct * rightProduct
// 		}

// 		fmt.Println(i, product, leftProduct, rightProduct)
// 		sum := -1 * 1
// 		fmt.Println(">>sum",sum)
// 		res[i] = product

// 	}

// 	fmt.Println(res)

// 	return res
// }

func solution(nums []int) []int {
	res := make([]int, len(nums))

	prefix := 1
	for i := 0; i < len(nums); i++{
		res[i] = prefix 
		prefix = prefix * nums[i]
	}

	suffix := 1
	for i := len(nums) -1; i >=0; i-- {
		res[i] = res[i] * suffix
		suffix = suffix * nums[i]
	}

	fmt.Println(res)

	return res
}

func Main() {
	solution([]int{1, 2, 3, 4})
}
