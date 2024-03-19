package advanced_math

import "errors"

var DecimalPlaces = 5

func Sin(x float64) float64 {
	ans := x - Power(x, 3)/float64(Factorial(3)) +
		Power(x, 5)/float64(Factorial(5)) -
		Power(x, 7)/float64(Factorial(7)) +
		Power(x, 9)/float64(Factorial(9)) -
		Power(x, 11)/float64(Factorial(11))
	return RoundToDecimal(ans, DecimalPlaces)
}

func Cos(x float64) float64 {
	ans := 1 - Power(x, 2)/float64(Factorial(2)) +
		Power(x, 4)/float64(Factorial(4)) -
		Power(x, 6)/float64(Factorial(6)) +
		Power(x, 8)/float64(Factorial(8)) -
		Power(x, 10)/float64(Factorial(10))
	return RoundToDecimal(ans, DecimalPlaces)
}

func Tan(x float64) (float64, error) {
	if Cos(x) == 0 {
		return -1, errors.New("Not defined")
	}
	return RoundToDecimal(Sin(x) / Cos(x)), nil
}
