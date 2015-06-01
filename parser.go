package exparser

import (
	"fmt"
	"strconv"

	"github.com/blake-wilson/exparser/functions"
	"github.com/blake-wilson/exparser/types"
)

var opa = map[string]struct {
	prec   int
	rAssoc bool
}{
	"^": {4, true},
	"*": {3, false},
	"/": {3, false},
	"+": {2, false},
	"-": {2, false},
}

// parse an expression and build and AST which can be evaluated
func EvalExpression(expr string) (types.AstNode, error) {
	tokens := tokenize(expr)
	tokens = tokenizePostfix(tokens)
	return evaluatePostfix(tokens)
}

func tokenize(expr string) []string {
	// values which start with a number will be parsed into number tokens,
	// values which start with a letter will be parsed into variable
	// tokens, and variables which are operators will be parsed into
	// operator tokens
	var tokens []string
	var token string
	for pos := 0; pos < len(expr); {
		if isWhitespace(expr[pos]) {
			pos++
		}
		if isDigit(expr[pos]) {
			token, pos = tokenizeNumber(expr, pos)
		} else if isLetter(expr[pos]) {
			token, pos = tokenizeVariable(expr, pos)
		} else if isOperator(string(expr[pos])) {
			token, pos = tokenizeOperator(expr, pos)
		}
		tokens = append(tokens, token)
	}
	return tokens
}

func tokenizeNumber(expr string, pos int) (string, int) {
	var res string
	var i int
	for i = pos; i < len(expr) && isDigit(expr[i]); i++ {
		res += string(expr[i])
	}
	return res, i
}

func tokenizeVariable(expr string, pos int) (string, int) {
	var res string
	var i int
	for i = pos; i < len(expr) && (isDigit(expr[i]) || isLetter(expr[i])); i++ {
		res += string(expr[i])
	}
	return res, i
}

func tokenizeOperator(expr string, pos int) (string, int) {
	if len(expr) > pos+2 {
		if isOperator(string(expr[pos])) && !isOperator(string(expr[pos+1])) {
			return string(expr[pos]), pos + 1
		} else {
			// cannot have two operators in a row
			return "", pos
		}
	} else {
		return string(expr[pos]), pos + 1
	}
}

// return a tokenized postfix expression
func tokenizePostfix(tokens []string) []string {

	// stacks for operators and operands
	var operators []string

	// postfix stack
	var postfix []string

	for _, tok := range tokens {
		switch tok {
		case "(":
			operators = append(operators, tok)
		case ")":
			var op string
			for {
				op, operators = operators[len(operators)-1], operators[:len(operators)-1]
				if op == "(" {
					break // discard "("
				}
				// add node
				postfix = append(postfix, op)
			}
		default:
			if o1, isOp := opa[tok]; isOp {
				// token is an operator
				for len(operators) > 0 {
					// consider top item on stack
					op := operators[len(operators)-1]
					if o2, isOp := opa[op]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rAssoc {
						break
					}

					operators = operators[:len(operators)-1] // pop operator from stack
					postfix = append(postfix, op)

				}
				// push operator to stack
				operators = append(operators, tok)

			} else { // token is an operand
				postfix = append(postfix, tok)
			}
		}
	}
	for len(operators) > 0 {
		postfix = append(postfix, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}
	return postfix
}

// Generate an AST based on a postfix expression
func evaluatePostfix(expr []string) (types.AstNode, error) {
	var evalStack []types.AstNode

	// slice of the variable names parsed in the expression
	var varNames []string

	for i := 0; i < len(expr); i++ {
		if len(expr[i]) < 1 {
			continue
		}
		if isOperator(expr[i]) {
			if len(evalStack) < 2 {
				// cannot pop two elements from stack
				return nil, fmt.Errorf("evaluation stack is in bad state")
			}
			rhs := evalStack[len(evalStack)-1]
			lhs := evalStack[len(evalStack)-2]

			// Pop top two elements off stack
			evalStack = evalStack[:len(evalStack)-2]

			// Create a function node for operators
			var node types.AstNode
			if nodeFunc, ok := functions.FMap[expr[i]]; ok {
				node = types.NewFunctionNode(lhs, rhs, nodeFunc)
			} else {
				// Not a valid operator
				return nil, fmt.Errorf("%s is not a valid operator", expr[i])
			}
			evalStack = append(evalStack, node)
		} else {
			// Non-operators are assumed to be variables or
			// terminals. Variables start with a letter. Terminals
			// are simply numbers
			var node types.AstNode
			if isLetter(expr[i][0]) {
				// Make a variable node
				node = types.NewVariableNode(expr[i])
				varNames = append(varNames, expr[i])

			} else {
				// try to parse a number/ terminal
				if value, err := strconv.ParseFloat(expr[i], 64); err == nil {
					// make a new Terminal
					node = types.NewTerminalNode(value)
				} else {
					// error parsing number
					return nil, fmt.Errorf("error parsing number")
				}
			}

			evalStack = append(evalStack, node)
		}
	}
	if len(evalStack) > 1 {
		// error
		return nil, fmt.Errorf("Error evaluating postfix expression")
	}
	return evalStack[0], nil
}

func isOperator(val string) bool {
	return val == "+" || val == "-" || val == "*" || val == "/" || val == "^"
}

func isDigit(char uint8) bool {
	return char >= '0' && char <= '9'
}

func isLetter(char uint8) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isWhitespace(char uint8) bool {
	return char == ' ' || char == '\t' || char == '\n'
}
