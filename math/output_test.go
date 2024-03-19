package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	num1 := 3
	num2 := 4
	sum := Add(num1, num2)
	assert.Equal(t, sum, 7)
}

func TestSubtract(t *testing.T) {
	num1 := 3
	num2 := 5
	diff := Subtract(num1, num2)
	assert.Equal(t, diff, -2)
}

func TestMultiply(t *testing.T) {
	num1 := 7
	num2 := 4
	prod := Multiply(num1, num2)
	assert.Equal(t, prod, 28)
}

func TestIntegerDivisionHappyCase(t *testing.T) {
	num1 := 15
	num2 := 2
	quo, rem, err := IntegerDivision(num1, num2)
	assert.Nil(t, err)
	assert.Equal(t, rem, 1)
	assert.Equal(t, quo, 7)
}

func TestIntegerDivisionSadCase(t *testing.T) {
	num1 := 4
	num2 := 0
	quo, rem, err := IntegerDivision(num1, num2)
	assert.NotNil(t, err)
	assert.Equal(t, rem, -1)
	assert.Equal(t, quo, -1)
}

func TestFloatDivisionHappyCase(t *testing.T) {
	num1 := 15
	num2 := 2
	quo, err := FloatDivision(num1, num2)
	assert.Nil(t, err)
	assert.Equal(t, quo, 7)
}

func TestFloatDivisionSadCase(t *testing.T) {
	num1 := 4
	num2 := 0
	quo, err := FloatDivision(num1, num2)
	assert.NotNil(t, err)
	assert.Equal(t, quo, 0)
}

func TestModulo(t *testing.T) {
	num1 := 5
	num2 := 3
	mod := Modulo(num1, num2)
	assert.Equal(t, mod, 2)
}
