package cmp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIf(t *testing.T) {
	assert.True(t, If(true, 1, 2) == 1)
	assert.True(t, If(true, "a", "b") == "a")
	assert.True(t, IfNumber(true, 312, 3121) == 312)
}
