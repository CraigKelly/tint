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

// TODO: atom editor plugin for linter (should work with our JSON mode)

func processFile(filename string, report chan *Warning, maxWarnings int) int {
	fm, err := NewFileMap(filename)
	check(err)

	// We launch a goroutine per list of TextChecks and write any warnings to
	// our gather channel. Then we pump from the gather channel to our caller.
	// There are many things we COULD do as we gather the warnings, but
	// currently this is only used to handle the max warnings per file feature.
	wg := sync.WaitGroup{}
	gather := make(chan *Warning, 64)

	launch := func(checks []TextCheck) {
		wg.Add(1)

		go func(fm *FileMap, checks []TextCheck) {
			defer wg.Done()

			count := 0
			for _, c := range checks {
				for _, warn := range c.Match(fm) {
					gather <- warn
					count += 1
					if maxWarnings > 0 && count >= maxWarnings {
						// optimization: see below for actual maxWarnings work
						return
					}
				}
			}
		}(fm, checks)
	}

	launch(ShouldNotExist())
	launch(ShouldNotCliche())
	launch(ShouldNotRedundant())
	launch(ShouldNotProfane())
	launch(ShouldNotBias())
	launch(ShouldNotPassive())

	go func() {
		wg.Wait()
		close(gather)
	}()

	// Actually funnel warnings back to caller
	finalCount := 0

	for warn := range gather {
		report <- warn
		finalCount += 1
		if maxWarnings > 0 && finalCount >= maxWarnings {
			break
		}
	}

	return finalCount
}

func main() {
	flags := flag.NewFlagSet("tint", flag.ExitOnError)
	verbosePtr := flags.Bool("v", false, "Verbose output (written to stderr")
	sortPtr := flags.Bool("s", false, "Sort output (by file name, line number, column number")
	jsonPtr := flags.Bool("j", false, "Use JSON output instead of default")
	maxAllPtr := flags.Int("max", 0, "Maximum warnings to output")
	maxFilePtr := flags.Int("filemax", 0, "Maximum warnings per file")
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
			processFile(fn, report, *maxFilePtr)
		}(filename)
	}

	go func() {
		wg.Wait()
		close(report)
	}()

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

	written := 0
	for warn := range output {
		outputter(warn)
		written += 1
		if *maxAllPtr > 0 && written >= *maxAllPtr {
			verb.Printf("Maximum warnings reached (%d)\n", written)
			break
		}
	}
}
