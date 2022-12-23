package main

import "fmt"

func updateThirdElement(x *string) {
	*x = "Texas"
	fmt.Println(*x)
	fmt.Println(x)

}

func main() {

	arr := [4]string{"tokyo", "florida", "dallas", "bali"}
	fmt.Println(arr)
	updateThirdElement(&arr[3])
	fmt.Println(arr)

}
