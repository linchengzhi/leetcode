package _00_199

import "testing"

func Test_singleNumber(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want int
	}{
		// TODO: Add test cases.
		{"test1", []int{2, 2, 1}, 1},
		{"test1", []int{4, 1, 2, 1, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
