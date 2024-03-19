package advanced_math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInverseHappyCase(t *testing.T) {
	num1 := 4.0
	inv, err := Inverse(num1)
	assert.Nil(t, err)
	assert.Equal(t, inv, 0.25)
}

func TestInverseSadCase(t *testing.T) {
	num1 := 0
	inv, err := Inverse(num1)
	assert.NotNil(t, err)
	assert.Equal(t, inv, 0)
}

func TestGCDCoprime(t *testing.T) {
	num1 := 12
	num2 := 7
	gcd := GCD(num1, num2)
	assert.Equal(t, gcd, 1)
}

func TestGCDNonCoprime(t *testing.T) {
	num1 := 36
	num2 := 24
	gcd := GCD(num1, num2)
	assert.Equal(t, gcd, 12)
}

func TestLCM(t *testing.T) {
	num1 := 12
	num2 := 6
	lcm, err := LCM(num1, num2)
	assert.Nil(t, err)
	assert.Equal(t, lcm, 12)
}

func TestPower(t *testing.T) {
	num1 := 2.5
	num2 := 4
	exp := Power(num1, num2)
	assert.Equal(t, exp, 39.0625)
}

func TestFactorial(t *testing.T) {
	num := 5
	fact := Factorial(num)
	assert.Equal(t, fact, 120)
}

func TestAreaOfCircle(t *testing.T) {
	radius := 1
	area := AreaOfCircle(radius)
	assert.Equal(t, area, PI)
}

func TestRoundToInt(t *testing.T) {
	val := 1.49
	round := RoundToInt(val)
	assert.Equal(t, round, 1)

	val = 1.51
	round = RoundToInt(val)
	assert.Equal(t, round, 2)
}

func TestRoundToDecimal(t *testing.T) {
	val := 1.2345678
	round := RoundToDecimal(val)
	assert.Equal(t, round, 1.23)

	val = 1.8765432
	round = RoundToDecimal(val, 4)
	assert.Equal(t, round, 1.8765)
}
