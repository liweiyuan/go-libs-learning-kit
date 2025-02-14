package cgo

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestAdd(t *testing.T) {
	result := Add(3, 4)
	assert.Equal(t, 7, result, "they should be equal")
}
