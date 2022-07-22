package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() > 0 {
		for _, arg := range flag.Args() {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if _, err := io.Copy(os.Stdout, f); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	} else {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
