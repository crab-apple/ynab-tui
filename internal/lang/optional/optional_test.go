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

func TestOrShouldReturnDefaultValueIfEmpty(t *testing.T) {
	// Given
	optional := Empty[string]()

	// When
	result := optional.Or("default")

	// Then
	assert.Equal(t, "default", result)
}

func TestShouldCreateEmptyFromNilPointer(t *testing.T) {
	// Given
	var pointer *string = nil

	// When
	result := OfPointer(pointer)

	// Then
	_, present := result.Get()
	assert.False(t, present)
}

func TestShouldCreatePresentFromNonNilPointer(t *testing.T) {
	// Given
	str := "hello"
	var pointer *string = &str

	// When
	result := OfPointer(pointer)

	// Then
	value, present := result.Get()
	assert.True(t, present)
	assert.Equal(t, "hello", value)
}
