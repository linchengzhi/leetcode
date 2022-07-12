package _00_199

import "testing"

func Test_majorityElement(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		{"test1", []int{3, 2, 3}, 3},
		{"test2", []int{2, 2, 1, 1, 1, 2, 2}, 2},
		{"test3", []int{2, 1, 2, 1, 1, 2, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
