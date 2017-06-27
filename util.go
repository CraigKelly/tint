package main

import "strings"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// IndexAll returns all indexes in s where sep occurs
func IndexAll(s, sep string) []int {
	ret := []int{}
	if len(sep) < 1 {
		return ret
	}

	p := 0
	for p < len(s) {
		idx := strings.Index(s[p:], sep)
		if idx < 0 {
			break
		}

		if len(ret) < 1 {
			ret = make([]int, 0, 4)
		}

		actualIdx := p + idx
		ret = append(ret, actualIdx)
		p = actualIdx + len(sep)
	}

	return ret
}
