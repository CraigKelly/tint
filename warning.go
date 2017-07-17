package main

import "fmt"

// Warning emitted by a checker for a file
type Warning struct {
	Filename string
	Offset   int
	Row      int
	Col      int
	Msg      string
}

// NewWarning creates a warning, complete with formatting
func NewWarning(fm *FileMap, offset int, msgf string, args ...interface{}) *Warning {
	r, c := fm.FindOffset(offset)
	msg := fmt.Sprintf(msgf, args...)
	w := Warning{
		Filename: fm.Filename,
		Offset:   offset,
		Row:      r,
		Col:      c,
		Msg:      msg,
	}
	return &w
}
