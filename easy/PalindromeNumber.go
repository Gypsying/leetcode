// 2021.06.27

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if 0 == x {
		return true
	}
	if 0 == x%10 {
		return false
	}

	reverse := ""
	s := fmt.Sprintf("%d", x)
	for idx := len(s) - 1; idx >= 0; idx-- {
		reverse += string(s[idx])
	}
	return s == reverse
}
