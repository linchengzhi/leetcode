package _00_199

func singleNumber(nums []int) int {
	var num = nums[0]
	for i := 1; i < len(nums); i++ {
		num = num ^ nums[i]
	}
	return num
}
