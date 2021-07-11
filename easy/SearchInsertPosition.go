// 2021.07.11

func searchInsert(nums []int, target int) int {
	if 0 == len(nums) {
		return 0
	}
	if target < nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	var idx int
	length := len(nums)
	low := 0
	high := length - 1
	for {
		mid := (low + high) / 2
		if low >= high {
			min := low
			if low > high {
				min = high
			}
			if nums[min] >= target {
				idx = min
			} else {
				idx = min + 1
			}
			break
		}
		if nums[mid] == target {
			idx = mid
			break
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return idx
}
