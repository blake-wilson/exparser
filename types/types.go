package types

type NodeType int

const (
	Function NodeType = iota
	Terminal
)

type ScalarFunc func(time, value float64)

// Information needed for the evaluation of an expression
type Context struct {
	variables map[string]float64
}

// Ast which is composed of AstNodes and evaluated
type Ast struct {
	nodes []*AstNode
}

type AstNode interface {
	// Function to evaluate the AST Node
	Eval(ctx *Context) float64
}

type FunctionNode struct {
	leftChild, rightChild AstNode

	// Function which defines what the node does with its children
	nodeFunc func(leftChild, rightChild AstNode, ctx *Context) float64
}

type VariableNode struct {
	// VariableNode is not self-contained.  Requires context to provide the
	// value of the variable
	name string
}

type TerminalNode struct {
	value float64
}

func (n *FunctionNode) Eval(ctx *Context) float64 {
	return n.nodeFunc(n.leftChild, n.rightChild, ctx)
}

func (n *VariableNode) Eval(ctx *Context) float64 {
	if _, ok := ctx.variables[n.name]; ok {
		return ctx.variables[n.name]
	}
	// Variable did not map to a value
	return -1
}

func (n *TerminalNode) Eval(ctx *Context) float64 {
	return n.value
}

func NewFunctionNode(leftChild, rightChild AstNode, nodeFunc func(leftChild, rightChild AstNode, ctx *Context) float64) *FunctionNode {
	return &FunctionNode{
		leftChild:  leftChild,
		rightChild: rightChild,
		nodeFunc:   nodeFunc,
	}
}

func NewVariableNode(name string) *VariableNode {
	return &VariableNode{
		name: name,
	}
}

func NewTerminalNode(value float64) *TerminalNode {
	return &TerminalNode{
		value: value,
	}
}
