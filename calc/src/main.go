package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	expression, errReader := reader.ReadString('\n')
	if errReader != nil && errReader != io.EOF {
		fmt.Println(errReader)
		os.Exit(-1)
	}

	result, errEval := EvalFromString(expression)
	if errEval != nil {
		fmt.Println(errEval)
		os.Exit(-1)
	}

	fmt.Print(result)
}
