package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglePassive(t *testing.T) {
	assert := assert.New(t)

	fm, err := NewFileMap("res/passive.txt")
	assert.NoError(err)

	warnCount := 0
	for _, check := range ShouldNotPassive() {
		for _, w := range check.Match(fm) {
			warnCount++
			t.Logf("WARN %4d: %v\n", warnCount, w)
		}
	}

	assert.Equal(1, warnCount)
}
