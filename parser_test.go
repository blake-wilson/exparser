package exparser

import (
	"fmt"
	"testing"

	"github.com/blake-wilson/exparser/types"
	"github.com/stretchr/testify/assert"
)

func TestParseExpr(t *testing.T) {
	res, err := EvalExpression("2 + 4 * 7 - 10 ^ 2")
	assert.Nil(t, err)
	assert.Equal(t, float64(-70), res.Eval(nil))

	res, err = EvalExpression("3 * 10 + 2 + 8 * 9")
	assert.Nil(t, err)
	assert.Equal(t, float64(104), res.Eval(nil))
}

func TestVariables(t *testing.T) {
	ctx := types.NewContext()
	ctx.AssignVariable("x", 3)
	res, err := EvalExpression("10 * x")

	assert.Nil(t, err)
	assert.Equal(t, float64(30), res.Eval(ctx))

	ctx.AssignVariable("y", 4)
	res, err = EvalExpression("x + y")
	assert.Nil(t, err)
	fmt.Printf("Context %s\n\n", ctx)
	assert.Equal(t, float64(7), res.Eval(ctx))
}
