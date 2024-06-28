package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var numberLinesFlag bool
	var err error
	var file *os.File
	lineNumber := 1

	flag.BoolVar(&numberLinesFlag, "n", false, "Number of Lines")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Wrong Number of Arguments")
	}

	for i := range args {

		file, err = os.Open(args[i])
		if err != nil {
			fmt.Println("Error trying to open file!")
			panic(err)
		}

		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			if numberLinesFlag {
				fmt.Printf("Line Number : %d   ", lineNumber)
			}
			fmt.Printf("%s\n\n", fileScanner.Text())
			lineNumber++
		}

		defer file.Close()
	}

	defer file.Close()

}
