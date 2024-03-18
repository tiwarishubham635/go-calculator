package main

import (
	"fmt"
	"github.com/tiwarishubham635/math"
)

func main() {
	var num1 int
	var num2 int

	_, err1 := fmt.Scan(&num1)
	_, err := fmt.Scan(&num2)
	if err1 != nil || err != nil {
		fmt.Println("Error1: ", err1)
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(math.Add(num1, num2))
	fmt.Println(math.Subtract(num1, num2))
	fmt.Println(math.Multiply(num1, num2))
	fmt.Println(math.Divide(num1, num2))
}
