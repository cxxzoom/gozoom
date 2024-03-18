package word1

import "unicode"

func IsPalindrome(x string) bool {
	var s []rune

	for _, v := range x {
		if unicode.IsLetter(v) {
			s = append(s, unicode.ToLower(v))
		}
	}

	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
