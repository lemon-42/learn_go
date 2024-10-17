package iteration

import "strings"

func Repeat(character string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}

func GoRepeat(character string, n int) string {
	return strings.Repeat(character, n)
}
