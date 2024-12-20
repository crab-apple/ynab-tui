package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateWithPresentValue(t *testing.T) {

	// When
	optional := Of("foo")
	value, present := optional.Get()

	// Then
	assert.Equal(t, "foo", value)
	assert.True(t, present)
}

func TestShouldCreateWithEmptyValue(t *testing.T) {

	// When
	optional := Empty[string]()
	_, present := optional.Get()

	// Then
	assert.False(t, present)
}

func TestOrShouldReturnValueIfPresent(t *testing.T) {
	// Given
	optional := Of("foo")

	// When
	result := optional.Or("default")

	// Then
	assert.Equal(t, "foo", result)
}

func TestOrShouldReturnDefaultValueIfEmpry(t *testing.T) {
	// Given
	optional := Empty[string]()

	// When
	result := optional.Or("default")

	// Then
	assert.Equal(t, "default", result)
}
