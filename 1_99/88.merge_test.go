package __99

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test1", args: args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}, want: []int{1, 2, 2, 3, 5, 6}},
		{name: "test2", args: args{[]int{0}, 0, []int{2}, 1}, want: []int{2}},
		{name: "test3", args: args{[]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "test4", args: args{[]int{1}, 1, []int{}, 0}, want: []int{1}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		if !reflect.DeepEqual(tt.want, tt.args.nums1) {
			t.Errorf("fial")
		}
	}
}
