// 2021.06.27
func twoSum(nums []int, target int) []int {
	var indices []int
	for idx := 0; idx < len(nums); idx++ {
		n1 := nums[idx]
		for idx2 := idx + 1; idx2 < len(nums); idx2++ {
			n2 := nums[idx2]
			if n2 == target-n1 {
				indices = append(indices, idx)
				indices = append(indices, idx2)
				break
			}
		}
		if 2 == len(indices) {
			break
		}
	}
	return indices
}
