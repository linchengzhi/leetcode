package __100

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	if lengthOfLongestSubstring("abcabcbb") != 3 {
		t.Error("fail")
	}
	if lengthOfLongestSubstring("bbbbb") != 1 {
		t.Error("fail")
	}
	if lengthOfLongestSubstring("pwwkew") != 3 {
		t.Error("fail")
	}
	if lengthOfLongestSubstring("tmmzuxt") != 5 {
		t.Error("fail")
	}

}
