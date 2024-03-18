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

func Divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return -1, errors.New("cannot divide by zero")
	}
	return num1 / num2, nil
}
