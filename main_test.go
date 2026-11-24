package main

import "testing"

func TestNumDecodings(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"12", 2},  // AB, L
		{"226", 3}, // BZ, VF, BBF
		{"06", 0},  // leading zero never decodes
		{"0", 0},
		{"10", 1},  // J only
		{"27", 1},  // 27 > 26: only 2,7
		{"100", 0}, // the second zero strands
		{"101", 1}, // J A
		{"11106", 2},
		{"2101", 1}, // 2 10 1? and 21 01 invalid -> B J A
		{"", 0},
	}
	for _, c := range cases {
		if got := NumDecodings(c.in); got != c.want {
			t.Errorf("NumDecodings(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
