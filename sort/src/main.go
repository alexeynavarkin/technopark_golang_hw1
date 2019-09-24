package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	reverse := flag.Bool("r", false, "column to sort")
	ignoreCase := flag.Bool("f", false, "ignore case")
	numerical := flag.Bool("n", false, "numerical sort")
	columnNum := flag.Int("k", 0, "column to sort")
	uniq := flag.Bool("u", false, "do not repeat")
	outputFilePath := flag.String("o", "", "output to file")

	flag.Parse()

	inputFilePath := flag.Arg(0)

	var reader io.Reader
	if inputFilePath != "" {
		inputFile, err := os.Open(inputFilePath)
		if err != nil {
			fmt.Println("Error opening file to read.")
			os.Exit(-1)
		}
		defer inputFile.Close()
		reader = inputFile
	} else {
		reader = os.Stdin
	}

	var writer io.Writer
	if *outputFilePath != "" {
		outputFile, err := os.Create(*outputFilePath)
		if err != nil {
			fmt.Println("Error opening file to write.")
			os.Exit(-1)
		}
		defer outputFile.Close()
		writer = outputFile
	} else {
		writer = os.Stdout
	}

	strings := ReadSliceStrings(reader)

	customSort := CustomSort{
		strings:    strings,
		reverse:    *reverse,
		ignoreCase: *ignoreCase,
		numerical:  *numerical,
		columnNum:  *columnNum,
	}

	sort.Sort(customSort)

	WriteSliceStrings(writer, strings, *uniq, *ignoreCase)
}
