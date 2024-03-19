package math

import "errors"

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func Subtract(num1 int, num2 int) int {
	return num1 - num2
}

func Multiply(num1 int, num2 int) int {
	return num1 * num2
}

func IntegerDivision(num1 int, num2 int) (int, int, error) {
	if num2 == 0 {
		return -1, -1, errors.New("cannot divide by zero")
	}
	return num1 / num2, num1 % num2, nil
}

func FloatDivision(num1 int, num2 int) (float64, error) {
	if num2 == 0 {
		return -1.0, errors.New("cannot divide by zero")
	}
	return float64(num1) / float64(num2), nil
}
