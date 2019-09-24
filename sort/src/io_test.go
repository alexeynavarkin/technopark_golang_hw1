package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testWriter struct {
	inputs *[]byte
}

func (t testWriter) Write(p []byte) (n int, err error) {
	*t.inputs = append(*t.inputs, p...)
	return len(p), nil
}

func TestOutput(t *testing.T) {
	var inputs []byte
	writer := testWriter{
		inputs: &inputs,
	}

	test := []string{
		"a",
		"a",
		"A",
	}

	wanted := []byte{
		'a',
		'\n',
		'a',
		'\n',
		'A',
		'\n',
	}

	WriteSliceStrings(writer, test, false, false)
	assert.True(t, reflect.DeepEqual(*writer.inputs, wanted))
}

func TestOutputUniq(t *testing.T) {
	var inputs []byte

	writer := testWriter{
		inputs: &inputs,
	}

	test := []string{
		"a",
		"a",
		"A",
	}

	wanted := []byte{
		'a',
		'\n',
		'A',
		'\n',
	}
	WriteSliceStrings(writer, test, false, true)
	assert.True(t, reflect.DeepEqual(*writer.inputs, wanted))
}

func TestOutputUniqIgnoreCase(t *testing.T) {
	var inputs []byte

	writer := testWriter{
		inputs: &inputs,
	}

	test := []string{
		"a",
		"a",
		"A",
	}

	wanted := []byte{
		'a',
		'\n',
	}
	WriteSliceStrings(writer, test, true, true)
	assert.True(t, reflect.DeepEqual(*writer.inputs, wanted))
}
