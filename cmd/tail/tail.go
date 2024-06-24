package main

import (
	"flag"
    "fmt"
	"strconv"
	"os"
	"bufio"
)

func main() {
	var numberofLinesFlag bool
	var numberofLines int
	var err error
	var fileName string
	var file *os.File
	var linetoStartPrinting int

	totalNumberofLines := 0 
	currentLineCount := 0

    flag.BoolVar(&numberofLinesFlag, "n", false, "Number of Lines")

	flag.Parse()

	args := flag.Args()
	if numberofLinesFlag && len(args) == 2 {
		numberofLines, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error trying to convert string to integer!")
		    panic(err)
		}
		fileName = args[1]
	} else if len(args) == 1 {
		numberofLines = 10
		fileName = args[0]
	} else {
		fmt.Println("Wrong Number of Arguments")
		os.Exit(0)
	}

	file, err = os.Open(fileName)
	if err != nil {
		fmt.Println("Error trying to open file!")
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
 
	for fileScanner.Scan() {
		totalNumberofLines++
	}

	file.Close()

	file, err = os.Open(fileName)
	if err != nil {
		fmt.Println("Error trying to open file!")
		panic(err)
	}
	
	linetoStartPrinting = totalNumberofLines - numberofLines

	fileScanner = bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if currentLineCount >= linetoStartPrinting  {
			fmt.Printf("%s\n\n",fileScanner.Text())
		}
		currentLineCount++
	}

	file.Close()

}