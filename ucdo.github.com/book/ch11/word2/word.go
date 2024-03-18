package word2

import "unicode"

func IsPalindrome(x string) bool {
	s := make([]rune, len(x))
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
