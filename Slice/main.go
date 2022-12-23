package main

import "fmt"

func main() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	//Question 2.

	fmt.Println(s)
	fmt.Printf("%T", (s))
	fmt.Println()

	//Question 3.

	s1 := s[0:5]
	s2 := s[5:10]
	s3 := s[2:7]
	s4 := s[1:6]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	//Question 4. 1.

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x1 := append(x, 52)
	fmt.Println(x1)

	// 2.

	fmt.Println(x1, (append(x1, 53, 54, 55)))

	// 3.

	y := []int{56, 57, 58, 59, 60}

	y1 := append(x1, y...)
	fmt.Println(y1)
}
