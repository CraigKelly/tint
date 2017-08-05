package main

import "testing"

var warnings int
var returned int

func benchmarkFile(fn string, b *testing.B) {
	report := make(chan *Warning, 64)
	warnCount := 0
	go func() {
		for w := range report {
			if w.Offset >= 0 {
				warnCount++
			} else {
				panic("Negative offset discovered!")
			}
		}
	}()

	retCount := 0
	for n := 0; n < b.N; n++ {
		retCount += processFile(fn, report, 0)
	}

	// All this counting and module-level var storage is to insure that the
	// optimizer doesn't eliminate anything
	warnings = warnCount
	returned = retCount

	close(report)
}

func BenchmarkSmallFile(b *testing.B) {
	benchmarkFile("res/line_map_test.txt", b)
}

func BenchmarkRandomFile(b *testing.B) {
	benchmarkFile("res/random.txt", b)
}
