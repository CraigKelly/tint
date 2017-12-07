package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

// Warning emitted by a checker for a file
type Warning struct {
	Filename string
	Offset   int
	Row      int
	Col      int
	Msg      string
}

// TextCheck is our generic check interface is currently only implemented by BadTerm
type TextCheck interface {
	Match(*FileMap) []*Warning
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

// WarningDefSort provides our default sort order for output: ascending
// by (File Name, File Position)
type WarningDefSort []*Warning

func (a WarningDefSort) Len() int {
	return len(a)
}

func (a WarningDefSort) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a WarningDefSort) Less(i int, j int) bool {
	if a[i].Filename < a[j].Filename {
		return true
	}
	if a[i].Filename > a[j].Filename {
		return false
	}

	return a[i].Offset < a[j].Offset
}

// WarningJSON conforms to http://steelbrain.me/linter/types/linter-message-v2.html
type WarningJSON struct {
	Location WarningLocation `json:"location"`
	Excerpt  string          `json:"excerpt"`
	Severity string          `json:"severity"`
}

// WarningLocation conforms to http://steelbrain.me/linter/types/linter-message-v2.html
type WarningLocation struct {
	File     string  `json:"file"`
	Position [][]int `json:"position"`
}

// CreateJSON returns a warning in JSON format matching http://steelbrain.me/linter/types/linter-message-v2.html
func (w Warning) CreateJSON() string {
	p1 := []int{w.Row, w.Col}
	p2 := []int{w.Row, w.Col}

	fn, err := filepath.Abs(w.Filename)
	check(err)

	data := WarningJSON{
		Location: WarningLocation{
			File:     fn,
			Position: [][]int{p1, p2},
		},
		Excerpt:  w.Msg,
		Severity: "warning",
	}

	body, err := json.Marshal(data)
	check(err)
	return string(body)
}
