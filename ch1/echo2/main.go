// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s := ""
	for i, arg := range os.Args[1:] {
		s += fmt.Sprintf("%d = %s\n", i, arg)
	}
	fmt.Printf("%s\n%s", os.Args[0], s)
	fmt.Printf("%.4fs elapsed\n", time.Since(start).Seconds())
}

//!-
