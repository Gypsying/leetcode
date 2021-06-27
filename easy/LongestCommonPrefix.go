// 2021.06.27

func longestCommonPrefix(strs []string) string {
	if 1 == len(strs) {
		return strs[0]
	}

	minStr := strs[0]
	length := len(strs[0])
	for _, s := range strs {
		if length > len(s) {
			length = len(s)
			minStr = s
		}
	}

	var prefix string
	if 0 == length {
		return prefix
	}
	for idx, _ := range minStr {
		var flag int
		for _, s := range strs {
			if minStr[idx] == s[idx] {
				continue
			} else {
				flag = 1
				break
			}
		}
		if 1 == flag {
			break
		}
		prefix += string(minStr[idx])
	}
	return prefix
}
