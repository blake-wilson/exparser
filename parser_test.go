package exparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExpr(t *testing.T) {
	tokens := ParseExpr("2 + 4 * 7 - 10 ^ 2")
	res, err := EvaluatePostfix(tokens)
	assert.Nil(t, err)
	assert.Equal(t, float64(-70), res.Eval(nil))

	tokens = ParseExpr("3 * 10 + 2 + 8 * 9")
	res, err = EvaluatePostfix(tokens)
	assert.Nil(t, err)
	assert.Equal(t, float64(104), res.Eval(nil))
}
