package main

import (
	"fmt"
	"strings"
)

func main() {
	var length, delta int
	var input string

	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &input)
	fmt.Scanf("%d\n", &delta)

	alphabetLower := "abcdefghijklmnopqrstuvwxyz"
	alphabetUpper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	ret := ""
	for _, ch := range input {
		switch {
		case strings.IndexRune(alphabetLower, ch) >= 0:
			ret += string(rotate(ch, delta, []rune(alphabetLower)))
		case strings.IndexRune(alphabetUpper, ch) >= 0:
			ret += string(rotate(ch, delta, []rune(alphabetUpper)))
		default:
			ret += string(ch)
		}
	}
	fmt.Println(ret)
}

func rotate(s rune, delta int, key []rune) rune {
	idx := strings.IndexRune(string(key), s)
	// for i, r := range key {
	// 	if r == s {
	// 		idx = i
	// 		break
	// 	}
	// }

	if idx < 0 {
		panic("idx <0")
	}

	idx = (idx + delta) % len(key)
	// for i := 0; i < delta; i++ {
	// 	idx++
	// 	if idx >= len(key) {
	// 		idx = 0
	// 	}
	// }
	return key[idx]
}
