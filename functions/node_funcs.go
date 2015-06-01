package functions

import (
	"math"

	"github.com/blake-wilson/exparser/types"
)

var FMap = map[string]func(leftChild, rightChild types.AstNode) float64{
	"+": AddFunction,
	"-": SubtractFunction,
	"*": MultiplyFunction,
	"/": DivideFunction,
	"^": PowerFunction,
}

func AddFunction(leftChild, rightChild types.AstNode) float64 {
	return leftChild.Eval(nil) + rightChild.Eval(nil)
}

func SubtractFunction(leftChild, rightChild types.AstNode) float64 {
	return leftChild.Eval(nil) - rightChild.Eval(nil)
}

func MultiplyFunction(leftChild, rightChild types.AstNode) float64 {
	return leftChild.Eval(nil) * rightChild.Eval(nil)
}

func DivideFunction(leftChild, rightChild types.AstNode) float64 {
	return leftChild.Eval(nil) / rightChild.Eval(nil)
}

func PowerFunction(leftChild, rightChild types.AstNode) float64 {
	return math.Pow(leftChild.Eval(nil), rightChild.Eval(nil))
}
