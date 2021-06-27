// 2021.06.27

func romanToInt(s string) int {
	d := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var numeral int
	var idx int
	for idx = 0; idx < len(s)-1; idx++ {
		current := string(s[idx])
		next := string(s[idx+1])
		v1 := d[current]
		v2 := d[next]
		if v1 >= v2 {
			numeral += v1
		} else {
			numeral += v2 - v1
			idx += 1
		}
	}
	if idx == len(s)-1 {
		numeral += d[string(s[len(s)-1])]
	}
	return numeral
}
