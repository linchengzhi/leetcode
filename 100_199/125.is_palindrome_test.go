package _00_199

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"test1", "A man, a plan, a canal: Panama", true},
		{"test1", "A man, a plana", false},
		{"test1", "A", true},
		{"test1", "race a car", false},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
