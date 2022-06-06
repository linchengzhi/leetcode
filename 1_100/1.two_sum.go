package __100

func twoSum(nums []int, target int) []int {
	mList := make(map[int]int)
	for key, val := range nums {
		d := target - val
		if _, ok := mList[d]; ok {
			return []int{mList[d], key}
		}
		mList[val] = key
	}
	return []int{}
}
