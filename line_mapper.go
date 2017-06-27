package main

import "unicode"

// TODO: unit test for all of the below

type lineEntry struct {
	start int
	end   int
}

// FileMap is a simple ytpe to hold the text of a file
// It includes a mapping from offset to line/col
type FileMap struct {
	Text  string
	Lines []lineEntry
}

// NewFileMap creates a File Map, complete with pre-loaded line entries to
// map from offset to (row,col) coordinates.
// TODO: creation from file path instead of buffer
func NewFileMap(buf string) *FileMap {
	fm := FileMap{
		Text:  buf,
		Lines: make([]lineEntry, 0, (len(buf)/32)+1),
	}

	// create lines
	begin := 0
	for idx, run := range fm.Text {
		if unicode.IsSpace(run) && string(run) == "\n" {
			end := idx
			fm.Lines = append(fm.Lines, lineEntry{begin, end})
			begin = end + 1
		}
	}

	// Anything left?
	end := len(fm.Text)
	if begin <= end {
		fm.Lines = append(fm.Lines, lineEntry{begin, end})
	}

	return &fm
}

// FindOffset returns ZERO-BASED (row, col) numbers for the given offset
func (fm *FileMap) FindOffset(offset int) (int, int) {
	if offset < 1 {
		return 0, 0
	}

	if offset >= len(fm.Text)-1 {
		r := len(fm.Lines) - 1
		c := fm.Lines[r].end - r
		return r, c
	}

	// TODO: use sort.Search instead of our crap linear scan
	for i, le := range fm.Lines {
		if offset >= le.start && offset <= le.end {
			return i, offset - le.start
		}
	}

	panic("Impossible logic condition in FileMap->FindOffset")
}
