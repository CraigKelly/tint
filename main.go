package main

import (
	"fmt"
	"os"
	"sync"
)

// TODO: more checks

func processFile(filename string, report chan *Warning) int {
	fm, err := NewFileMap(filename)
	check(err)

	count := 0

	for _, term := range ShouldNotExist() {
		for _, warn := range term.Match(fm) {
			report <- warn
			count++
		}
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
	// TODO: select output format on command line - should include JSON fmt from: http://steelbrain.me/linter/types/linter-message-v2.html
	// TODO: max warnings per file
	// TODO: max warnings overall
	// TODO: option to sort warnings before output
	for warn := range report {
		fmt.Printf("%s:%d:%d:warning: %s\n",
			warn.Filename, warn.Row+1, warn.Col+1, warn.Msg,
		)
	}
}
