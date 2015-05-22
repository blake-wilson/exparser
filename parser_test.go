package eqparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExpr(t *testing.T) {
	tokens := ParseExpr("2 + 4 * 7 - 10 ^ 2")
	res, err := EvaluatePostfix(tokens)
	assert.Nil(t, err)
	assert.Equal(t, -70, res)
}
