package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

var ErrorUnknowOperation = errors.New("error got unknown operation token")

func calcBinary(binary *ast.BinaryExpr) (int, error) {
	opLeft, errLeft := evalExpr(binary.X)
	opRight, errRight := evalExpr(binary.Y)
	if errLeft != nil {
		return 0, errLeft
	}
	if errRight != nil {
		return 0, errRight
	}

	switch binary.Op {
	case token.ADD:
		return opLeft + opRight, nil
	case token.SUB:
		return opLeft - opRight, nil
	case token.MUL:
		return opLeft * opRight, nil
	case token.QUO:
		return opLeft / opRight, nil
	default:
		return 0, ErrorUnknowOperation
	}
}

var ErrorCastBinaryExpr = errors.New("error casting to binary expression")
var ErrorCastBasicLit = errors.New("error casting to basic literal")
var ErrorCastParentExpr = errors.New("error casting to parent expression")
var ErrorUnexpectedToken = errors.New("error got unexpected token")

func evalExpr(node interface{}) (int, error) {
	fmt.Printf("Calc get %T, %v\n", node, node)
	switch node.(type) {
	case (*ast.BinaryExpr):
		binaryExpr, ok := node.(*ast.BinaryExpr)
		if !ok {
			return 0, ErrorCastBinaryExpr
		}
		result, err := calcBinary(binaryExpr)
		return result, err
	case (*ast.BasicLit):
		basicLit, ok := node.(*ast.BasicLit)
		if !ok {
			return 0, ErrorCastBasicLit
		}
		return strconv.Atoi(basicLit.Value)
	case (*ast.ParenExpr):
		parentExpr, ok := node.(*ast.ParenExpr)
		if !ok {
			return 0, ErrorCastParentExpr
		}
		return evalExpr(parentExpr.X)
	default:
		return 0, ErrorUnexpectedToken
	}
}

func EvalFromString(expression string) (int, error) {
	astRoot, _ := parser.ParseExpr(expression)
	fs := token.NewFileSet()
	ast.Print(fs, astRoot)
	return evalExpr(astRoot)
}
