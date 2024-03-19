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

func TestDivisionHappyCase(t *testing.T) {
	num1 := 14
	num2 := 2
	quo, err := Divide(num1, num2)
	assert.Nil(t, err)
	assert.Equal(t, quo, 7)
}

func TestDivisionSadCase(t *testing.T) {
	num1 := 4
	num2 := 0
	quo, err := Divide(num1, num2)
	assert.NotNil(t, err)
	assert.Equal(t, quo, -1)
}
