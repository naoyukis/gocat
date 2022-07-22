package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func open(arg string) error {
	f, err := os.Open(arg)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := io.Copy(os.Stdout, f); err != nil {
		return err
	}
	return nil
}

func main() {
	flag.Parse()

	if flag.NArg() > 0 {
		for _, arg := range flag.Args() {
			if err := open(arg); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	} else {
		if _, err := io.Copy(os.Stdout, os.Stdin); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
