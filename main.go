package main

import "fmt"
import "io/ioutil"
import "os"
import "strings"

// TODO: need to start replacing with actual checks...
func indexAll(s, sep string) []int {
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

func main() {
	// TODO: all this is just testing until we have real functionality
	textb, err := ioutil.ReadAll(os.Stdin)
	check(err)
	fm := NewFileMap(string(textb))
	fmt.Printf("Read %d bytes\n", len(fm.Text))

	// TODO: unit testing
	for _, arg := range os.Args[1:] {
		fmt.Printf("Looking for %v\n", arg)
		for i, sidx := range indexAll(fm.Text, arg) {
			r, c := fm.FindOffset(sidx)
			// Note that we print 1-based
			fmt.Printf(" [%5d]=> offset:%8d (row %d, col %d)\n", i, sidx, r+1, c+1)
		}
		fmt.Printf("----\n")
	}
}
