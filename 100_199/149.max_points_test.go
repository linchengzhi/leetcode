package _00_199

import "testing"

func Test_maxPoints(t *testing.T) {

	tests := []struct {
		name string
		args [][]int
		want int
	}{
		{"test1", [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}, 4},
		{"test2", [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4},
		{"test3", [][]int{{0, 0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
