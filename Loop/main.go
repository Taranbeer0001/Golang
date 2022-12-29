package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//1. Using the for loop print 1 to 100 numbers

	var i int
	fmt.Println("Using the for loop printing numbers 1 to 100")
	for i = 1; i <= 100; i++ {
		fmt.Println(i)
	}

	//2. Create a for loop using this syntax
	//for condition { }
	//print out the odd numbers in 1 to 50.

	fmt.Println("Using for loop printing all the odd numbers between 1 to 50")
	for j := 1; j <= 50; j++ {

		if j%2 != 0 {
			fmt.Println(j)
		}
	}

	//3. Create a for loop using this syntax
	//for { }
	//print out the even numbers in 1 to 50.
	fmt.Println("Using for loop printing all the even numbers between 1 to 50")
	for k := 1; k <= 50; k++ {

		if k%2 == 0 {
			fmt.Println(k)
		}
	}

	//4. Print out the quotient for each number between 50 and 105 when it is divided by 6.

	fmt.Println("Print out the quotient for each number between 50 and 105 when it is divided by 6")
	for l := 50; l <= 105; l++ {

		m := l / 6
		fmt.Println(m)
	}

	//5. Read the user input.
	//If the input is equal to Golang tutorial print welcome else print end

	fmt.Println("Enter the input")

	comp := "Golang tutorial"

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	if input == comp {
		fmt.Println("Welcome")
	} else {
		fmt.Println("End")
	}

	//6. Write a program to print the numbers from 1 to 80.
	//But, for multiples of two print Golang instead of the number and for the multiples of four print tutorial.
	//For numbers which are multiples of both two and four print Golang tutorial.

	fmt.Println("Print the numbers from 1 to 80 with certain conditions.")
	for k := 1; k <= 80; k++ {

		if k%2 == 0 {
			if k%4 == 0 {
				fmt.Println("Golang tutorial")
			} else {
				fmt.Println("Golang")
			}
		} else {
			fmt.Println(k)
		}

	}

}
