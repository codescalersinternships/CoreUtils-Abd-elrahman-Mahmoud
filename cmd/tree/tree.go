package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var numberofDirectories int = 1
var numberofFiles int = 0

func printDirectory(entry string, level int) {
	fmt.Printf("|_%s\n", entry)
	printDirectoryRecursion(entry, level, "    ")
	fmt.Printf("%d directories, %d files\n", numberofDirectories, numberofFiles)
}

func printDirectoryRecursion(entryName string, level int, space string) {
	if level == 0 { //first exit condition
		return
	}

	c, err := os.ReadDir(entryName)
	if err != nil {
		fmt.Println("Error trying to read directory")
		panic(err)
	}

	if len(c) == 0 {
		return //second exit condition
	}
	level--

	for _, entry := range c {
		if entry.IsDir() {
			numberofDirectories++
			fmt.Printf("%s|_%s\n", space, entry.Name())
			printDirectoryRecursion(filepath.Join(entryName, entry.Name()), level, space+"    ")
		} else {
			numberofFiles++
			fmt.Printf("%s|_%s\n", space, entry.Name())
		}
	}
}

func main() {
	var treeLevelFlag bool
	var treeLevel int
	var directoryName string
	var err error

	flag.BoolVar(&treeLevelFlag, "L", false, "Tree level to stop at")

	flag.Parse()

	args := flag.Args()
	if treeLevelFlag && len(args) == 2 {
		treeLevel, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error trying to convert string to integer!")
			panic(err)
		}
		directoryName = args[1]
	} else if len(args) == 1 {
		treeLevel = -1
		directoryName = args[0]
	} else if len(args) == 0 {
		treeLevel = -1
		directoryName, err = os.Getwd()
		if err != nil {
			fmt.Println("Error trying to get current directory!")
			panic(err)
		}
	} else {
		log.Fatal("Wrong Number of Arguments")
	}

	printDirectory(directoryName, treeLevel)

}
