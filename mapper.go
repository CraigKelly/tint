package main

import (
	"io/ioutil"
	"sort"
	"strings"
	"unicode"

	"github.com/mvdan/xurls"
)

type lineEntry struct {
	start int
	end   int
}

// FileMap is a simple type to hold the text of a file
// It includes a mapping from offset to line/col
// Text is the actual text read from the file, while CheckText is the
// processed text used for searching. Mainly it is set to lower case
// by convention so that our checks are faster. Doing this with regex's
// instead of the (?i) flag is an order of magnitude faster.
type FileMap struct {
	Filename  string
	Text      string
	CheckText string
	Lines     []lineEntry
}

// createCheckText returns a string that is suitable for fast searching. The
// returned string will be the same length as the source string.
func createCheckText(text string) string {
	// Always lower case the string so we don't need case-insensitive regex's
	newText := strings.ToLower(text)

	// Space out anything that we shouldn't search...

	// URL's
	newText = xurls.Strict.ReplaceAllStringFunc(newText, func(toRep string) string {
		return strings.Repeat(" ", len(toRep))
	})

	return newText
}

// NewFileMap creates a File Map, complete with pre-loaded line entries to
// map from offset to (row,col) coordinates.
func NewFileMap(filename string) (*FileMap, error) {
	bufb, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	buf := string(bufb)
	fm := FileMap{
		Filename:  filename,
		Text:      buf,
		CheckText: createCheckText(buf),
		Lines:     make([]lineEntry, 0, (len(buf)/32)+1),
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

	return &fm, nil
}

// FindOffset returns ZERO-BASED (row, col) numbers for the given offset
func (fm *FileMap) FindOffset(offset int) (int, int) {
	if offset < 1 {
		return 0, 0
	}

	if offset >= len(fm.Text)-1 {
		r := len(fm.Lines) - 1
		c := fm.Lines[r].end - fm.Lines[r].start
		return r, c
	}

	// Find smallest index such that offset occurs before the end of the line
	idx := sort.Search(len(fm.Lines), func(i int) bool {
		return offset <= fm.Lines[i].end
	})
	// Insure line finding is working - note we've already checked boundary case
	// offsets above, so if we fail here, it's a full-on bug
	if idx < len(fm.Lines) {
		le := fm.Lines[idx]
		if offset >= le.start && offset <= le.end {
			return idx, offset - le.start
		}
	}

	panic("Impossible logic condition in FileMap->FindOffset")
}
