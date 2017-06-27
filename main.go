package main

import (
	"fmt"
	"os"
	"sync"
)

// TODO: need actual checks...

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

func processFile(filename string, report chan *Warning) int {
	fm, err := NewFileMap(filename)
	check(err)
	count := 0
	for _, sidx := range IndexAll(fm.Text, "X") {
		report <- NewWarning(fm, sidx, "Found an X at offset %d!", sidx)
		count++
	}
	return count
}

func main() {
	report := make(chan *Warning, 64)
	wg := sync.WaitGroup{}

	for _, filename := range os.Args[1:] {
		wg.Add(1)
		go func(fn string) {
			defer wg.Done()
			processFile(fn, report)
		}(filename)
	}

	go func() {
		wg.Wait()
		close(report)
	}()

	// TODO: better output
	// TODO: select output format on command line
	// TODO: max warnings per file
	// TODO: max warnings overall
	// TODO: sort warnings before output
	for warn := range report {
		fmt.Printf("%s:%d:%d:warning: %s\n",
			warn.Filename, warn.Row+1, warn.Col+1, warn.Msg,
		)
	}
}
