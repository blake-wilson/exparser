package eqparser

import (
	"math"
	"strconv"
	"strings"
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

// return a tokenized postfix expression
func ParseExpr(expr string) []string {

	// stacks for operators and operands
	var operators []string

	// postfix stack
	var postfix []string

	for _, tok := range strings.Fields(expr) {
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

// evaluate postfix expression
func EvaluatePostfix(expr []string) float64 {
	var evalStack []float64

	for i := 0; i < len(expr); i++ {
		if isOperator(expr[i]) {
			rhs := evalStack[len(evalStack)-1]
			lhs := evalStack[len(evalStack)-2]
			evalStack = evalStack[:len(evalStack)-2]
			evalStack = append(evalStack, evaluate(expr[i], lhs, rhs))
		} else {
			exprF, _ := strconv.ParseFloat(expr[i], 64)
			evalStack = append(evalStack, exprF)
		}
	}
	if len(evalStack) > 1 {
		// error
		return -1
	}
	return evalStack[0]
}

func isOperator(val string) bool {
	return val == "+" || val == "-" || val == "*" || val == "/" || val == "^"
}

func evaluate(operator string, lhs, rhs float64) float64 {
	if operator == "+" {
		return lhs + rhs
	} else if operator == "-" {
		return lhs - rhs
	} else if operator == "*" {
		return lhs * rhs
	} else if operator == "/" {
		return lhs / rhs
	} else if operator == "^" {
		return math.Pow(lhs, rhs)
	}
	return -1
}

func isDigit(char uint8) bool {
	return char >= '0' && char <= '9'
}
