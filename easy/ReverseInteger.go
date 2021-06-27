// 2021.06.27

import (
	"fmt"
	"strconv"
)

func power(x, y int) int {
	if 0 == x {
		return 0
	}
	if 0 == y {
		return 1
	}
	return x * power(x, y-1)
}

func reverse(x int) int {
	factor := 1
	if x < 0 {
		x = x * -1
		factor = -1
	}

	var lst []int
	s := fmt.Sprintf("%d", x)
	for idx := 0; idx < len(s); idx++ {
		// ch := fmt.Sprintf("%c", s[idx])
		num, _ := strconv.Atoi(string(s[idx]))
		lst = append([]int{num}, lst...)
	}

	sum := 0
	for idx := 0; idx < len(lst); idx++ {
		num := lst[idx]
		if 0 == num {
			continue
		}
		sum += num * power(10, len(lst)-idx-1)
	}

	y := factor * sum
	min := power(2, 31)
	max := min - 1
	min = -1 * min
	if y < min || y > max {
		return 0
	}
	return y
}
