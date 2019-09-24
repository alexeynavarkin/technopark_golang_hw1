package main

import (
	"bufio"
	"io"
	"strings"
)

// ReadSliceStrings is function to read strings divided '/n' from stdin to slice
func ReadSliceStrings(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	var strings []string
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}
	return strings
}

// WriteSliceStrings is a function to write slice of strings
func WriteSliceStrings(writer io.Writer, lines []string, ignoreCase, uniq bool) {
	var prevStr string
	for _, str := range lines {
		if ignoreCase {
			str = strings.ToLower(str)
		}

		if uniq && prevStr == str {
			continue
		} else {
			io.WriteString(writer, str+"\n")
		}

		if ignoreCase {
			prevStr = strings.ToLower(str)
			continue
		}
		prevStr = str
	}
}
