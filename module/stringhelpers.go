package module

import (
	"strings"
)

func Join(s []string) string {
	return strings.Join(s, "")
}

func Reverse(s string) (r string) {
	for _, v := range s {
		r = string(v) + r
	}
	return
}

func Interleave(s []string) (r string) {
	// need?
	str := make([]string, len(s))
	copy(str, s)
	var maxLen int

	for _, v := range str {
		if maxLen < len([]rune(v)) {
			maxLen = len([]rune(v))
		}
	}

	var result string
	for i := 0; i < maxLen; i++ {
		for j := 0; j < len(str); j++ {
			for _, v := range str[j] {
				result = result + string(v)
				break
			}
			if len([]rune(str[j])) != 0 {
				str[j] = string([]rune(str[j])[1:])
			}
		}
	}
	return result
}
