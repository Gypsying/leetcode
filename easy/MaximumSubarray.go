// 2021.07.18

// 暴力求解，时间复杂度是O(N^3)
// func maxSubArray(nums []int) int {
// 	sum := nums[0]
// 	for i := 0; i < len(nums); i++ {
// 		for j := i; j < len(nums); j++ {
// 			var s int
// 			for idx := i; idx <= j; idx++ {
// 				s += nums[idx]
// 			}
// 			if s > sum {
// 				sum = s
// 			}
// 		}
// 	}
// 	return sum
// }

// 暴力求解优化，时间复杂度是O(N^2)
// 事实上，上面的代码有一些重复计算，
// 这是因为相同前缀的区间求和，即后一个区间的和=当前值+前一个区间的和。
// func maxSubArray(nums []int) int {
//     sum := nums[0]
//     for i := 0; i < len(nums); i++ {
//         var s int
//         for j := i; j < len(nums); j++ {
//             s += nums[j]
//             if s > sum {
//                 sum = s
//             }
//         }
//     }
//     return sum
// }

// 动态规划，时间复杂度是O(N)
// https://leetcode-cn.com/problems/maximum-subarray/solution/zheng-li-yi-xia-kan-de-dong-de-da-an-by-lizhiqiang/
// func maxSubArray(nums []int) int {
//     if 0 == len(nums) {
//         return 0
//     }
//     // 先计算每个子组的最大值
//     idx2max := map[int]int{}
//     idx2max[0] = nums[0]
//     for idx:=1; idx<len(nums); idx++ {
//         if idx2max[idx-1] > 0 {
//             idx2max[idx] = nums[idx] + idx2max[idx-1]
//         } else {
//             idx2max[idx] = nums[idx]
//         }
//     }
//     // 找到子组的最大值即全局最大值
//     sum := idx2max[0]
//     for _, max := range idx2max {
//         if max > sum {
//             sum = max
//         }
//     }
//     return sum
// }

// 动态规划优化，降低空间复杂度
func maxSubArray(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	sum := nums[0]    // 记录全局最大值
	subMax := nums[0] // 只需要一个int变量保存前面子组合的最大值
	for idx := 1; idx < len(nums); idx++ {
		if subMax > 0 {
			subMax += nums[idx]
		} else {
			subMax = nums[idx]
		}
		if subMax > sum {
			sum = subMax
		}
	}
	return sum
}
