package _00_199

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	var result = make([][]string, 0)
	result = append(result, []string{string(s[0])})
	for i := 1; i < len(s); i++ {
		result = partitionCycle(result, string(s[i]))
	}
	return result
}

func partitionCycle(list [][]string, s string) [][]string {
	var result = make([][]string, 0)
	for _, val := range list {
		d := make([]string, 0)
		d = append(d, val...)
		d = append(d, s)
		result = append(result, d)
		endStr := val[len(val)-1]
		if checkPalindrome(endStr + s) {
			d2 := make([]string, 0)
			d2 = append(d2, val[:len(val)-1]...)
			d2 = append(d2, endStr+s)
			result = append(result, d2)
		} else {
			if len(val) <= 1 {
				continue
			}
			endStr = val[len(val)-2] + val[len(val)-1]
			if checkPalindrome(endStr + s) {
				d2 := make([]string, 0)
				d2 = append(d2, val[:len(val)-2]...)
				d2 = append(d2, endStr+s)
				result = append(result, d2)
			}
		}
	}
	return result
}

func checkPalindrome(list string) bool {
	if len(list) == 0 {
		return false
	}
	if len(list) == 1 {
		return true
	}
	var i, j = 0, len(list) - 1
	for i < j {
		if list[i] != list[j] {
			return false
		}
		i++
		j--
	}
	return true
}
