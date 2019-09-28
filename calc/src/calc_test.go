package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalFromString(t *testing.T) {
	test := []string{
		"2+2",
		"2+4*2",
		"(2-(2*(3)))",
		"    2 + 22 - 22 -  2    ",
		"2/2",
		"2/3",
	}
	wanted := []int{
		4,
		10,
		-4,
		0,
		1,
		0,
	}

	for idx, expr := range test {
		result, err := EvalFromString(expr)
		assert.Equal(t, result, wanted[idx], "Error on test %d, with expr %v. Wanted %v, got %d", idx, expr, wanted[idx], result)
		assert.Nil(t, err, "Got not nill err on test %d", idx)
	}
}
