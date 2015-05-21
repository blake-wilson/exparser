package eqparser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExpr(t *testing.T) {
	tokens := ParseExpr("2 + 4 * 7 - 10 ^ 2")
	assert.Equal(t, -70, EvaluatePostfix(tokens))
	fmt.Printf("%f", EvaluatePostfix(tokens))
}
