// 2021.06.27
/*
// quick sort "Time Limit Exceeded": 11 / 13 test cases passed.
func sortArray(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }
    
    var left []int
    var right []int
    pivot := nums[0]
    nums = nums[1:]
    for _, num := range nums {
        if num < pivot {
            left = append(left, num)
        } else {
            right = append(right, num)
        }
    }
    
    left = append(sortArray(left), pivot)
    lst := append(left, sortArray(right)...)
    return lst
}
*/

func merge(left, right []int) []int {
	var lst []int
	for {
		if 0 == len(left) || 0 == len(right) {
			break
		}
		if left[0] < right[0] {
			lst = append(lst, left[0])
			left = left[1:]
		} else {
			lst = append(lst, right[0])
			right = right[1:]
		}
	}

	if 0 != len(left) {
		for _, num := range left {
			lst = append(lst, num)
		}
	}
	if 0 != len(right) {
		for _, num := range right {
			lst = append(lst, num)
		}
	}
	return lst
}

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	midx := len(nums) / 2
	return merge(sortArray(nums[:midx]), sortArray(nums[midx:]))
}
