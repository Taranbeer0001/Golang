package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	ar := [5]int{2, 4, 6, 8, 10}

	fmt.Println(arr)
	fmt.Println()
	fmt.Printf("%T", (arr))
	fmt.Println()
	fmt.Println(ar)

}
