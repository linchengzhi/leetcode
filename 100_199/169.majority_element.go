package _00_199

func majorityElement(nums []int) int {
	var count = 1
	var num = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == num {
			count++
			continue
		}
		count--
		if count == 0 {
			num = nums[i+1]
		}
	}
	return num
}
