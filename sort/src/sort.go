package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// CustomSort is implementation of custom sort logic
type CustomSort struct {
	reverse    bool
	ignoreCase bool
	numerical  bool
	columnNum  int
	strings    []string
}

func (c CustomSort) Len() int {
	return len(c.strings)
}

func (c CustomSort) Swap(i, j int) {
	c.strings[i], c.strings[j] = c.strings[j], c.strings[i]
}

func getColumn(str string, n int) string {
	words := strings.Split(str, " ")
	if n >= len(words) {
		return ""
	}
	return words[n]
}

// ErrorParse is error displaying that string has no leading number
var ErrorParse = errors.New("val: can not parse leading number")

var numRegexp = regexp.MustCompile(`^-?[0-9]+`)

func parseNum(str string) (int, error) {
	num := numRegexp.FindString(str)
	if num == "" {
		return 0, ErrorParse
	}
	integer, err := strconv.Atoi(num)
	// This check could be skipped because of regexp
	if err != nil {
		return 0, ErrorParse
	}
	return integer, nil
}

func lessNum(left, right string) bool {
	leftNum, _ := parseNum(left)
	rightNum, _ := parseNum(right)
	return leftNum < rightNum
}

func (c CustomSort) Less(i, j int) bool {
	var isLess bool

	left := c.strings[i]
	right := c.strings[j]

	if c.columnNum != 0 {
		left = getColumn(left, c.columnNum)
		right = getColumn(right, c.columnNum)
	}

	if c.ignoreCase {
		left = strings.ToLower(left)
		right = strings.ToLower(right)
	}

	if c.numerical {
		isLess = lessNum(left, right)
	} else {
		isLess = left < right
	}

	if c.reverse {
		isLess = !isLess
	}

	return isLess
}
