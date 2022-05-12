// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	fileCount := make(map[string]map[string]int) // make map of filename to counts
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileCount)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileCount)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			for fn, fc := range fileCount[line] {
				fmt.Printf("%s\t%d\t%s\n", fn, fc, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileCount map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		fileName := f.Name()
		counts[text]++
		if fileCount[text] == nil {
			fileCount[text] = make(map[string]int)
		}
		fileCount[text][fileName]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
