package __100

import "testing"

func TestName(t *testing.T) {
	if findMedianSortedArrays([]int{1, 3}, []int{2}) != 2 {
		t.Error("fail")
	}

	if findMedianSortedArrays([]int{1, 2}, []int{3, 4}) != 2.5 {
		t.Error("fail")
	}

}
