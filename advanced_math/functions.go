package advanced_math

import (
	"errors"
	"github.com/tiwarishubham635/go-calculator/v3/math"
)

func Inverse[number math.Number](num number) (number, error) {
	if num == 0 {
		return 0, errors.New("not defined")
	}
	return math.FloatDivision(1, num)
}

func GCD(num1 int, num2 int) int {
	if num1 < num2 {
		return GCD(num2, num1)
	}
	if num2 == 0 {
		return num1
	}
	return GCD(num2, math.Modulo(num1, num2))
}

func LCM(num1 int, num2 int) (int, error) {
	product := math.Multiply(num1, num2)
	gcd := GCD(num1, num2)
	quo, err := math.FloatDivision(product, gcd)
	if err == nil {
		return quo, nil
	}
	return -1, err
}

func Power[number math.Number](a number, b int) number {
	if a == 0 {
		return 0
	}
	if b == 0 || a == 1 {
		return 1
	}
	mid := Power(a, b/2)
	exp := math.Multiply(mid, mid)
	if b%2 == 0 {
		return exp
	}
	return math.Multiply(exp, a)
}

func Factorial(num int) int {
	if num < 0 {
		return 0
	}
	if num == 0 || num == 1 {
		return 1
	}
	return math.Multiply(num, Factorial(math.Subtract(num, 1)))
}

var PI = 3.1415926

func AreaOfCircle[number math.Number](radius number) float64 {
	return math.Multiply(PI, float64(Power(radius, 2)))
}

func RoundToInt[number math.Number](num number) int {
	return int(float64(num) + 0.5)
}

func RoundToDecimal(num float64, places ...int) float64 {
	decimalPlaces := 2
	if len(places) > 0 {
		decimalPlaces = places[0]
	}
	factor := Power(10, decimalPlaces)
	quo, _ := math.FloatDivision(float64(RoundToInt(math.Multiply(num, float64(factor)))), float64(factor))
	return quo
}
