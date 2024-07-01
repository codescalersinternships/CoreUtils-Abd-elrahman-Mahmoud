package main

import (
	"flag"
	"fmt"
)

func main() {
	var newLineFlag bool

	flag.BoolVar(&newLineFlag, "n", false, "New Line")

	flag.Parse()

	args := flag.Args()

	if len(args) != 0 {
		for i := range args {
			fmt.Printf("%s ", args[i])
			if newLineFlag {
				fmt.Printf("\n")
			}
		}
	}
}
