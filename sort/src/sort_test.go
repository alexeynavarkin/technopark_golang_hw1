package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sliceStringCompare(left, right []string) bool {
	if len(left) != len(right) {
		return false
	}
	for idx, elem := range left {
		if elem != right[idx] {
			return false
		}
	}
	return true
}

func TestParseNumPositive(t *testing.T) {
	test := []string{
		"123lol",
		"-1sd",
		"0asdf",
		"50000",
		"-259",
		"123l123",
	}
	wanted := []int{
		123,
		-1,
		0,
		50000,
		-259,
		123,
	}
	for idx, str := range test {
		num, err := parseNum(str)
		assert.Nil(t, err)
		assert.Equal(t, num, wanted[idx], "Wrong parse leading int")
	}
}

func TestCustomSort(t *testing.T) {
	test := [][]string{
		{
			"3",
			"2",
			"11",
		},
		{
			"3",
			"2",
			"11",
		},
		{
			"3",
			"2",
			"11",
		},
		{
			"1 3",
			"2 2",
			"3 1",
		},
		{
			"1 3",
			"2 2",
			"3 1",
		},
		{
			"3 11",
			"2 2",
			"1 3",
		},
		{
			"a",
			"A",
		},
	}

	wanted := [][]string{
		{
			"11",
			"2",
			"3",
		},
		{
			"3",
			"2",
			"11",
		},
		{
			"2",
			"3",
			"11",
		},
		{
			"1 3",
			"2 2",
			"3 1",
		},
		{
			"3 1",
			"2 2",
			"1 3",
		},
		{
			"3 11",
			"1 3",
			"2 2",
		},
		{
			"a",
			"A",
		},
	}

	sorts := []CustomSort{
		{},
		{
			reverse: true,
		},
		{
			numerical: true,
		},
		{
			columnNum: 1,
			reverse:   true,
		},
		{
			columnNum: 1,
			reverse:   false,
		},
		{
			columnNum: 1,
			reverse:   true,
			numerical: true,
		},
		{
			ignoreCase: true,
		},
	}

	for idx, conf := range sorts {
		conf.strings = test[idx]
		sort.Sort(conf)
		if !sliceStringCompare(test[idx], wanted[idx]) {
			t.Errorf("Expected:\n%v\nGot:\n%v\nOn test %v\n", wanted[idx], test[idx], idx)
		}
	}
}
