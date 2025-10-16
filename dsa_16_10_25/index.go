package dsa161025



func solution(nums []int, target int) int {

	high, low := len(nums)-1, 0

	for low < high {
		mid := (low + high) / 2

		if nums[mid] == target {
			return mid
		}

		// if this part of the array is sorted
		if nums[low] <= nums[mid] {
			// if the target is not in the sorted part go to the other part of the array
			if target < nums[low] || target > nums[mid] {
				low = mid + 1
			} else {
				// move the higher pointer  to the sorted part of the array
				high = mid - 1
			}
		} else {
			// if the target is not in the sorted part go to the other part of the array
			if target > nums[mid] || target < nums[high] {
				high = mid - 1
			} else {
				// move the lower pointer  to the sorted part of the array
				low = mid + 1
			}
		}
	}
	return -1
}

//chatgpt
func search(nums []int, target int) int {
    low, high := 0, len(nums)-1

    for low <= high {
        mid := (low + high) / 2

        if nums[mid] == target {
            return mid
        }

        // Left side is sorted
        if nums[low] <= nums[mid] {
            if target >= nums[low] && target < nums[mid] {
                high = mid - 1
            } else {
                low = mid + 1
            }
        } else {
            // Right side is sorted
            if target > nums[mid] && target <= nums[high] {
                low = mid + 1
            } else {
                high = mid - 1
            }
        }
    }

    return -1
}


// brute force works

func bruteForce(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		}

	}
	return 0
}

func Main() {
	solution([]int{4, 5, 6, 7, 0, 1, 2}, 0)
}
