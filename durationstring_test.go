package durationstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse_Parses(t *testing.T) {
	d, err := Parse("1y2mo3d4h5m6s7ms8us9ns")

	assert.Nil(t, err)
	assert.Equal(t, 1, d.Years)
	assert.Equal(t, 2, d.Months)
	assert.Equal(t, 3, d.Days)
	assert.Equal(t, 4, d.Hours)
	assert.Equal(t, 5, d.Minutes)
	assert.Equal(t, 6, d.Seconds)
	assert.Equal(t, 7, d.Milliseconds)
	assert.Equal(t, 8, d.Microseconds)
	assert.Equal(t, 9, d.Nanoseconds)
}

func TestParse_AnyOrder_Parses(t *testing.T) {
	d, err := Parse("42h1y")

	assert.Nil(t, err)
	assert.Equal(t, 1, d.Years)
	assert.Equal(t, 0, d.Months)
	assert.Equal(t, 0, d.Days)
	assert.Equal(t, 42, d.Hours)
	assert.Equal(t, 0, d.Minutes)
	assert.Equal(t, 0, d.Seconds)
	assert.Equal(t, 0, d.Milliseconds)
	assert.Equal(t, 0, d.Microseconds)
	assert.Equal(t, 0, d.Nanoseconds)
}

func TestParse_WithWhitespace_Parses(t *testing.T) {
	d, err := Parse("1y 2mo	3d\n4h\r5m6s7ms8us9ns\n")

	assert.Nil(t, err)
	assert.Equal(t, 1, d.Years)
	assert.Equal(t, 2, d.Months)
	assert.Equal(t, 3, d.Days)
	assert.Equal(t, 4, d.Hours)
	assert.Equal(t, 5, d.Minutes)
	assert.Equal(t, 6, d.Seconds)
	assert.Equal(t, 7, d.Milliseconds)
	assert.Equal(t, 8, d.Microseconds)
	assert.Equal(t, 9, d.Nanoseconds)
}

func TestParse_MissingDigit_ReturnsError(t *testing.T) {
	_, err := Parse("y42h")

	assert.NotNil(t, err)
	assert.Equal(t, "Digit not supplied for unit 'y'", err.Error())
}

func TestParse_MissingUnit_ReturnsError(t *testing.T) {
	_, err := Parse("42")

	assert.NotNil(t, err)
	assert.Equal(t, "Unit not supplied for digit '42'", err.Error())
}

func TestParse_InvalidDigit_ReturnsError(t *testing.T) {
	_, err := Parse("999999999999999999999999999y")

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Failed to parse digit")
}

func TestParse_InvalidUnit_ReturnsError(t *testing.T) {
	_, err := Parse("4invalid")

	assert.NotNil(t, err)
	assert.Equal(t, "invalid unit 'invalid'", err.Error())
}

func TestString_ReturnsExpected(t *testing.T) {
	d := NewDuration(1, 2, 3, 4, 5, 6, 7, 8, 9)
	s := d.String()

	assert.Equal(t, "1y2mo3d4h5m6s7ms8µS9ns", s)
}

func TestString_ExcludesInvalid(t *testing.T) {
	d := NewDuration(1, 0, 0, 4, 0, 0, 0, 0, 0)
	s := d.String()

	assert.Equal(t, "1y4h", s)
}
