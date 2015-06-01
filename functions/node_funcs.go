package functions

import (
	"math"

	"github.com/blake-wilson/exparser/types"
)

var FMap = map[string]func(leftChild, rightChild types.AstNode, ctx *types.Context) float64{
	"+": AddFunction,
	"-": SubtractFunction,
	"*": MultiplyFunction,
	"/": DivideFunction,
	"^": PowerFunction,
}

func AddFunction(leftChild, rightChild types.AstNode, ctx *types.Context) float64 {
	return leftChild.Eval(ctx) + rightChild.Eval(ctx)
}

func SubtractFunction(leftChild, rightChild types.AstNode, ctx *types.Context) float64 {
	return leftChild.Eval(ctx) - rightChild.Eval(ctx)
}

func MultiplyFunction(leftChild, rightChild types.AstNode, ctx *types.Context) float64 {
	return leftChild.Eval(ctx) * rightChild.Eval(ctx)
}

func DivideFunction(leftChild, rightChild types.AstNode, ctx *types.Context) float64 {
	return leftChild.Eval(ctx) / rightChild.Eval(ctx)
}

func PowerFunction(leftChild, rightChild types.AstNode, ctx *types.Context) float64 {
	return math.Pow(leftChild.Eval(ctx), rightChild.Eval(ctx))
}
