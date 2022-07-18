package _00_199

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	var i, j = 0, len(s) - 1
	for i < j {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i++
			continue
		}
		if !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j--
			continue
		}
		if strings.ToUpper(string(s[i])) != strings.ToUpper(string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}
