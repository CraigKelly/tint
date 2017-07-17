package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Xlate and assertion using 1-based expectations
func assertRC(assert *assert.Assertions, fm *FileMap, offset int, expR int, expC int) {
	actR, actC := fm.FindOffset(offset)
	assert.Equal(expR, actR+1, "Row from offset did not match expected")
	assert.Equal(expC, actC+1, "Col from offset did not match expected")
}

func TestCornerCaseMapping(t *testing.T) {
	assert := assert.New(t)

	fm, err := NewFileMap("res/line_map_test.txt")
	assert.NoError(err)

	assertRC(assert, fm, -1, 1, 1)
	assertRC(assert, fm, 0, 1, 1)
	assertRC(assert, fm, 1, 1, 2)

	lineCount := len(fm.Lines)
	line := fm.Lines[lineCount-1]
	finalCol := line.end - line.start
	assertRC(assert, fm, 100000, lineCount, finalCol+1) //1-based
}

// indexAll returns all indexes in s where sep occurs
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

func TestBasicFileMapping(t *testing.T) {
	assert := assert.New(t)

	fm, err := NewFileMap("res/line_map_test.txt")
	assert.NoError(err)

	xfound := indexAll(fm.Text, "X")
	assertRC(assert, fm, xfound[0], 9, 1)
	assertRC(assert, fm, xfound[1], 9, 3)
	assertRC(assert, fm, xfound[2], 9, 5)
	assertRC(assert, fm, xfound[3], 10, 2)
	assertRC(assert, fm, xfound[4], 10, 4)
	assertRC(assert, fm, xfound[5], 11, 1)
	assertRC(assert, fm, xfound[6], 11, 3)
	assertRC(assert, fm, xfound[7], 11, 5)
	assertRC(assert, fm, xfound[8], 12, 2)
	assertRC(assert, fm, xfound[9], 12, 4)

	yfound := indexAll(fm.Text, "Y")
	assertRC(assert, fm, yfound[0], 9, 2)
	assertRC(assert, fm, yfound[1], 9, 4)
	assertRC(assert, fm, yfound[2], 10, 1)
	assertRC(assert, fm, yfound[3], 10, 3)
	assertRC(assert, fm, yfound[4], 10, 5)
	assertRC(assert, fm, yfound[5], 11, 2)
	assertRC(assert, fm, yfound[6], 11, 4)
	assertRC(assert, fm, yfound[7], 12, 1)
	assertRC(assert, fm, yfound[8], 12, 3)
	assertRC(assert, fm, yfound[9], 12, 5)
}
