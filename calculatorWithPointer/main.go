package main

import "fmt"

func Add(x *int, y *int) (int, int, int) {
	return *x, *y, *x + *y
}
func Sub(x *int, y *int) (int, int, int) {
	return *x, *y, *x - *y
}
func Mul(x *int, y *int) (int, int, int) {
	return *x, *y, *x * *y
}
func Div(x *int, y *int) (int, int, int) {
	return *x, *y, *x / *y
}
func Mod(x *int, y *int) (int, int, int) {
	return *x, *y, *x % *y
}

func main() {
	var x int = 100
	var y int = 10
	var a, b, c int

	fmt.Println("Addition of two numbers : ")
	a, b, c = Add(&x, &y)
	fmt.Println(a, "+", b, "=", c)

	fmt.Println("Subtraction of two numbers : ")
	a, b, c = Sub(&x, &y)
	fmt.Println(a, "-", b, "=", c)

	fmt.Println("Multiplication of two numbers : ")
	a, b, c = Mul(&x, &y)
	fmt.Println(a, "*", b, "=", c)

	fmt.Println("Division of two numbers : ")
	a, b, c = Div(&x, &y)
	fmt.Println(a, "/", b, "=", c)

	fmt.Println("Modular of two numbers : ")
	a, b, c = Mod(&x, &y)
	fmt.Println(a, "%", b, "=", c)
}

//Calculator
