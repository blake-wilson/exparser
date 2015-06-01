package exparser

import (
	"testing"

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
