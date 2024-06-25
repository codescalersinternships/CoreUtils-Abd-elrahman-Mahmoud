package main

import (
	"flag"
    "fmt"
	"os"
	"bufio"
)

func main() {
	var numberLinesFlag bool
	var err error
	var file *os.File
	lineNumber :=1

    flag.BoolVar(&numberLinesFlag, "n", false, "Number of Lines")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0{
		log.fatal("Wrong Number of Arguments")
	}

	for i := range len(args) {

		file, err = os.Open(args[i])
		if err != nil {
			fmt.Println("Error trying to open file!")
			panic(err)
		}

		fileScanner := bufio.NewScanner(file)

		fileScanner.Split(bufio.ScanLines)
	
		for fileScanner.Scan() {
			if numberLinesFlag {
				fmt.Printf("Line Number : %d   ",lineNumber)
			}
			fmt.Printf("%s\n\n",fileScanner.Text())
			lineNumber++
		}

		file.Close()
	}

	file.Close()

}