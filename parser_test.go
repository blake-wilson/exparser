package exparser

import (
	"testing"

	"github.com/blake-wilson/exparser/types"
	"github.com/stretchr/testify/assert"
)

func TestParseExpr(t *testing.T) {
	res, err := EvalExpression("2 + 4 * 7 - 10 ^ 2")
	assert.NoError(t, err)
	assert.Equal(t, float64(-70), res.Eval(nil))

	res, err = EvalExpression("3 * 10 + 2 + 8 * 9")
	assert.NoError(t, err)
	assert.Equal(t, float64(104), res.Eval(nil))
}

func TestVariables(t *testing.T) {
	ctx := types.NewContext()
	ctx.AssignVariable("x", 3)
	res, err := EvalExpression("10 * x")

	assert.NoError(t, err)
	assert.Equal(t, float64(30), res.Eval(ctx))

	ctx.AssignVariable("y", 4)
	res, err = EvalExpression("x + y")
	assert.NoError(t, err)
	assert.Equal(t, float64(7), res.Eval(ctx))

	ctx.AssignVariable("z", 5)
	res, err = EvalExpression("x*y^z")
	assert.NoError(t, err)
	assert.Equal(t, float64(3072), res.Eval(ctx))
}

func TestParentheses(t *testing.T) {
	ctx := types.NewContext()
	ctx.AssignVariable("x", 7)
	res, err := EvalExpression("3*(x+2)")

	assert.Equal(t, float64(27), res.Eval(ctx))
	assert.NoError(t, err)
}

func TestDecimals(t *testing.T) {
	ctx := types.NewContext()
	ctx.AssignVariable("x", 2)

	res, err := EvalExpression("1.5*x")
	assert.Equal(t, float64(3), res.Eval(ctx))
	assert.NoError(t, err)
}

func TestTokenize(t *testing.T) {
	tokens, err := tokenize("3*x + 2")
	assert.NoError(t, err)
	assert.Equal(t, []string{"3", "*", "x", "+", "2"}, tokens)
}
