package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertAction(t *testing.T) {
	// assert equal
	assert.Equal(t, 1, 1, "they should be equal")
	// assert Null
	assert.Nil(t, nil)
	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")
}
