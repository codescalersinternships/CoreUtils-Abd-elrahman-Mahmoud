package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) != 0 {
		for {
			for i := range len(args) {
				fmt.Printf("%s ", args[i])
			}
			fmt.Printf("\n")
		}
	} else {
		for {
			fmt.Printf("y\n")
		}
	}
}
