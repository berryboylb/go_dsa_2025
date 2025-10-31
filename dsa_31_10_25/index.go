package dsa_31_10_25

func solution(nums []int) int {
	arr := make([]int, len(nums))
	maxVal :=1

	for i := 0; i < len(nums); i++ {
		arr[i] = 1
	}


	for j := 1; j < len(nums); j++ {
		for k := 0; k <j; k++{
			if nums[j] > nums[k]{
				arr[j] = max(arr[j], 1+arr[k])
				maxVal = max(maxVal, arr[j])
			}
		}
	}

	return  maxVal;
}

func Main() {
	solution([]int{1, 2, 5})
}
