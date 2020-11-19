package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(s string) bool {

	dict := make(map[rune]bool)

	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) && dict[r] {

			return false
		}
		dict[r] = true
	}
	return true
}
