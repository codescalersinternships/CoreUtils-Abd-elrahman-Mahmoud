package main

import (
	"flag"
    "fmt"
	"strconv"
	"os"
	"bufio"
)

func main() {
	var NumberofLinesFlag bool
	var NumberofLines int
	var err error
	var fileName string
	var file *os.File

    flag.BoolVar(&NumberofLinesFlag, "n", false, "Number of Lines")

	flag.Parse()

	args := flag.Args()
	if NumberofLinesFlag && len(args) == 2 {
		NumberofLines, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error trying to convert string to integer!")
		    panic(err)
		}
		fileName = args[1]
	} else if len(args) == 1 {
		NumberofLines = 10
		fileName = args[0]
	} else {
		fmt.Println("Wrong Number of Arguments")
	}

	file, err = os.Open(fileName)
	if err != nil {
		fmt.Println("Error trying to open file!")
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
 
	for fileScanner.Scan() {
		if NumberofLines == 0 {
			break
		}
		fmt.Printf("%s\n\n",fileScanner.Text())
		NumberofLines--
	}

	file.Close()

}