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
// TODO: passive voice detection (optional on command line) -- see here:
//       http://matt.might.net/articles/shell-scripts-for-passive-voice-weasel-words-duplicates/

func processFile(filename string, report chan *Warning) int {
	fm, err := NewFileMap(filename)
	check(err)

	wg := sync.WaitGroup{}
	counts := make(chan int, 16)

	launch := func(checks []TextCheck) {
		wg.Add(1)
		go func(fm *FileMap, checks []TextCheck) {
			defer wg.Done()
			count := 0
			for _, c := range checks {
				for _, warn := range c.Match(fm) {
					report <- warn
					count += 1
				}
			}
			counts <- count
		}(fm, checks)
	}

	launch(ShouldNotExist())
	launch(ShouldNotCliche())
	launch(ShouldNotRedundant())

	go func() {
		wg.Wait()
		close(counts)
	}()

	finalCount := 0
	for c := range counts {
		finalCount += c
	}
	return finalCount
}

func main() {
	flags := flag.NewFlagSet("tint", flag.ExitOnError)
	verbosePtr := flags.Bool("v", false, "Verbose output (written to stderr")
	sortPtr := flags.Bool("s", false, "Sort output (by file name, line number, column number")
	jsonPtr := flags.Bool("j", false, "Use JSON output instead of default")
	check(flags.Parse(os.Args[1:]))
	args := flags.Args()

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
	var outputter func(w *Warning)
	if *jsonPtr {
		outputter = func(w *Warning) {
			fmt.Printf("%s\n", w.CreateJSON())
		}
	} else {
		outputter = func(w *Warning) {
			fmt.Printf("%s:%d:%d:warning: %s\n",
				w.Filename, w.Row+1, w.Col+1, msgEscape(w.Msg),
			)
		}
	}

	for warn := range output {
		outputter(warn)
	}
}
