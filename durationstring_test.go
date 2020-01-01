package durationstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_Parses(t *testing.T) {
	years, months, days, hours, minutes, seconds, milliseconds, microseconds, nanoseconds, err := Parse("1y2mo3d4h5m6s7ms8us9ns")

	assert.Nil(t, err)
	assert.Equal(t, 1, years)
	assert.Equal(t, 2, months)
	assert.Equal(t, 3, days)
	assert.Equal(t, 4, hours)
	assert.Equal(t, 5, minutes)
	assert.Equal(t, 6, seconds)
	assert.Equal(t, 7, milliseconds)
	assert.Equal(t, 8, microseconds)
	assert.Equal(t, 9, nanoseconds)
}

func TestParse_AnyOrder_Parses(t *testing.T) {
	years, months, days, hours, minutes, seconds, milliseconds, microseconds, nanoseconds, err := Parse("42h1y")

	assert.Nil(t, err)
	assert.Equal(t, 1, years)
	assert.Equal(t, 0, months)
	assert.Equal(t, 0, days)
	assert.Equal(t, 42, hours)
	assert.Equal(t, 0, minutes)
	assert.Equal(t, 0, seconds)
	assert.Equal(t, 0, milliseconds)
	assert.Equal(t, 0, microseconds)
	assert.Equal(t, 0, nanoseconds)
}

func TestParse_WithWhitespace_Parses(t *testing.T) {
	years, months, days, hours, minutes, seconds, milliseconds, microseconds, nanoseconds, err := Parse("1y 2mo	3d\n4h\r5m6s7ms8us9ns\n")

	assert.Nil(t, err)
	assert.Equal(t, 1, years)
	assert.Equal(t, 2, months)
	assert.Equal(t, 3, days)
	assert.Equal(t, 4, hours)
	assert.Equal(t, 5, minutes)
	assert.Equal(t, 6, seconds)
	assert.Equal(t, 7, milliseconds)
	assert.Equal(t, 8, microseconds)
	assert.Equal(t, 9, nanoseconds)
}

func TestParse_MissingDigit_ReturnsError(t *testing.T) {
	_, _, _, _, _, _, _, _, _, err := Parse("y42h")

	assert.NotNil(t, err)
	assert.Equal(t, "Digit not supplied for unit 'y'", err.Error())
}

func TestParse_MissingUnit_ReturnsError(t *testing.T) {
	_, _, _, _, _, _, _, _, _, err := Parse("42")

	assert.NotNil(t, err)
	assert.Equal(t, "Unit not supplied for digit '42'", err.Error())
}

func TestParse_InvalidDigit_ReturnsError(t *testing.T) {
	_, _, _, _, _, _, _, _, _, err := Parse("999999999999999999999999999y")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Failed to parse digit")
}

func TestParse_InvalidUnit_ReturnsError(t *testing.T) {
	_, _, _, _, _, _, _, _, _, err := Parse("4invalid")

	assert.NotNil(t, err)
	assert.Equal(t, "invalid unit 'invalid'", err.Error())
}

func TestString_ReturnsExpected(t *testing.T) {
	s := String(1, 2, 3, 4, 5, 6, 7, 8, 9)

	assert.Equal(t, "1y2mo3d4h5m6s7ms8µS9ns", s)
}

func TestString_ExcludesInvalid(t *testing.T) {
	s := String(1, 0, 0, 4, 0, 0, 0, 0, 0)

	assert.Equal(t, "1y4h", s)
}
