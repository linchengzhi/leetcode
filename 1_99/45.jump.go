package __99

func jump(nums []int) int {
	l := len(nums)
	maxPos := 0
	endPos := 0
	step := 0
	for i := 0; i < l-1; i++ {
		if nums[i]+i > maxPos {
			maxPos = nums[i] + i
		}
		if endPos == i {
			step++
			endPos = maxPos
		}
	}
	return step
}
