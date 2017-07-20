package main

import "strconv"

func msgEscape(s string) string {
	esc := strconv.Quote(s)
	ln := len(esc)
	if ln <= 2 {
		return ""
	}
	return esc[1 : ln-1]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
