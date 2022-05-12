// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var s string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += fmt.Sprintf("%d = %s\n", i, os.Args[i])
	}
	fmt.Printf("%s\n%s", os.Args[0], s)
	fmt.Printf("%.4fs elapsed\n", time.Since(start).Seconds())
}

//!-
