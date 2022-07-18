package _00_199

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	tests := []struct {
		name string
		args string
		want [][]string
	}{
		//{"test1", "a", [][]string{{"a"}}},
		//{"test2", "ab", [][]string{{"a", "b"}}},
		//{"test3", "aba", [][]string{{"a", "b", "a"}, {"aba"}}},
		{"test3", "cabac", [][]string{{"c", "a", "b", "a", "c"}, {"c", "aba", "c"}, {"cabac"}}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
