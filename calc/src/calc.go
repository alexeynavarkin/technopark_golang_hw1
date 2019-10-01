package main

import (
	"errors"
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
	switch node.(type) {
	case (*ast.BinaryExpr):
		binaryExpr, _ := node.(*ast.BinaryExpr)
		result, err := calcBinary(binaryExpr)
		return result, err
	case (*ast.BasicLit):
		basicLit, _ := node.(*ast.BasicLit)
		return strconv.Atoi(basicLit.Value)
	case (*ast.ParenExpr):
		parentExpr, _ := node.(*ast.ParenExpr)
		return evalExpr(parentExpr.X)
	default:
		return 0, ErrorUnexpectedToken
	}
}

// EvalFromString evals mathematical expression from string
func EvalFromString(expression string) (int, error) {
	astRoot, err := parser.ParseExpr(expression)
	if err != nil {
		return 0, err
	}
	return evalExpr(astRoot)
}
