package main

import "fmt"

// NumDecodings counts the ways to decode a digit string where A=1 … Z=26.
// Each position extends decodings by a one-digit letter (1-9) and/or a
// two-digit letter (10-26); zeros are only legal inside 10 or 20.
func NumDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	prev, cur := 1, 1 // ways for prefixes of length i-1 and i
	for i := 1; i < len(s); i++ {
		next := 0
		if s[i] != '0' {
			next = cur // single-digit decode of s[i]
		}
		two := (s[i-1]-'0')*10 + (s[i] - '0')
		if two >= 10 && two <= 26 {
			next += prev // two-digit decode of s[i-1:i+1]
		}
		if next == 0 {
			return 0 // stranded zero or impossible pair: no decoding at all
		}
		prev, cur = cur, next
	}
	return cur
}

func main() {
	fmt.Println(NumDecodings("226"))
	fmt.Println(NumDecodings("06"))
}
