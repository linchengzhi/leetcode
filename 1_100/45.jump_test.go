package __100

import "testing"

func TestJump(t *testing.T) {
	//var num = []int{2,3,1,1,4}
	//if jump(num) != 2 {
	//	t.Error("test fail")
	//}
	//
	//var num2 = []int{2,3,0,1,4}
	//if jump(num2) != 2 {
	//	t.Error("test fail")
	//}

	var num3 = []int{10,3,0,1,4}
	if jump(num3) != 1 {
		t.Error("test fail")
	}
	var num4 = []int{2,1}
	if jump(num4) != 1 {
		t.Error("test fail")
	}
}
