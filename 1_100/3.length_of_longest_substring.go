package __100

func lengthOfLongestSubstring(s string) int {
	mList := make(map[rune]int)
	start, length := -1, 0
	for key, val := range s {
		if _, ok := mList[val]; ok && start < mList[val]{
			start = mList[val]
		} else {
			if key - start > length {
				length = key - start
			}
		}
		mList[val] = key
	}
	return length
}