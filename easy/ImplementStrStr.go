// 2021.06.27

func recurse(haystack string, needle string) int {
	if 0 == len(needle) || needle == haystack {
		return 0
	}
	if haystack[0] == needle[0] {
		return recurse(haystack[1:], needle[1:])
	} else {
		return -1
	}
}

func strStr(haystack string, needle string) int {
	if 0 == len(needle) {
		return 0
	}
	if len(needle) > len(haystack) || (len(needle) == len(haystack) && needle != haystack) {
		return -1
	}

	index := -1
	ch := needle[0]
	for j, _ := range haystack {
		if haystack[j] == ch && len(haystack[j:]) >= len(needle) {
			tmp := recurse(haystack[j+1:], needle[1:])
			// !!critical!!
			if tmp != -1 && (0 == tmp && len(needle) != 0) {
				index = j
				break
			}

		}
	}
	return index
}
