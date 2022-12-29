package main

import (
	"fmt"

	"Assignments/Calculator"
)

func main() {
	var x int
	var y int
	var a, b, c int
	fmt.Println("Enter the first digit")
	fmt.Scanln(&x)
	fmt.Println("Enter the Second digit")
	fmt.Scanln(&y)
	fmt.Println("Addition of two numbers : ")
	a, b, c = Calculator.Add(x, y)
	fmt.Println(a, "+", b, "=", c)

	fmt.Println("Subtraction of two numbers : ")
	a, b, c = Calculator.Sub(x, y)
	fmt.Println(a, "-", b, "=", c)

	fmt.Println("Multiplication of two numbers : ")
	a, b, c = Calculator.Mul(x, y)
	fmt.Println(a, "*", b, "=", c)

	fmt.Println("Division of two numbers : ")
	a, b, c = Calculator.Div(x, y)
	fmt.Println(a, "/", b, "=", c)

	fmt.Println("Modular of two numbers : ")
	a, b, c = Calculator.Mod(x, y)
	fmt.Println(a, "%", b, "=", c)
}

//Calculator
