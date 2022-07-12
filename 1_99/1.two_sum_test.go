package __99

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := []int{0, 1}
	d := twoSum(nums, target)
	if !IntSliceEqualBCE(result, d) {
		t.Log("fail")
	}

	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := []int{1, 2}
	d2 := twoSum(nums2, target2)
	if !IntSliceEqualBCE(result2, d2) {
		t.Log("fail")
	}

	nums3 := []int{3, 3}
	target3 := 6
	result3 := []int{0, 1}
	d3 := twoSum(nums3, target3)
	if !IntSliceEqualBCE(result3, d3) {
		t.Log("fail")
	}
}

func IntSliceEqualBCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
