package mock

import (
	"fmt"
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

	// assert true
	assert.True(t, true, "it should be true")

	// assert false
	assert.False(t, false, "it should be false")

	// assert contains
	assert.Contains(t, "hello world", "world", "\"hello world\" should contain \"world\"")

	// assert not contains
	assert.NotContains(t, "hello world", "planet", "\"hello world\" should not contain \"planet\"")

	// assert len
	assert.Len(t, []int{1, 2, 3}, 3, "slice should have length 3")

	// assert empty
	assert.Empty(t, []int{}, "slice should be empty")

	// assert not empty
	assert.NotEmpty(t, []int{1}, "slice should not be empty")

	// assert type of
	assert.IsType(t, 123, int(0), "value should be of type int")

	// assert no error
	assert.NoError(t, nil, "there should be no error")

	// assert error
	assert.Error(t, fmt.Errorf("an error occurred"), "there should be an error")

	// assert panic
	assert.Panics(t, func() { panic("a problem occurred") }, "the code should panic")

	// assert no panic
	assert.NotPanics(t, func() {}, "the code should not panic")
}
