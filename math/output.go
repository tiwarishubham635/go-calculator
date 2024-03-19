package math

import "errors"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Add[number Number](num1 number, num2 number) number {
	return num1 + num2
}

func Subtract[number Number](num1 number, num2 number) number {
	return num1 - num2
}

func Multiply[number Number](num1 number, num2 number) number {
	return num1 * num2
}

func IntegerDivision(num1 int, num2 int) (int, int, error) {
	if num2 == 0 {
		return -1, -1, errors.New("cannot divide by zero")
	}
	return num1 / num2, num1 % num2, nil
}

func FloatDivision[number Number](num1 number, num2 number) (number, error) {
	if num2 == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return num1 / num2, nil
}

func Modulo(num1 int, num2 int) int {
	return num1 % num2
}
