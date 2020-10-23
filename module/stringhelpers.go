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
