package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

// TODO: go over preferred form and other not-exist-checks and use them
// TODO: look at write good and alex, especially passive voice detection

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
	flags := flag.NewFlagSet("tint", flag.ExitOnError)
	verbosePtr := flags.Bool("v", false, "Verbose output (written to stderr")
	sortPtr := flags.Bool("s", false, "Sort output (by file name, line number, column number")
	jsonPtr := flags.Bool("j", false, "Use JSON output instead of default")
	check(flags.Parse(os.Args[1:]))
	args := flags.Args()

	// TODO: JSON fmt from: http://steelbrain.me/linter/types/linter-message-v2.html
	if *jsonPtr {
		panic("JSON output not implemented")
	}

	// If it should always be printed, we use log. If it should only be printed when
	// verbose=true, then we use verb
	var verb *log.Logger
	if *verbosePtr {
		verb = log.New(os.Stderr, "", 0)
	} else {
		verb = log.New(ioutil.Discard, "", 0)
	}

	verb.Printf("Verbose: %v\n", *verbosePtr)
	verb.Printf("Sort: %v\n", *sortPtr)

	wg := sync.WaitGroup{}
	report := make(chan *Warning, 64)

	for _, filename := range args {
		verb.Printf("FILE: %s\n", filename)
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

	// TODO: max warnings per file
	// TODO: max warnings overall

	// Final output channel might be filtered by sort or count options
	output := report

	// Sort filter
	if *sortPtr {
		sortInput := output
		output = make(chan *Warning, 64)

		go func(out chan *Warning) {
			defer close(out)

			warnings := make([]*Warning, 0, 128)
			for warn := range sortInput {
				warnings = append(warnings, warn)
			}

			verb.Printf("Sorting %d warnings\n", len(warnings))
			sort.Sort(WarningDefSort(warnings))

			for _, warn := range warnings {
				out <- warn
			}
		}(output)
	}

	// Output all warnings
	for warn := range output {
		fmt.Printf("%s:%d:%d:warning: %s\n",
			warn.Filename, warn.Row+1, warn.Col+1, msgEscape(warn.Msg),
		)
	}
}
