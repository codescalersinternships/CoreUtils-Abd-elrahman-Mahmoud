package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var numberofLinesFlag bool
	var numberofWordsFlag bool
	var numberofCharactersFlag bool
	var err error
	var fileName string
	var file *os.File

	numberofLines := 0
	numberofWords := 0
	numberofCharacters := 0

	flag.BoolVar(&numberofLinesFlag, "l", false, "Number of Lines")
	flag.BoolVar(&numberofWordsFlag, "w", false, "Number of Words")
	flag.BoolVar(&numberofCharactersFlag, "c", false, "Number of Characters")

	flag.Parse()

	if !numberofLinesFlag && !numberofWordsFlag && !numberofCharactersFlag {
		numberofLinesFlag = true
		numberofWordsFlag = true
		numberofCharactersFlag = true
	}

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Wrong Number of Arguments")
	} else {
		fileName = args[0]
	}

	if numberofLinesFlag {
		file, err = os.Open(fileName)
		if err != nil {
			fmt.Println("Error trying to open file!")
			panic(err)
		}

		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			numberofLines++
		}
		fmt.Println("Number of Lines :", numberofLines)

		defer file.Close()
	}
	if numberofWordsFlag {
		file, err = os.Open(fileName)
		if err != nil {
			fmt.Println("Error trying to open file!")
			panic(err)
		}

		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanWords)
		for fileScanner.Scan() {
			numberofWords++
		}
		fmt.Println("Number of Words :", numberofWords)

		defer file.Close()
	}
	if numberofCharactersFlag {
		file, err = os.Open(fileName)
		if err != nil {
			fmt.Println("Error trying to open file!")
			panic(err)
		}

		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanRunes)
		for fileScanner.Scan() {
			numberofCharacters++
		}
		fmt.Println("Number of Characters :", numberofCharacters)

		defer file.Close()
	}

}
