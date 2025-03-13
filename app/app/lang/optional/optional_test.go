package optional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOfShouldCreateWithPresentValue(t *testing.T) {

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

func TestOfPointerShouldCreateEmptyFromNilPointer(t *testing.T) {
	// Given
	var pointer *string = nil

	// When
	result := OfPointer(pointer)

	// Then
	_, present := result.Get()
	assert.False(t, present)
}

func TestOfPointerShouldCreatePresentFromNonNilPointer(t *testing.T) {
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

func TestAsPointerShouldReturnNilPointerIfEmpty(t *testing.T) {
	// Given
	opt := Empty[string]()

	// When
	pointer := opt.AsPointer()

	// Then
	assert.Nil(t, pointer)
}
func TestAsPointerShouldReturnPointerToValueIfPresent(t *testing.T) {
	// Given
	opt := Of("hello")

	// When
	pointer := opt.AsPointer()

	// Then
	assert.Equal(t, "hello", *pointer)
}

type StrContainer struct {
	myString string
}

func TestMapShouldReturnEmptyOptionalIfEmpty(t *testing.T) {
	// Given
	opt := Empty[StrContainer]()

	// When
	result := Map(opt, func(x StrContainer) string { return x.myString })

	// Then
	_, present := result.Get()
	assert.False(t, present)
}

func TestMapShouldReturnOptionalWithResultIfPresent(t *testing.T) {
	// Given
	opt := Of[StrContainer](StrContainer{myString: "hello"})

	// When
	result := Map(opt, func(x StrContainer) string { return x.myString })

	// Then
	value, present := result.Get()
	assert.True(t, present)
	assert.Equal(t, "hello", value)
}
