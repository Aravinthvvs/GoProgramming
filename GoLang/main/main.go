package main

import (
	"fmt"
)

func getLongestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	p := str[0]

	for _, s := range str {
		i := 0
		for ; i < len(s) && i < len(p) && s[i] == p[i]; i++ {
		}
		p = p[:i]
	}
	return p
}

func main() {
	str := []string{"aravinth", "arav", "arav"}
	fmt.Println(getLongestCommonPrefix(str))
}
